# unity-agent-expert CLI Architecture Proposal

**Based on:** gentle-ai CLI architecture analysis  
**Date:** 2026-04-21  
**Change:** bootstrapper-cli-redesign

---

## Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         BOOTSTRAPPER (PowerShell)                            │
│  install.ps1                                                                 │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │ 1. Check prerequisites (git, PowerShell 5.1+)                        │   │
│  │ 2. Clone/download repo to temp                                       │   │
│  │ 3. Install Go binary:                                                │   │
│  │    • go install github.com/Ulysses-Alv/unity-agent-expert@latest     │   │
│  │    • OR download from GitHub Releases                                │   │
│  │ 4. Validate installation                                             │   │
│  │ 5. Call: unity-expert install --provider claude                      │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────────────────┘
                                       │
                                       ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                         Go CLI (unity-expert)                                │
│  cmd/unity-expert/main.go                                                    │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │ main()                                                                │   │
│  │   app.Version = "1.0.0"                                              │   │
│  │   app.Run(os.Args[1:])                                               │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
│                                       │                                      │
│                                       ▼                                      │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │ internal/app/app.go - Command Routing                                 │   │
│  │                                                                       │   │
│  │ switch args[0] {                                                      │   │
│  │   case "install":  return cli.RunInstall(args[1:])                   │   │
│  │   case "sync":     return cli.RunSync(args[1:])                      │   │
│  │   case "uninstall": return cli.RunUninstall(args[1:])                │   │
│  │   case "version":  print version                                     │   │
│  │   case "help":     print help                                        │   │
│  │ }                                                                     │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────────────────┘
                                       │
          ┌────────────────────────────┼────────────────────────────┐
          │                            │                            │
          ▼                            ▼                            ▼
┌──────────────────┐        ┌──────────────────┐        ┌──────────────────┐
│   INSTALL        │        │     SYNC         │        │   UNINSTALL      │
│   COMMAND        │        │    COMMAND       │        │    COMMAND       │
├──────────────────┤        ├──────────────────┤        ├──────────────────┤
│ Parse Flags:     │        │ Parse Flags:     │        │ Parse Flags:     │
│ • --agents       │        │ • --agents       │        │ • --agents       │
│ • --provider     │        │ • --provider     │        │ • --force        │
│ • --dry-run      │        │ • --dry-run      │        │                  │
│ • --force        │        │                  │        │                  │
│                  │        │                  │        │                  │
│ Pipeline:        │        │ Pipeline:        │        │ Pipeline:        │
│ 1. Prepare       │        │ 1. Prepare       │        │ 1. Prepare       │
│    • Backup      │        │    • Backup      │        │    • Backup      │
│ 2. Apply         │        │ 2. Apply         │        │ 2. Apply         │
│    • Merge JSON │        │    • Merge JSON  │        │    • Remove      │
│    • Copy files │        │    • Copy files  │        │    • Cleanup     │
│    • Set models │        │    • Set models  │        │                  │
│ 3. Verify        │        │ 3. Verify        │        │ 3. Verify        │
└──────────────────┘        └──────────────────┘        └──────────────────┘
          │                            │                            │
          └────────────────────────────┼────────────────────────────┘
                                       │
                                       ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                      internal/pipeline/orchestrator.go                       │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │ type Orchestrator struct {                                            │   │
│  │     rollbackPolicy RollbackPolicy                                    │   │
│  │     failurePolicy  FailurePolicy                                     │   │
│  │ }                                                                     │   │
│  │                                                                       │   │
│  │ func (o *Orchestrator) Execute(plan StagePlan) ExecutionResult {     │   │
│  │     // 1. Run Prepare steps (backup)                                 │   │
│  │     for _, step := range plan.Prepare {                              │   │
│  │         step.Run()                                                   │   │
│  │     }                                                                 │   │
│  │                                                                       │   │
│  │     // 2. Run Apply steps (inject)                                   │   │
│  │     for _, step := range plan.Apply {                                │   │
│  │         if err := step.Run(); err != nil {                           │   │
│  │             o.rollback()                                             │   │
│  │             return ExecutionResult{Err: err}                         │   │
│  │         }                                                             │   │
│  │     }                                                                 │   │
│  │                                                                       │   │
│  │     // 3. Verify                                                      │   │
│  │     return verify()                                                   │   │
│  │ }                                                                     │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────────────────┘
                                       │
                                       ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                         internal/agents/opencode.go                          │
