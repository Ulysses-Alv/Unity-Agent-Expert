# gentle-ai CLI Architecture Analysis

**Exploration Date:** 2026-04-21  
**Purpose:** Understand gentle-ai's CLI architecture to inform unity-agent-expert Go CLI redesign

---

## Executive Summary

gentle-ai implements a **production-grade Go CLI** with the following key characteristics:

1. **CLI Framework:** Uses Go's standard `flag` package (NOT cobra/urfave) with custom parsing
2. **Architecture:** Clean separation between CLI parsing (`internal/cli/`), business logic (`internal/`), and TUI (`internal/tui/`)
3. **Install Command:** Full pipeline with backup, component injection, and verification
4. **Sync Command:** Idempotent re-sync of managed assets without re-installing binaries
5. **Key Features:** Dry-run mode, automatic backups, rollback support, state persistence

---

## 1. gentle-ai CLI Structure

### 1.1 Repository Layout

```
gentle-ai/
├── cmd/gentle-ai/
│   └── main.go              # Entry point, version injection
├── internal/
│   ├── app/
│   │   └── app.go           # Main application logic, command routing
│   ├── cli/
│   │   ├── install.go       # Install command parsing
│   │   ├── sync.go          # Sync command parsing & execution
│   │   ├── uninstall.go     # Uninstall command
│   │   └── restore.go       # Backup restore
│   ├── model/
│   │   └── types.go         # Domain types (AgentID, ComponentID, etc.)
│   ├── pipeline/
│   │   └── orchestrator.go  # Stage execution with rollback
│   ├── backup/
│   │   └── snapshotter.go   # Backup creation/restore
│   ├── components/
│   │   ├── sdd/             # SDD orchestrator injection
│   │   ├── engram/          # Persistent memory setup
│   │   ├── skills/          # Skill file deployment
│   │   ├── mcp/             # MCP server config
│   │   └── ...
│   └── tui/
│       └── tui.go           # Bubble Tea interactive UI
└── scripts/
    ├── install.ps1          # Windows bootstrapper
    └── install.sh           # Unix bootstrapper
```

