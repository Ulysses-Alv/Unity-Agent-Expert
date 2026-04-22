package cli

import (
	"flag"
	"testing"
)

func TestParseUpdateFlags(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantOpts  UpdateOptions
		wantErr   bool
		errType   error
	}{
		{
			name:    "no flags",
			args:    []string{},
			wantOpts: UpdateOptions{CheckOnly: false},
			wantErr: false,
		},
		{
			name:    "check-only flag",
			args:    []string{"--check-only"},
			wantOpts: UpdateOptions{CheckOnly: true},
			wantErr: false,
		},
		{
			name:    "help flag",
			args:    []string{"--help"},
			wantOpts: UpdateOptions{},
			wantErr: true,
			errType: flag.ErrHelp,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts, err := ParseUpdateFlags(tt.args)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseUpdateFlags() expected error, got nil")
				} else if tt.errType != nil && err != tt.errType {
					t.Errorf("ParseUpdateFlags() error = %v, want %v", err, tt.errType)
				}
				return
			}
			if err != nil {
				t.Errorf("ParseUpdateFlags() unexpected error: %v", err)
				return
			}
			if opts.CheckOnly != tt.wantOpts.CheckOnly {
				t.Errorf("ParseUpdateFlags().CheckOnly = %v, want %v", opts.CheckOnly, tt.wantOpts.CheckOnly)
			}
		})
	}
}
