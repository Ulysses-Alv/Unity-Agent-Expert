package update

import (
	"fmt"
	"strconv"
	"strings"
)

// Version represents a semantic version with major, minor, and patch components.
type Version struct {
	Major int
	Minor int
	Patch int
}

// ParseVersion parses a semantic version string (e.g., "1.2.3" or "v1.2.3").
// Returns a Version struct and nil error on success.
// Returns an error on invalid input.
func ParseVersion(s string) (Version, error) {
	// Strip leading 'v' if present
	s = strings.TrimPrefix(s, "v")

	// Split by '.'
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return Version{}, &VersionParseError{Version: s, Err: fmt.Errorf("expected 3 parts separated by '.'")}
	}

	// Check for leading zeros (semver doesn't allow them except for 0 itself)
	for _, part := range parts {
		if len(part) > 1 && strings.HasPrefix(part, "0") {
			return Version{}, &VersionParseError{Version: s, Err: fmt.Errorf("leading zeros not allowed in semver")}
		}
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return Version{}, &VersionParseError{Version: s, Err: fmt.Errorf("invalid major component: %w", err)}
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return Version{}, &VersionParseError{Version: s, Err: fmt.Errorf("invalid minor component: %w", err)}
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return Version{}, &VersionParseError{Version: s, Err: fmt.Errorf("invalid patch component: %w", err)}
	}

	return Version{Major: major, Minor: minor, Patch: patch}, nil
}

// CompareVersions compares two versions.
// Returns -1 if a < b, 0 if a == b, +1 if a > b.
func CompareVersions(a, b Version) int {
	if a.Major != b.Major {
		if a.Major < b.Major {
			return -1
		}
		return 1
	}

	if a.Minor != b.Minor {
		if a.Minor < b.Minor {
			return -1
		}
		return 1
	}

	if a.Patch != b.Patch {
		if a.Patch < b.Patch {
			return -1
		}
		return 1
	}

	return 0
}

// IsDevVersion returns true if the version string looks like a dev/development build.
// Dev versions are not valid semver strings (e.g., "dev", "development", "main", "").
func IsDevVersion(s string) bool {
	if s == "" {
		return true
	}

	// Strip leading 'v' if present
	s = strings.TrimPrefix(s, "v")

	// Check if it contains any non-numeric parts (besides dots and v prefix)
	// A valid semver has exactly 3 numeric parts separated by dots
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return true
	}

	for _, part := range parts {
		if _, err := strconv.Atoi(part); err != nil {
			return true
		}
	}

	// Check for leading zeros (not allowed in semver)
	// But 0.0.0 is valid
	for i, part := range parts {
		if len(part) > 1 && strings.HasPrefix(part, "0") {
			// Leading zeros only allowed if the number is exactly 0
			// "01", "00" etc are not valid semver
			if i == 0 && part == "0" {
				continue
			}
			if part != "0" {
				return true
			}
		}
	}

	return false
}

// String returns the string representation of a Version.
func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}