### 1.2 Command Flow Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                         main.go                                  │
│  var version = "dev"                                            │
│  app.Version = app.ResolveVersion(version)                      │
│  app.Run()                                                      │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      internal/app/app.go                        │
│  Run() → RunArgs(os.Args[1:])                                   │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │ Command Routing:                                          │   │
│  │ • version/--version/-v → print version                    │   │
│  │ • help/--help/-h → print help                             │   │
│  │ • install → cli.RunInstall()                              │   │
│  │ • sync → cli.RunSync()                                    │   │
│  │ • uninstall → cli.RunUninstall()                          │   │
│  │ • restore → cli.RunRestore()                              │   │
│  │ • (no args) → TUI (Bubble Tea)                            │   │
│  └──────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      internal/cli/*.go                          │
│  ParseInstallFlags() → InstallFlags struct                      │
│  ParseSyncFlags() → SyncFlags struct                            │
│  RunInstall() / RunSync() → ExecutionResult                     │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    internal/pipeline/                           │
│  Orchestrator.Execute(StagePlan)                                │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │ Prepare Stage:                                            │   │
│  │ • backup.NewSnapshotter().Snapshot()                      │   │
│  │ • Create backup manifest                                  │   │
│  └──────────────────────────────────────────────────────────┘   │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │ Apply Stage:                                              │   │
│  │ • componentApplyStep (per component)                      │   │
│  │ • agentInstallStep (per agent)                            │   │
│  │ • engram.Setup()                                          │   │
│  │ • persona.Inject()                                        │   │
│  └──────────────────────────────────────────────────────────┘   │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │ Rollback (on failure):                                    │   │
│  │ • rollbackRestoreStep                                     │   │
│  └──────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

---

## 2. Install Command Deep Dive

### 2.1 Command Signature

```bash
gentle-ai install [FLAGS]
```

### 2.2 Accepted Flags

| Flag | Type | Description |
|------|------|-------------|
| `--agent`, `--agents` | comma-separated list | Specific agents to install |
| `--component`, `--components` | comma-separated list | Specific components to install |
| `--skill`, `--skills` | comma-separated list | Specific skills to install |
| `--persona` | string | Persona to apply (gentleman, neutral, custom) |
| `--preset` | string | Preset to apply (full-gentleman, ecosystem-only, minimal) |
| `--sdd-mode` | string | SDD orchestrator mode: `single` or `multi` |
| `--dry-run` | boolean | Preview plan without executing |

**Flag Parsing Code:**
```go
func ParseInstallFlags(args []string) (InstallFlags, error) {
    var opts InstallFlags
    fs := flag.NewFlagSet("install", flag.ContinueOnError)
    registerListFlag(fs, "agent", &opts.Agents)  // CSV parsing
    fs.StringVar(&opts.Persona, "persona", "", "...")
    fs.BoolVar(&opts.DryRun, "dry-run", false, "...")
    // ...
}
```

### 2.3 Install Pipeline Steps

```
┌─────────────────────────────────────────────────────────────────┐
│ INSTALL PIPELINE                                                │
├─────────────────────────────────────────────────────────────────┤
│ PREPARE STAGE:                                                  │
│  1. Prepare backup snapshot                                     │
│     • Snapshot target files (agent configs, MCP, etc.)          │
│     • Create backup manifest with timestamp                     │
│     • Compress to tar.gz                                        │
│                                                                 │
│ APPLY STAGE:                                                    │
│  2. Rollback restore step (no-op on first run)                  │
│  3. Component injection (per component):                        │
│     • sdd.Inject() → Write SDD orchestrator prompts             │
│     • engram.Inject() → MCP config + system prompt              │
│     • skills.Inject() → Copy skill files                        │
│     • mcp.Inject() → Context7 MCP server config                 │
│     • gga.Inject() → GGA runtime config                         │
│     • permissions.Inject() → Permission model config            │
│     • theme.Inject() → Theme configuration                      │
│  4. Agent installation (per agent):                             │
│     • adapter.Install() → Agent-specific setup                  │
│     • Merge agents.json into opencode.json                      │
│     • Apply model assignments from preset                       │
│  5. Engram setup:                                               │
│     • Create ~/.engram/ directory                               │
│     • Initialize project context                                │
│  6. Persona injection:                                          │
│     • Write AGENTS.md / CLAUDE.md with persona                  │
│                                                                 │
│ VERIFY STAGE:                                                   │
│  7. runPostInstallVerification()                                │
│     • Check all expected files exist                            │
│     • Validate JSON syntax                                      │
│     • Render verification report                                │
└─────────────────────────────────────────────────────────────────┘
```

### 2.4 File Operations

**JSON Merge (agents.json → opencode.json):**
```go
// Conceptual flow from agents.Inject()
func Inject(homeDir string, adapter Adapter) (Result, error) {
    // 1. Read existing opencode.json
    config := readJSON(filepath.Join(homeDir, ".config/opencode/opencode.json"))
    
    // 2. Read agents.json from embedded FS
    agentsToAdd := readEmbedded("opencode/config/agents.json")
    
    // 3. Merge agents (preserve existing, add new)
    for agentName, agentConfig := range agentsToAdd {
        config.Agent[agentName] = agentConfig
    }
    
    // 4. Apply model assignments from preset
    if presetModel, ok := preset.Models[agentName]; ok {
        config.Agent[agentName].Model = presetModel
    }
    
    // 5. Write back with UTF-8 encoding (no BOM)
    writeJSON(config, filepath.Join(homeDir, ".config/opencode/opencode.json"))
    
    return Result{Changed: true}, nil
}
```

**File Copy (skills):**
```go
// Skills injection copies from embedded FS to target
func Inject(homeDir string, adapter Adapter, skillIDs []SkillID) error {
    targetDir := filepath.Join(homeDir, ".config/opencode/skills/unity-6000")
    os.MkdirAll(targetDir, 0755)
    
    for _, skillID := range skillIDs {
        src := embedFS.Open("skills/" + skillID)
        dst := filepath.Join(targetDir, skillID)
        io.Copy(dst, src)
    }
}
```

**Prompt Copy:**
```go
// Prompts are copied to prompts/unity/ directory
func copyPrompts(homeDir string) error {
    targetDir := filepath.Join(homeDir, ".config/opencode/prompts/unity")
    os.MkdirAll(targetDir, 0755)
    
    for _, prompt := range embeddedPrompts {
        dst := filepath.Join(targetDir, prompt.Name)
        os.WriteFile(dst, prompt.Content, 0644)
    }
}
```

---

## 3. Sync Command (Idempotent Re-configuration)

### 3.1 Purpose

The `sync` command is **crucial** for unity-agent-expert because it:
- Re-applies configuration without re-installing binaries
- Preserves user model assignments
- Is idempotent (safe to run multiple times)
- Supports `--dry-run` for preview

### 3.2 Sync vs Install

| Aspect | Install | Sync |
|--------|---------|------|
| Binary install | ✅ Yes | ❌ No |
| Backup creation | ✅ Yes | ✅ Yes |
| Component injection | ✅ Yes | ✅ Yes (inject only) |
| Agent setup | ✅ Full setup | ❌ Config only |
| Engram binary | ✅ Install | ❌ Skip |
| Persona injection | ✅ Yes | ❌ No |
| State persistence | ✅ Write state | ✅ Update state |

### 3.3 Sync Flags

```bash
gentle-ai sync [FLAGS]
```

| Flag | Type | Description |
|------|------|-------------|
| `--agent`, `--agents` | comma-separated | Agents to sync |
| `--skill`, `--skills` | comma-separated | Skills to sync |
| `--sdd-mode` | string | SDD mode (single/multi) |
| `--strict-tdd` | boolean | Enable strict TDD mode |
| `--include-permissions` | boolean | Include permissions component |
| `--include-theme` | boolean | Include theme component |
| `--profile` | comma-separated | SDD profiles (name:provider/model) |
| `--profile-phase` | comma-separated | Per-phase model assignments |
| `--dry-run` | boolean | Preview without executing |

### 3.4 Sync Pipeline

```
┌─────────────────────────────────────────────────────────────────┐
│ SYNC PIPELINE                                                   │
├─────────────────────────────────────────────────────────────────┤
│ PREPARE STAGE:                                                  │
│  1. Prepare backup snapshot (same as install)                   │
│                                                                 │
│ APPLY STAGE:                                                    │
│  2. Rollback restore step                                       │
│  3. Component sync (inject ONLY, no binary install):            │
│     • engram.Inject() → MCP config + protocol                   │
│     • mcp.Inject() → Context7 config                            │
│     • sdd.Inject() → SDD prompts (preserve model assignments)   │
│     • skills.Inject() → Skill files                             │
│     • gga.Inject() → GGA config (no binary)                     │
│     • permissions.Inject() → (opt-in)                           │
│     • theme.Inject() → (opt-in)                                 │
│                                                                 │
│ VERIFY STAGE:                                                   │
│  4. runPostSyncVerification()                                   │
│     • Check synced files exist                                  │
└─────────────────────────────────────────────────────────────────┘
```

### 3.5 State Persistence

gentle-ai persists installation state to `~/.gentle-ai/state.json`:

```json
{
  "installedAgents": ["opencode", "claude-code"],
  "modelAssignments": {
    "unity-6000-expert": {
      "providerID": "anthropic",
      "modelID": "claude-sonnet-4-20250514"
    }
  },
  "claudeModelAssignments": {
    "orchestrator": "claude-sonnet-4-20250514",
    "sdd-apply": "claude-haiku-3-5"
  }
}
```

**Key insight:** Sync loads this state to preserve model assignments across runs.

---

## 4. Backup System

### 4.1 Backup Structure

```
~/.gentle-ai/backups/
├── 20260421154302.123456789/    # Timestamp-based ID
│   ├── manifest.json            # Backup metadata
│   └── snapshot.tar.gz          # Compressed config files
├── 20260422103015.987654321/
│   ├── manifest.json
│   └── snapshot.tar.gz
└── ...
```

### 4.2 Backup Manifest

```json
{
  "id": "20260421154302.123456789",
  "createdAt": "2026-04-21T15:43:02.123456789Z",
  "source": "install",  // or "sync", "upgrade", "uninstall"
  "description": "pre-install snapshot",
  "appVersion": "1.22.0",
  "targets": [
    "C:\\Users\\PC\\.config\\opencode\\opencode.json",
    "C:\\Users\\PC\\.config\\opencode\\agents.json"
  ],
  "pinned": false
}
```

### 4.3 Backup Policy

- **Retention:** 5 most recent backups (auto-pruned)
- **Compression:** tar.gz
- **Deduplication:** Identical configs are not re-backed up
- **Pinning:** Users can pin backups via TUI (`p` key) to protect from pruning

---

## 5. Bootstrapper Scripts

### 5.1 Windows Bootstrapper (install.ps1)

gentle-ai's PowerShell script has **two installation methods**:

```powershell
# Method 1: Binary download (preferred)
irm https://raw.githubusercontent.com/Gentleman-Programming/gentle-ai/main/scripts/install.ps1 | iex

# Method 2: Go install
.\install.ps1 -Method go
```

**Key features:**
1. Platform detection (AMD64/ARM64)
2. Checksum verification (SHA256)
3. PATH validation
4. Version check via GitHub API

**Install location:**
```
%LOCALAPPDATA%\gentle-ai\bin\gentle-ai.exe
```

### 5.2 Post-Install Flow

After binary installation, users run:
```bash
gentle-ai          # Opens TUI installer
gentle-ai install  # CLI install with flags
```

---

## 6. Key Design Patterns

### 6.1 Pipeline Pattern

```go
type StagePlan struct {
    Prepare []Step  // Backup, validation
    Apply   []Step  // Component injection, agent setup
}

type Step interface {
    ID() string
    Run() error
}

type Orchestrator struct {
    rollbackPolicy RollbackPolicy
    failurePolicy  FailurePolicy
    progressFunc   ProgressFunc
}

func (o *Orchestrator) Execute(plan StagePlan) ExecutionResult {
    // 1. Run all Prepare steps
    for _, step := range plan.Prepare {
        if err := step.Run(); err != nil {
            return ExecutionResult{Err: err}
        }
    }
    
    // 2. Run all Apply steps with rollback on failure
    for _, step := range plan.Apply {
        if err := step.Run(); err != nil {
            if o.failurePolicy == RollbackOnError {
                o.rollback()
            }
            return ExecutionResult{Err: err}
        }
    }
    
    return ExecutionResult{Success: true}
}
```

### 6.2 Adapter Pattern for Agents

```go
type Adapter interface {
    Agent() AgentID
    SettingsPath(homeDir string) string
    MCPPath(homeDir string) string
    SystemPromptPath(homeDir string) string
    Install(homeDir string, config AgentConfig) error
}

// Example: OpenCode adapter
type OpenCodeAdapter struct{}

func (a OpenCodeAdapter) SettingsPath(homeDir string) string {
    return filepath.Join(homeDir, ".config/opencode/opencode.json")
}

func (a OpenCodeAdapter) Install(homeDir string, config AgentConfig) error {
    // Merge into opencode.json
}
```

### 6.3 Embedded Filesystem

gentle-ai embeds all config templates:

```go
//go:embed opencode/config/agents.json
var agentsConfig []byte

//go:embed prompts/unity/*.md
var promptFS embed.FS

//go:embed skills/unity-*/*
var skillFS embed.FS
```

---

## 7. Comparison: gentle-ai vs unity-agent-expert

### 7.1 Current unity-agent-expert (PowerShell)

| Feature | Current PS1 | gentle-ai |
|---------|-------------|-----------|
| Binary install | ❌ N/A | ✅ Yes |
| Config merge | ✅ Yes | ✅ Yes |
| Backup | ❌ No | ✅ Yes (compressed) |
| Dry-run | ✅ Yes | ✅ Yes |
| Rollback | ❌ No | ✅ Yes |
| State persistence | ❌ No | ✅ Yes (state.json) |
| TUI | ❌ No | ✅ Yes (Bubble Tea) |
| Sync command | ❌ No | ✅ Yes |
| Multi-agent | ✅ Yes | ✅ Yes |
| Model presets | ✅ Yes | ✅ Yes |

### 7.2 What to Move from PS1 to Go CLI

**Move to Go:**
1. ✅ JSON merge logic (agents.json → opencode.json)
2. ✅ Prompt file copying
3. ✅ Skill file copying
4. ✅ Model preset application
5. ✅ Dry-run mode
6. ✅ Force reinstall flag
7. ✅ Provider selection (`--provider`)
8. ✅ Backup creation (NEW)
9. ✅ State persistence (NEW)
10. ✅ Sync command (NEW - idempotent re-apply)

**Keep in PS1 (Bootstrapper only):**
1. ✅ Git clone / repo download
2. ✅ Go binary installation (or binary download)
3. ✅ PATH validation
4. ✅ Prerequisites check (git, PowerShell version)
5. ✅ First-run detection

---

## 8. Recommended Go CLI Structure for unity-expert

### 8.1 Proposed Command Set

```bash
unity-expert install [FLAGS]     # Install agents, prompts, skills
unity-expert sync [FLAGS]        # Re-apply config (idempotent)
unity-expert uninstall [FLAGS]   # Remove agents
unity-expert version             # Show version
unity-expert help                # Show help
```

### 8.2 Install Flags

```bash
unity-expert install \
  --agents unity-6000-expert,unity-graphics-expert \
  --provider claude \
  --dry-run \
  --force