│  ┌──────────────────────────────────────────────────────────────────────┐   │
│  │ type OpenCodeAdapter struct {}                                        │   │
│  │                                                                       │   │
│  │ func (a OpenCodeAdapter) Inject(homeDir string, config AgentConfig)  │   │
│  │     error {                                                           │   │
│  │     // 1. Read opencode.json                                          │   │
│  │     // 2. Merge agents from embedded agents.json                      │   │
│  │     // 3. Apply model assignments from presets.json                   │   │
│  │     // 4. Write back (UTF-8, no BOM)                                  │   │
│  │ }                                                                     │   │
│  └──────────────────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────────────────┘
                                       │
                                       ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                         Target: opencode.json                                │
│  ~/.config/opencode/opencode.json                                            │
│  {                                                                           │
│    "agent": {                                                                │
│      "unity-6000-expert": {                                                  │
│        "model": "anthropic/claude-sonnet-4-20250514",                        │
│        "prompt": "{file:../prompts/unity/unity-6000-expert.md}",             │
│        ...                                                                   │
│      },                                                                      │
│      "unity-graphics-expert": { ... },                                       │
│      ...                                                                     │
│    }                                                                         │
│  }                                                                           │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Component Comparison Matrix

| Component | Current PS1 | gentle-ai | Proposed Go CLI |
|-----------|-------------|-----------|-----------------|
| **Binary Installation** | ❌ N/A | ✅ Binary download / go install | ✅ go install / Releases |
| **Repo Download** | ✅ Git clone | ✅ Git clone / Releases | ✅ Git clone / Releases |
| **JSON Merge** | ✅ PowerShell ConvertFrom/To-Json | ✅ Go encoding/json | ✅ Go encoding/json |
| **File Copy** | ✅ Copy-Item | ✅ embed.FS + io.Copy | ✅ embed.FS + io.Copy |
| **Model Presets** | ✅ presets.json | ✅ presets.json | ✅ presets.json |
| **Provider Flag** | ✅ `-Provider claude` | ✅ `--preset` | ✅ `--provider claude` |
| **Dry Run** | ✅ `-DryRun` | ✅ `--dry-run` | ✅ `--dry-run` |
| **Force Reinstall** | ✅ `-Force` | ✅ State-based | ✅ `--force` |
| **Backup** | ❌ No | ✅ tar.gz + manifest | ✅ tar.gz + manifest |
| **Rollback** | ❌ No | ✅ Restore from backup | ✅ Restore from backup |
| **State Persistence** | ❌ No | ✅ state.json | ✅ state.json |
| **Sync Command** | ❌ No | ✅ `gentle-ai sync` | ✅ `unity-expert sync` |
| **Uninstall** | ❌ No | ✅ `gentle-ai uninstall` | ✅ `unity-expert uninstall` |
| **TUI** | ❌ No | ✅ Bubble Tea | ❌ (CLI only for now) |
| **Verification** | ✅ Basic checks | ✅ Comprehensive | ✅ Basic + JSON validate |

---

## Logic Migration: PS1 → Go

### What Moves to Go

| Logic | Current PS1 | Go Implementation |
|-------|-------------|-------------------|
| **JSON Merge** | `ConvertFrom-Json` + `Add-Member` | `encoding/json` + map merge |
| **Model Assignment** | Loop through presets | Map lookup + apply |
| **File Copy** | `Copy-Item -Recurse` | `embed.FS` + `io.Copy` |
| **Dry Run** | `if ($DryRun)` checks | `flags.DryRun` conditional |
| **Force Flag** | `if (-not $Force)` | `flags.Force` check |
| **Provider Selection** | `$ProviderChoice` param | `--provider` flag |
| **Backup** | ❌ N/A | `backup.Snapshotter` |
| **State** | ❌ N/A | `state.Read/Write` |

### What Stays in PS1

| Logic | Reason |
|-------|--------|
| **Git clone** | Bootstrapping only, Go doesn't need git |
| **Binary install** | PS1 installs the Go binary |
| **PATH check** | OS-level concern |
| **Prerequisites** | git, PowerShell version check |
| **First-run detection** | Bootstrapper concern |

---

## File Operations Deep Dive

### 1. JSON Merge (agents.json → opencode.json)

