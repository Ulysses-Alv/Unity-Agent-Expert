package install

// InstallOptions wraps the CLI options for the install pipeline.
type InstallOptions struct {
	Provider string
	DryRun   bool
	Force    bool
	Agents   []string
}

// SyncOptions wraps the CLI options for the sync pipeline.
type SyncOptions struct {
	Provider string
	DryRun   bool
	Force    bool
}