```

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--agents` | CSV | all | Specific agents to install |
| `--provider` | string | opencode | Model provider (opencode, claude, gpt, gemini, custom) |
| `--dry-run` | bool | false | Preview without installing |
| `--force` | bool | false | Overwrite existing agents |
| `--dev` | bool | false | Install dev dependencies (future) |

### 8.3 Sync Flags

```bash
unity-expert sync \
  --agents unity-6000-expert \
  --provider claude \
  --dry-run
```

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--agents` | CSV | installed | Agents to sync |
| `--provider` | string | preserved | Update model provider |
| `--dry-run` | bool | false | Preview without syncing |

### 8.4 Directory Structure

```
unity-agent-expert/
├── cmd/unity-expert/
│   └── main.go              # Entry point
├── internal/
│   ├── app/
│   │   └── app.go           # Command routing
│   ├── cli/
│   │   ├── install.go       # Install command
│   │   ├── sync.go          # Sync command
│   │   └── uninstall.go     # Uninstall command
│   ├── model/
│   │   └── types.go         # Domain types
│   ├── pipeline/
│   │   └── orchestrator.go  # Pipeline execution
│   ├── backup/
│   │   └── snapshotter.go   # Backup management
│   ├── agents/
│   │   └── opencode.go      # OpenCode adapter
│   └── embed/
│       ├── config/
│       │   ├── agents.json
│       │   └── presets.json
│       └── prompts/
│           └── unity/*.md
├── scripts/
│   ├── install.ps1          # Windows bootstrapper
│   └── install.sh           # Unix bootstrapper
└── go.mod
```

### 8.5 Bootstrapper Flow (PS1)

```powershell
# 1. Check prerequisites (git, PowerShell version)
# 2. Clone/download repo to temp
# 3. Install Go binary (go install or download)
# 4. Call: unity-expert install --provider claude
# 5. Show status
```

---

## 9. Implementation Priority

### Phase 1: Core CLI (MVP)
1. ✅ `cmd/unity-expert/main.go` - Entry point
2. ✅ `internal/cli/install.go` - Install command with flags
3. ✅ `internal/agents/opencode.go` - OpenCode adapter
4. ✅ `internal/pipeline/orchestrator.go` - Simple pipeline
5. ✅ Embedded config files (agents.json, presets.json)

### Phase 2: Backup & Sync
6. ✅ `internal/backup/snapshotter.go` - Backup creation
7. ✅ `internal/cli/sync.go` - Sync command
8. ✅ `internal/state/state.go` - State persistence

### Phase 3: Polish
9. ✅ `internal/cli/uninstall.go` - Uninstall command
10. ✅ Dry-run rendering
11. ✅ Verification reports
12. ✅ Bootstrapper script updates

---

## 10. Key Learnings from gentle-ai

### ✅ Do This
1. **Use standard `flag` package** - cobra is overkill for simple CLI
2. **Separate CLI parsing from business logic** - `cli.ParseFlags()` returns struct, `RunXxx()` executes
3. **Pipeline with rollback** - Prepare stage (backup) → Apply stage (inject)
4. **State persistence** - `~/.unity-expert/state.json` for installed agents
5. **Sync command** - Idempotent re-apply without binary install
6. **Embedded filesystem** - `//go:embed` for config templates
7. **Dry-run mode** - Show plan without executing
8. **Backup before changes** - tar.gz with manifest

### ❌ Avoid This
1. **Don't mix TUI and CLI logic** - Keep them separate (gentle-ai does this well)
2. **Don't use cobra unless you need subcommands** - flag package is simpler
3. **Don't skip verification** - Always verify after install/sync
4. **Don't hardcode paths** - Use `os.UserHomeDir()` and `filepath.Join()`

---

## 11. Files to Reference

| File | Purpose |
|------|---------|
| `internal/cli/install.go` | Flag parsing, install execution |
| `internal/cli/sync.go` | Sync command (idempotent) |
| `internal/pipeline/orchestrator.go` | Stage execution |
| `internal/backup/snapshotter.go` | Backup creation |
| `internal/model/types.go` | Domain types |
| `internal/agents/` | Agent adapters |
| `scripts/install.ps1` | Windows bootstrapper |

---

## Next Recommended Steps

1. **Create Go module structure** - Initialize `go.mod` with cmd/internal layout
2. **Implement install command** - Start with flag parsing + JSON merge
3. **Add embedded configs** - Move agents.json, presets.json, prompts to embed
4. **Implement backup system** - Simple tar.gz before changes
5. **Add sync command** - Idempotent re-apply
6. **Update bootstrapper** - PS1 calls `unity-expert install` after binary install