**Current PS1:**
```powershell
$config = Get-Content $configFile -Raw | ConvertFrom-Json
$agentsToAdd = Get-Content $agentsFile -Raw | ConvertFrom-Json

foreach ($agentName in $agentsToAdd.PSObject.Properties.Name) {
    $agentEntry = $agentsToAdd.$agentName | Select-Object *
    
    # Apply model from preset
    if ($presetModel) {
        $agentEntry.model = $presetModel
    }
    
    $config.agent | Add-Member -NotePropertyName $agentName -NotePropertyValue $agentEntry -Force
}

$jsonOutput = $config | ConvertTo-Json -Depth 20
[System.IO.File]::WriteAllText($configFile, $jsonOutput, [System.Text.UTF8Encoding]::new($false))
```

**Go Implementation:**
```go
func (a OpenCodeAdapter) Inject(homeDir string, agents map[string]AgentConfig, presets PresetConfig) error {
    // 1. Read existing opencode.json
    configPath := filepath.Join(homeDir, ".config/opencode/opencode.json")
    configData, err := os.ReadFile(configPath)
    if err != nil {
        return fmt.Errorf("read opencode.json: %w", err)
    }
    
    var config OpenCodeConfig
    if err := json.Unmarshal(configData, &config); err != nil {
        return fmt.Errorf("parse opencode.json: %w", err)
    }
    
    // 2. Merge agents
    for name, agent := range agents {
        // Apply model from preset if available
        if presetModel, ok := presets.Models[name]; ok {
            agent.Model = presetModel
        }
        config.Agents[name] = agent
    }
    
    // 3. Write back (UTF-8, no BOM)
    output, err := json.MarshalIndent(config, "", "  ")
    if err != nil {
        return fmt.Errorf("marshal config: %w", err)
    }
    
    if err := os.WriteFile(configPath, output, 0644); err != nil {
        return fmt.Errorf("write opencode.json: %w", err)
    }
    
    return nil
}
```

### 2. File Copy (Prompts)

**Current PS1:**
```powershell
$promptsSource = Join-Path $RepoDir "opencode\prompts\unity"
$promptsTarget = Join-Path $INSTALL_DIR "prompts\unity"

if (-not (Test-Path $promptsTarget)) {
    New-Item -ItemType Directory -Path $promptsTarget -Force | Out-Null
}

Get-ChildItem $promptsSource -Filter "*.md" | ForEach-Object {
    $targetFile = Join-Path $promptsTarget $_.Name
    Copy-Item $_.FullName -Destination $targetFile -Force
}
```

**Go Implementation:**
```go
//go:embed prompts/unity/*.md
var promptFS embed.FS

func copyPrompts(homeDir string) error {
    targetDir := filepath.Join(homeDir, ".config/opencode/prompts/unity")
    
    if err := os.MkdirAll(targetDir, 0755); err != nil {
        return fmt.Errorf("create prompts dir: %w", err)
    }
    
    entries, err := promptFS.ReadDir("prompts/unity")
    if err != nil {
        return fmt.Errorf("read embedded prompts: %w", err)
    }
    
    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }
        
        src, err := promptFS.Open(filepath.Join("prompts/unity", entry.Name()))
        if err != nil {
            return err
        }
        defer src.Close()
        
        dstPath := filepath.Join(targetDir, entry.Name())
        dst, err := os.Create(dstPath)
        if err != nil {
            return err
        }
        defer dst.Close()
        
        if _, err := io.Copy(dst, src); err != nil {
            return err
        }
    }
    
    return nil
}
```

### 3. Backup Creation

**gentle-ai Approach:**
```go
func (s *Snapshotter) Snapshot(targets []string, backupRoot string) (Manifest, error) {
    timestamp := time.Now().UTC().Format("20060102150405.000000000")
    backupDir := filepath.Join(backupRoot, timestamp)
    
    if err := os.MkdirAll(backupDir, 0755); err != nil {
        return Manifest{}, err
    }
    
    // Create tar.gz
    archivePath := filepath.Join(backupDir, "snapshot.tar.gz")
    archive, err := createTarGz(targets, archivePath)
    if err != nil {
        return Manifest{}, err
    }
    
    // Write manifest
    manifest := Manifest{
        ID:          timestamp,
        CreatedAt:   time.Now().UTC(),
        Source:      "install",
        Description: "pre-install snapshot",
        Targets:     targets,
    }
    
    manifestPath := filepath.Join(backupDir, "manifest.json")
    if err := writeJSON(manifest, manifestPath); err != nil {
        return Manifest{}, err
    }
    
    return manifest, nil
}
```

---

## Data Flow Diagrams

### Install Command Flow

