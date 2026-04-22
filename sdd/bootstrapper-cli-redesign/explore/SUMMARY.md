# Exploration Summary: gentle-ai CLI Architecture Analysis

**Status:** ✅ COMPLETE  
**Date:** 2026-04-21  
**Change:** bootstrapper-cli-redesign  
**Artifacts:** `sdd/bootstrapper-cli-redesign/explore/`

---

## Executive Summary

Successfully analyzed **gentle-ai's Go CLI architecture** to inform the redesign of unity-agent-expert's installer from PowerShell-only to a **Go-based CLI with PS1 bootstrapper**.

### Key Findings

1. **CLI Framework:** gentle-ai uses Go's standard `flag` package (NOT cobra/urfave) - simpler and sufficient for our needs
2. **Install vs Sync:** Two separate commands - `install` for first-time setup, `sync` for idempotent re-apply
3. **Pipeline Pattern:** Prepare (backup) → Apply (inject) → Verify with automatic rollback on failure
4. **State Persistence:** `~/.gentle-ai/state.json` tracks installed agents and model assignments
5. **Backup System:** Compressed tar.gz with manifest.json, auto-pruned to 5 most recent
6. **Embedded Configs:** Uses `//go:embed` for agents.json, presets.json, prompts, skills

---

## Artifacts Created

| File | Purpose | Lines |
|------|---------|-------|
| `cli-architecture-analysis.md` | Deep dive into gentle-ai's CLI implementation | ~600 |
| `architecture-proposal.md` | Proposed architecture for unity-expert CLI | ~800 |

**Total:** ~1400 lines of documentation

---

## Answers to Key Questions

### 1. How does the install command work internally?

**gentle-ai install flow:**
```
Parse Flags → Build StagePlan → Pipeline.Execute() → Verify
                   ↓
         PREPARE STAGE:
         • backup.Snapshot() → tar.gz + manifest.json
                   ↓
         APPLY STAGE:
         • componentApplyStep (per component)
         • agentInstallStep (per agent)
         • engram.Setup()
         • persona.Inject()
                   ↓
         VERIFY STAGE:
         • Check files exist
         • Validate JSON
```

### 2. What flags does it accept?

**gentle-ai install flags:**
```bash
gentle-ai install \
  --agents claude-code,opencode \
  --components sdd,engram,skills \
  --persona gentleman \
  --preset full-gentleman \
  --sdd-mode multi \
  --dry-run
```

**Proposed unity-expert flags:**
```bash
unity-expert install \
  --agents unity-6000-expert,unity-graphics-expert \
  --provider claude \
  --dry-run \
  --force
```

### 3. How does it handle file operations?

**JSON Merge (agents.json → opencode.json):**
```go
// Read existing config
config := readJSON("~/.config/opencode/opencode.json")

// Read embedded agents
agents := readEmbedded("config/agents.json")

// Merge (preserve existing, add new)
for name, agent := range agents {
    if presetModel, ok := presets.Models[name]; ok {
        agent.Model = presetModel
    }
    config.Agent[name] = agent
}

// Write back (UTF-8, no BOM)
writeJSON(config, "opencode.json")
```

**File Copy (prompts, skills):**
```go
//go:embed prompts/unity/*.md
var promptFS embed.FS

func copyPrompts(homeDir string) {
    entries, _ := promptFS.ReadDir("prompts/unity")
    for _, entry := range entries {
        src, _ := promptFS.Open(...)
        dst, _ := os.Create(filepath.Join(targetDir, entry.Name()))
        io.Copy(dst, src)
    }
}
```

### 4. How is the CLI structured?

**gentle-ai uses standard `flag` package:**
```go
// internal/cli/install.go
func ParseInstallFlags(args []string) (InstallFlags, error) {
    fs := flag.NewFlagSet("install", flag.ContinueOnError)
    
    // CSV list flags
    registerListFlag(fs, "agent", &opts.Agents)
    
    // String flags
    fs.StringVar(&opts.Persona, "persona", "", "persona to apply")
    
    // Bool flags
    fs.BoolVar(&opts.DryRun, "dry-run", false, "preview plan")
    
    fs.Parse(args)
    return opts, nil
}

// internal/app/app.go
func Run() error {
    switch args[0] {
    case "install":
        result, err := cli.RunInstall(args[1:])
        // Handle result
    case "sync":
        result, err := cli.RunSync(args[1:])
    }
}
```

**NOT cobra/urfave** - simpler approach with standard library.

---

## Comparison: PS1 vs Go CLI

