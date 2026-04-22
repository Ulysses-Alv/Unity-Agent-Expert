package update

import (
	"testing"
)

func TestParseVersion(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Version
		wantErr bool
	}{
		// Valid versions
		{
			name:    "standard version",
			input:   "1.2.3",
			want:    Version{Major: 1, Minor: 2, Patch: 3},
			wantErr: false,
		},
		{
			name:    "version with v prefix",
			input:   "v1.2.3",
			want:    Version{Major: 1, Minor: 2, Patch: 3},
			wantErr: false,
		},
		{
			name:    "zero version",
			input:   "0.0.0",
			want:    Version{Major: 0, Minor: 0, Patch: 0},
			wantErr: false,
		},
		{
			name:    "large numbers",
			input:   "100.200.300",
			want:    Version{Major: 100, Minor: 200, Patch: 300},
			wantErr: false,
		},

		// Invalid versions
		{
			name:    "dev string",
			input:   "dev",
			want:    Version{},
			wantErr: true,
		},
		{
			name:    "only two parts",
			input:   "1.2",
			want:    Version{},
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			want:    Version{},
			wantErr: true,
		},
		{
			name:    "non-numeric major",
			input:   "a.2.3",
			want:    Version{},
			wantErr: true,
		},
		{
			name:    "non-numeric minor",
			input:   "1.b.3",
			want:    Version{},
			wantErr: true,
		},
		{
			name:    "non-numeric patch",
			input:   "1.2.c",
			want:    Version{},
			wantErr: true,
		},

		// Edge cases - leading zeros (valid semver forbids these)
		{
			name:    "leading zero in major",
			input:   "01.2.3",
			want:    Version{},
			wantErr: true,
		},
		{
			name:    "leading zero in minor",
			input:   "1.02.3",
			want:    Version{},
			wantErr: true,
		},
		{
			name:    "leading zero in patch",
			input:   "1.2.03",
			want:    Version{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseVersion(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseVersion(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ParseVersion(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		name string
		a    Version
		b    Version
		want int
	}{
		{
			name: "equal versions",
			a:    Version{Major: 1, Minor: 2, Patch: 3},
			b:    Version{Major: 1, Minor: 2, Patch: 3},
			want: 0,
		},
		{
			name: "major less",
			a:    Version{Major: 1, Minor: 2, Patch: 3},
			b:    Version{Major: 2, Minor: 0, Patch: 0},
			want: -1,
		},
		{
			name: "major greater",
			a:    Version{Major: 2, Minor: 0, Patch: 0},
			b:    Version{Major: 1, Minor: 2, Patch: 3},
			want: 1,
		},
		{
			name: "minor less",
			a:    Version{Major: 1, Minor: 1, Patch: 3},
			b:    Version{Major: 1, Minor: 2, Patch: 3},
			want: -1,
		},
		{
			name: "minor greater",
			a:    Version{Major: 1, Minor: 3, Patch: 3},
			b:    Version{Major: 1, Minor: 2, Patch: 3},
			want: 1,
		},
		{
			name: "patch less",
			a:    Version{Major: 1, Minor: 2, Patch: 2},
			b:    Version{Major: 1, Minor: 2, Patch: 3},
			want: -1,
		},
		{
			name: "patch greater",
			a:    Version{Major: 1, Minor: 2, Patch: 4},
			b:    Version{Major: 1, Minor: 2, Patch: 3},
			want: 1,
		},
		{
			name: "zero versions",
			a:    Version{Major: 0, Minor: 0, Patch: 0},
			b:    Version{Major: 0, Minor: 0, Patch: 0},
			want: 0,
		},
		{
			name: "zero vs non-zero",
			a:    Version{Major: 0, Minor: 0, Patch: 0},
			b:    Version{Major: 1, Minor: 0, Patch: 0},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CompareVersions(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("CompareVersions(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestIsDevVersion(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		// Dev versions (should return true)
		{
			name:  "empty string",
			input: "",
			want:  true,
		},
		{
			name:  "dev",
			input: "dev",
			want:  true,
		},
		{
			name:  "development",
			input: "development",
			want:  true,
		},
		{
			name:  "main",
			input: "main",
			want:  true,
		},
		{
			name:  "only one part",
			input: "1",
			want:  true,
		},
		{
			name:  "only two parts",
			input: "1.2",
			want:  true,
		},
		{
			name:  "v prefix only",
			input: "v",
			want:  true,
		},
		{
			name:  "four parts",
			input: "1.2.3.4",
			want:  true,
		},
		{
			name:  "with suffix",
			input: "1.2.3-beta",
			want:  true,
		},
		{
			name:  "leading zero in major",
			input: "01.2.3",
			want:  true,
		},

		// Valid versions (should return false)
		{
			name:  "standard version",
			input: "1.2.3",
			want:  false,
		},
		{
			name:  "v prefix version",
			input: "v1.2.3",
			want:  false,
		},
		{
			name:  "zero version",
			input: "0.0.0",
			want:  false,
		},
		{
			name:  "large numbers",
			input: "100.200.300",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsDevVersion(tt.input)
			if got != tt.want {
				t.Errorf("IsDevVersion(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestVersionString(t *testing.T) {
	v := Version{Major: 1, Minor: 2, Patch: 3}
	want := "v1.2.3"
	got := v.String()
	if got != want {
		t.Errorf("Version.String() = %q, want %q", got, want)
	}
}