```
┌──────────────┐
│ User runs:   │
│ unity-expert │
│ install      │
│ --provider   │
│ claude       │
└──────┬───────┘
       │
       ▼
┌──────────────────────────────┐
│ cli.ParseInstallFlags()      │
│ Returns:                     │
│ • Agents: ["unity-6000-..."] │
│ • Provider: "claude"         │
│ • DryRun: false              │
│ • Force: false               │
└──────┬───────────────────────┘
       │
       ▼
┌──────────────────────────────┐
│ cli.RunInstall()             │
│ 1. Discover agents           │
│ 2. Load presets              │
│ 3. Build StagePlan           │
└──────┬───────────────────────┘
       │
       ▼
┌──────────────────────────────┐
│ Pipeline.Execute()           │
│                              │
│ PREPARE:                     │
│ • backup.Snapshot()          │
│   → manifest.json            │
│   → snapshot.tar.gz          │
│                              │
│ APPLY:                       │
│ • opencodeAdapter.Inject()   │
│   → Merge agents.json        │
│   → Apply claude preset      │
│ • copyPrompts()              │
│   → prompts/unity/*.md       │
│ • copySkills()               │
│   → skills/unity-6000/       │
│                              │
│ VERIFY:                      │
│ • Check files exist          │
│ • Validate JSON              │
└──────┬───────────────────────┘
       │
       ▼
┌──────────────────────────────┐
│ cli.RenderInstallReport()    │
│ [OK] Added 18 agent(s)       │
│ [OK] Copied 18 prompt(s)     │
│ [OK] Copied 19 skill(s)      │
└──────────────────────────────┘
```

### Sync Command Flow (Idempotent)

```
┌──────────────┐
│ User runs:   │
│ unity-expert │
│ sync         │
│ --dry-run    │
└──────┬───────┘
       │
       ▼
┌──────────────────────────────┐
│ cli.ParseSyncFlags()         │
│ Returns:                     │
│ • Agents: from state.json    │
│ • DryRun: true               │
└──────┬───────────────────────┘
       │
       ▼
┌──────────────────────────────┐
│ cli.RunSync()                │
│ 1. Load state.json           │
│ 2. Discover installed agents │
│ 3. Build sync selection      │
└──────┬───────────────────────┘
       │
       ▼
┌──────────────────────────────┐
│ Pipeline.Execute()           │
│                              │
│ PREPARE:                     │
│ • backup.Snapshot()          │
│                              │
│ APPLY:                       │
│ • opencodeAdapter.Inject()   │
│   → Merge (preserve models)  │
│ • copyPrompts()              │
│   → Overwrite if changed     │
│ • copySkills()               │
│   → Overwrite if changed     │
│                              │
│ VERIFY:                      │
│ • Check files exist          │
└──────┬───────────────────────┘
       │
       ▼
┌──────────────────────────────┐
│ cli.RenderSyncReport()       │
│ gentle-ai sync — no managed  │
│ sync actions needed          │
│ All managed assets are up    │
│ to date. No files changed.   │
└──────────────────────────────┘
```

---

## State Management

### state.json Schema

```json
{
  "version": 1,
  "installedAt": "2026-04-21T23:30:00Z",
  "installedAgents": [
    "unity-6000-expert",
    "unity-graphics-expert",
    "unity-2d-expert"
  ],
  "provider": "claude",
  "modelAssignments": {
    "unity-6000-expert": "anthropic/claude-sonnet-4-20250514",
    "unity-graphics-expert": "anthropic/claude-sonnet-4-20250514"
  }
}
```

### State Read/Write

```go
// internal/state/state.go
type InstallState struct {
    Version          int                    `json:"version"`
    InstalledAt      time.Time              `json:"installedAt"`
    InstalledAgents  []string               `json:"installedAgents"`
    Provider         string                 `json:"provider"`
    ModelAssignments map[string]string      `json:"modelAssignments"`
}

func Read(homeDir string) (InstallState, error) {
    path := filepath.Join(homeDir, ".unity-expert", "state.json")
    data, err := os.ReadFile(path)
    if err != nil {
        return InstallState{}, err
    }
    
    var state InstallState
    if err := json.Unmarshal(data, &state); err != nil {
        return InstallState{}, err
    }
    
    return state, nil
}

func Write(homeDir string, state InstallState) error {
    dir := filepath.Join(homeDir, ".unity-expert")
    if err := os.MkdirAll(dir, 0755); err != nil {
        return err
    }
    
    path := filepath.Join(dir, "state.json")
    data, err := json.MarshalIndent(state, "", "  ")
    if err != nil {
        return err
    }
    
    return os.WriteFile(path, data, 0644)
}
```

---

## Testing Strategy

### Unit Tests