| Feature | Current PS1 | Proposed Go CLI |
|---------|-------------|-----------------|
| Binary install | ❌ N/A | ✅ go install / Releases |
| JSON merge | ✅ PowerShell | ✅ Go encoding/json |
| File copy | ✅ Copy-Item | ✅ embed.FS + io.Copy |
| Model presets | ✅ presets.json | ✅ presets.json |
| Provider flag | ✅ `-Provider` | ✅ `--provider` |
| Dry run | ✅ `-DryRun` | ✅ `--dry-run` |
| Force reinstall | ✅ `-Force` | ✅ `--force` |
| **Backup** | ❌ No | ✅ **tar.gz + manifest** |
| **Rollback** | ❌ No | ✅ **Restore from backup** |
| **State persistence** | ❌ No | ✅ **state.json** |
| **Sync command** | ❌ No | ✅ **Idempotent re-apply** |
| **Uninstall** | ❌ No | ✅ **Clean removal** |

---

## Recommended Architecture for unity-expert

### Command Structure

```bash
unity-expert install [FLAGS]     # First-time installation
unity-expert sync [FLAGS]        # Re-apply config (idempotent)
unity-expert uninstall [FLAGS]   # Remove agents
unity-expert version             # Show version
unity-expert help                # Show help
```

### Directory Layout

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

### PS1 Bootstrapper Flow

```powershell
# 1. Check prerequisites (git, PowerShell 5.1+)
# 2. Clone/download repo
# 3. Install Go binary:
#    go install github.com/Ulysses-Alv/unity-agent-expert/cmd/unity-expert@latest
# 4. Call: unity-expert install --provider claude
# 5. Show status
```

---

## Implementation Priority

### Phase 1: Core CLI (Week 1)
- [ ] Initialize Go module
- [ ] Create cmd/unity-expert/main.go
- [ ] Implement internal/app/app.go (command routing)
- [ ] Implement internal/cli/install.go (flag parsing)
- [ ] Create internal/agents/opencode.go (JSON merge)
- [ ] Embed agents.json and presets.json

### Phase 2: File Operations (Week 2)
- [ ] Implement prompt file copying (embed.FS)
- [ ] Implement skill file copying
- [ ] Add model preset application
- [ ] Add --provider flag support
- [ ] Add --force flag support

### Phase 3: Backup & State (Week 3)
- [ ] Implement internal/backup/snapshotter.go
- [ ] Implement internal/state/state.go
- [ ] Add backup to install pipeline
- [ ] Add state persistence after install

### Phase 4: Sync Command (Week 4)
- [ ] Implement internal/cli/sync.go
- [ ] Implement idempotent sync logic
- [ ] Add state loading for sync
- [ ] Add --dry-run for sync

### Phase 5: Polish (Week 5)
- [ ] Implement internal/cli/uninstall.go
- [ ] Add verification reports
- [ ] Update install.ps1 bootstrapper
- [ ] E2E testing

---

## Key Learnings

### ✅ Patterns to Adopt

1. **Standard `flag` package** - cobra is overkill for simple CLI
2. **Separate install vs sync** - install for first-time, sync for re-apply
3. **Pipeline with rollback** - backup before changes, restore on failure
4. **State persistence** - track installed agents and model assignments
5. **Embedded filesystem** - `//go:embed` for config templates
6. **Dry-run mode** - show plan without executing
7. **Clean separation** - cli.ParseFlags() → cli.RunXxx() → pipeline.Execute()

### ❌ Anti-patterns to Avoid

1. **Don't mix TUI and CLI logic** - keep them separate
2. **Don't skip verification** - always verify after install/sync
3. **Don't hardcode paths** - use os.UserHomeDir() and filepath.Join()
4. **Don't overwrite without backup** - always snapshot first

---

## Next Recommended Steps

1. **Create Go module structure** - `go mod init github.com/Ulysses-Alv/unity-agent-expert`
2. **Implement install command** - Start with flag parsing + JSON merge
3. **Add embedded configs** - Move agents.json, presets.json, prompts to embed
4. **Implement backup system** - Simple tar.gz before changes
5. **Add sync command** - Idempotent re-apply
6. **Update bootstrapper** - PS1 calls `unity-expert install` after binary install

---

## Files to Reference

| File | Purpose |
|------|---------|
| `cli-architecture-analysis.md` | Deep dive into gentle-ai's implementation |
| `architecture-proposal.md` | Full architecture proposal with diagrams |
| `D:\Projects\Skill creation\install.ps1` | Current PS1 installer (to refactor) |
| `D:\Projects\Skill creation\opencode\config\agents.json` | Agent definitions |
| `D:\Projects\Skill creation\opencode\config\presets.json` | Model presets |

---

**Exploration Status:** ✅ COMPLETE  
**Ready for:** Design phase (sdd-design)  
**Confidence Level:** HIGH - gentle-ai provides proven reference implementation
