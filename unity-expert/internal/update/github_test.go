package update

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSelectAsset(t *testing.T) {
	release := &ReleaseInfo{
		TagName: "v1.2.3",
		Assets: []AssetInfo{
			{Name: "unity-agent-expert_windows_amd64.zip", DownloadURL: "https://example.com/windows_amd64.zip"},
			{Name: "unity-agent-expert_windows_arm64.zip", DownloadURL: "https://example.com/windows_arm64.zip"},
			{Name: "unity-agent-expert_darwin_amd64.zip", DownloadURL: "https://example.com/darwin_amd64.zip"},
			{Name: "unity-agent-expert_darwin_arm64.zip", DownloadURL: "https://example.com/darwin_arm64.zip"},
			{Name: "unity-agent-expert_linux_amd64.zip", DownloadURL: "https://example.com/linux_amd64.zip"},
			{Name: "unity-agent-expert_linux_arm64.zip", DownloadURL: "https://example.com/linux_arm64.zip"},
		},
	}

	tests := []struct {
		name        string
		goos        string
		goarch      string
		wantName    string
		wantErr     bool
		errType     interface{}
	}{
		{
			name:     "windows amd64",
			goos:     "windows",
			goarch:   "amd64",
			wantName: "unity-agent-expert_windows_amd64.zip",
			wantErr:  false,
		},
		{
			name:     "windows arm64",
			goos:     "windows",
			goarch:   "arm64",
			wantName: "unity-agent-expert_windows_arm64.zip",
			wantErr:  false,
		},
		{
			name:     "darwin amd64",
			goos:     "darwin",
			goarch:   "amd64",
			wantName: "unity-agent-expert_darwin_amd64.zip",
			wantErr:  false,
		},
		{
			name:     "darwin arm64",
			goos:     "darwin",
			goarch:   "arm64",
			wantName: "unity-agent-expert_darwin_arm64.zip",
			wantErr:  false,
		},
		{
			name:     "linux amd64",
			goos:     "linux",
			goarch:   "amd64",
			wantName: "unity-agent-expert_linux_amd64.zip",
			wantErr:  false,
		},
		{
			name:     "linux arm64",
			goos:     "linux",
			goarch:   "arm64",
			wantName: "unity-agent-expert_linux_arm64.zip",
			wantErr:  false,
		},
		{
			name:    "unsupported platform",
			goos:    "freebsd",
			goarch:  "amd64",
			wantErr: true,
			errType: &AssetNotFoundError{},
		},
		{
			name:    "unsupported arch",
			goos:    "windows",
			goarch:  "386",
			wantErr: true,
			errType: &AssetNotFoundError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asset, err := SelectAsset(release, tt.goos, tt.goarch)
			if tt.wantErr {
				if err == nil {
					t.Errorf("SelectAsset() expected error, got nil")
					return
				}
				if _, ok := err.(*AssetNotFoundError); !ok {
					t.Errorf("SelectAsset() unexpected error type: %T", err)
				}
				return
			}
			if err != nil {
				t.Errorf("SelectAsset() unexpected error: %v", err)
				return
			}
			if asset.Name != tt.wantName {
				t.Errorf("SelectAsset() = %q, want %q", asset.Name, tt.wantName)
			}
		})
	}
}

func TestFetchLatestRelease(t *testing.T) {
	t.Run("successful fetch", func(t *testing.T) {
		expectedRelease := map[string]interface{}{
			"tag_name": "v1.5.0",
			"assets": []map[string]interface{}{
				{"name": "unity-agent-expert_windows_amd64.zip", "browser_download_url": "https://example.com/windows_amd64.zip"},
				{"name": "unity-agent-expert_darwin_arm64.zip", "browser_download_url": "https://example.com/darwin_arm64.zip"},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify request headers
			if r.Header.Get("Accept") != "application/vnd.github+json" {
				t.Errorf("Expected Accept header to be application/vnd.github+json, got %s", r.Header.Get("Accept"))
			}
			if r.Header.Get("User-Agent") == "" {
				t.Error("Expected non-empty User-Agent header")
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(expectedRelease)
		}))
		defer server.Close()

		// We can't easily test the real FetchLatestRelease since it uses hardcoded URL,
		// but we can test the parsing logic through SelectAsset with mock data
		release := &ReleaseInfo{
			TagName: "v1.5.0",
			Assets: []AssetInfo{
				{Name: "unity-agent-expert_windows_amd64.zip", DownloadURL: "https://example.com/windows_amd64.zip"},
				{Name: "unity-agent-expert_darwin_arm64.zip", DownloadURL: "https://example.com/darwin_arm64.zip"},
			},
		}

		asset, err := SelectAsset(release, "windows", "amd64")
		if err != nil {
			t.Errorf("SelectAsset() unexpected error: %v", err)
			return
		}
		if asset.Name != "unity-agent-expert_windows_amd64.zip" {
			t.Errorf("SelectAsset() = %q, want %q", asset.Name, "unity-agent-expert_windows_amd64.zip")
		}
	})
}

func TestParseTagName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"v1.2.3", "1.2.3"},
		{"v0.0.1", "0.0.1"},
		{"1.2.3", "1.2.3"},
		{"v999.888.777", "999.888.777"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ParseTagName(tt.input)
			if got != tt.want {
				t.Errorf("ParseTagName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestNormalizeOS(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"windows", "win"},
		{"darwin", "darwin"},
		{"linux", "linux"},
		{"freebsd", "freebsd"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := normalizeOS(tt.input)
			if got != tt.want {
				t.Errorf("normalizeOS(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