```go
// internal/cli/install_test.go
func TestParseInstallFlags(t *testing.T) {
    tests := []struct {
        name     string
        args     []string
        expected InstallFlags
    }{
        {
            name: "with provider",
            args: []string{"--provider", "claude"},
            expected: InstallFlags{
                Provider: "claude",
                DryRun:   false,
            },
        },
        {
            name: "dry run",
            args: []string{"--dry-run"},
            expected: InstallFlags{
                DryRun: true,
            },
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            flags, err := ParseInstallFlags(tt.args)
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            if flags.Provider != tt.expected.Provider {
                t.Errorf("Provider = %q, want %q", flags.Provider, tt.expected.Provider)
            }
        })
    }
}
```

### Integration Tests

```go
// internal/agents/opencode_test.go
func TestOpenCodeAdapter_Inject(t *testing.T) {
    // Create temp dir
    tempDir := t.TempDir()
    
    // Create mock opencode.json
    configPath := filepath.Join(tempDir, "opencode.json")
    os.WriteFile(configPath, []byte(`{"agent":{}}`), 0644)
    
    // Load embedded agents.json
    agents := loadEmbeddedAgents()
    presets := loadEmbeddedPresets()
    
    // Inject
    adapter := OpenCodeAdapter{}
    err := adapter.Inject(tempDir, agents, presets)
    if err != nil {
        t.Fatalf("Inject failed: %v", err)
    }
    
    // Verify
    data, _ := os.ReadFile(configPath)
    var config OpenCodeConfig
    json.Unmarshal(data, &config)
    
    if _, ok := config.Agents["unity-6000-expert"]; !ok {
        t.Error("unity-6000-expert not found in config")
    }
}
```

---

## Implementation Checklist

### Phase 1: Core CLI (Week 1)
- [ ] Initialize Go module
- [ ] Create cmd/unity-expert/main.go
- [ ] Implement internal/app/app.go (command routing)
- [ ] Implement internal/cli/install.go (flag parsing)
- [ ] Create internal/agents/opencode.go (JSON merge)
- [ ] Embed agents.json and presets.json
- [ ] Test: `unity-expert install --dry-run`

### Phase 2: File Operations (Week 2)
- [ ] Implement prompt file copying (embed.FS)
- [ ] Implement skill file copying
- [ ] Add model preset application
- [ ] Add --provider flag support
- [ ] Add --force flag support
- [ ] Test: Full install with --provider claude

### Phase 3: Backup & State (Week 3)
- [ ] Implement internal/backup/snapshotter.go
- [ ] Implement internal/state/state.go
- [ ] Add backup to install pipeline
- [ ] Add state persistence after install
- [ ] Test: Backup creation and restore

### Phase 4: Sync Command (Week 4)
- [ ] Implement internal/cli/sync.go
- [ ] Implement idempotent sync logic
- [ ] Add state loading for sync
- [ ] Add --dry-run for sync
- [ ] Test: Multiple sync runs (idempotent)

### Phase 5: Uninstall & Polish (Week 5)
- [ ] Implement internal/cli/uninstall.go
- [ ] Add verification reports
- [ ] Add error handling improvements
- [ ] Update install.ps1 bootstrapper
- [ ] Test: Full E2E flow

---

## Risk Mitigation

| Risk | Mitigation |
|------|------------|
| **JSON merge corrupts config** | Backup before merge, validate JSON syntax |
| **File copy overwrites user content** | Check for existing files, prompt or use --force |
| **State file lost** | Non-fatal: re-discover from filesystem |
| **Backup disk full** | Auto-prune old backups (keep 5 most recent) |
| **Provider preset not found** | Fallback to opencode default, warn user |
| **Embedded FS too large** | Only embed configs, not binaries |

---

## Success Criteria

1. ✅ `unity-expert install --provider claude` installs all agents with Claude models
2. ✅ `unity-expert install --dry-run` shows plan without changes
3. ✅ `unity-expert sync` re-applies config idempotently
4. ✅ Backup created before every install/sync
5. ✅ State persisted to ~/.unity-expert/state.json
6. ✅ Bootstrapper install.ps1 calls Go CLI after binary install
7. ✅ All unit tests pass
8. ✅ E2E test: Fresh install → sync → uninstall

---

## References

- gentle-ai source: https://github.com/Gentleman-Programming/gentle-ai
- Go embed package: https://pkg.go.dev/embed
- Go flag package: https://pkg.go.dev/flag
- Unity-Agent-Expert install.ps1: `D:\Projects\Skill creation\install.ps1`
