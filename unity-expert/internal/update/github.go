package update

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ReleaseInfo holds parsed GitHub release data.
type ReleaseInfo struct {
	TagName string
	Assets  []AssetInfo
}

// AssetInfo holds a single release asset.
type AssetInfo struct {
	Name        string
	DownloadURL string
}

// FetchLatestRelease fetches the latest release from GitHub API.
// Returns ReleaseInfo on success or an error on failure.
func FetchLatestRelease(owner, repo string) (*ReleaseInfo, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, &NetworkError{Err: err}
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "unity-agent-expert-update-checker")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, &NetworkError{Err: err}
	}
	defer resp.Body.Close()

	// Check for rate limiting (HTTP 403 with X-RateLimit-Remaining: 0)
	if resp.StatusCode == http.StatusForbidden {
		if remaining := resp.Header.Get("X-RateLimit-Remaining"); remaining == "0" {
			return nil, &RateLimitError{}
		}
	}

	if resp.StatusCode != http.StatusOK {
		return nil, &NetworkError{Err: fmt.Errorf("unexpected status code: %d", resp.StatusCode)}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &NetworkError{Err: fmt.Errorf("failed to read response body: %w", err)}
	}

	// Parse the release JSON
	var release struct {
		TagName string `json:"tag_name"`
		Assets  []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
		} `json:"assets"`
	}

	if err := json.Unmarshal(body, &release); err != nil {
		return nil, &NetworkError{Err: fmt.Errorf("failed to parse release JSON: %w", err)}
	}

	// Build AssetInfo slice
	assets := make([]AssetInfo, 0, len(release.Assets))
	for _, asset := range release.Assets {
		assets = append(assets, AssetInfo{
			Name:        asset.Name,
			DownloadURL: asset.BrowserDownloadURL,
		})
	}

	return &ReleaseInfo{
		TagName: release.TagName,
		Assets:  assets,
	}, nil
}

// SelectAsset selects the correct asset for the current platform (GOOS/GOARCH).
// Asset naming convention: unity-agent-expert_{os}_{arch}.zip
// Returns the matching AssetInfo or an error if no matching asset is found.
func SelectAsset(release *ReleaseInfo, goos, goarch string) (*AssetInfo, error) {
	// Use GOOS and GOARCH directly - the asset naming follows standard Go naming
	osName := goos
	archName := goarch

	expectedPattern := fmt.Sprintf("unity-agent-expert_%s_%s.zip", osName, archName)

	for i := range release.Assets {
		if release.Assets[i].Name == expectedPattern {
			return &release.Assets[i], nil
		}
	}

	return nil, &AssetNotFoundError{OS: goos, Arch: goarch}
}

// normalizeOS converts GOOS to GitHub asset naming convention.
func normalizeOS(goos string) string {
	switch goos {
	case "windows":
		return "win"
	default:
		return goos
	}
}

// normalizeArch converts GOARCH to GitHub asset naming convention.
func normalizeArch(goarch string) string {
	// GitHub uses "x86_64" not "amd64" for some Linux distros, but
	// goreleaser typically uses amd64/arm64
	return goarch
}

// GetDownloadURL returns the full download URL for an asset.
// If the asset already has a full URL, it returns it as-is.
// Otherwise, it constructs the URL from the owner/repo and asset name.
func GetDownloadURL(asset *AssetInfo, owner, repo string) string {
	if asset.DownloadURL != "" {
		return asset.DownloadURL
	}
	// Fallback: construct URL from asset name
	return fmt.Sprintf("https://github.com/%s/%s/releases/download/latest/%s", owner, repo, url.PathEscape(asset.Name))
}

// ParseTagName extracts the version string from a tag name, stripping "v" prefix.
func ParseTagName(tagName string) string {
	return strings.TrimPrefix(tagName, "v")
}
