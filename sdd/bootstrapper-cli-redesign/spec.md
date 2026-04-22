# Bootstrapper CLI Redesign — Specification

## Purpose

Define the detailed requirements for replacing the 384-line PowerShell monolith (`install.ps1`) with a Go CLI binary (`unity-expert.exe`) plus a ~20-line PS1 bootstrapper. The CLI provides `install`, `sync`, `doctor`, and `version` commands with backup, rollback, state persistence, and dry-run capabilities.

---

## Domain: CLI Commands

### Requirement: Install Command

The system MUST provide an `install` command that performs first-time installation of Unity agents, prompts, and skills into the user's OpenCode configuration.

**Flags:**

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--provider` | string | `opencode` | Model provider preset: `gpt`, `claude`, `gemini`, `opencode`, `custom` |
| `--dry-run` | bool | `false` | Preview changes without applying |
| `--force` | bool | `false` | Overwrite existing Unity agents in opencode.json |
| `--agents` | CSV | all | Comma-separated list of specific agents to install (e.g., `unity-6000-expert,unity-graphics-expert`) |

**Preconditions:**
- `opencode.json` MUST exist at `~/.config/opencode/opencode.json`
- Embedded `agents.json` and `presets.json` MUST be present in the binary
- Embedded prompts (`prompts/unity/*.md`) and skills (`skills/unity-6000/*`) MUST be present

**Behavior:**
1. Parse flags and validate `--provider` against known presets
2. Check for existing Unity agents in `opencode.json` (agents matching `unity-*`)
3. If agents exist and `--force` is NOT set → skip with informational message
4. Create backup snapshot of `opencode.json` (tar.gz + manifest)
5. Merge agents from embedded `agents.json` into `opencode.json`
6. Apply model assignments from the selected provider's preset in `presets.json`
7. Copy embedded prompts to `~/.config/opencode/prompts/unity/`
8. Copy embedded skills to `~/.config/opencode/skills/unity-6000/`
9. Write state to `~/.unity-expert/state.json`
10. Verify all files exist and JSON is valid
11. On verification failure → rollback from backup
12. Print success report with installed components

#### Scenario: Fresh install with Claude provider

- GIVEN `opencode.json` exists at `~/.config/opencode/opencode.json` with no `unity-*` agents
- WHEN user runs `unity-expert install --provider claude`
- THEN all 18 agents are merged into `opencode.json` with `anthropic/claude-sonnet-4-20250514` model
- AND 18 prompt files are copied to `~/.config/opencode/prompts/unity/`
- AND 19 skill folders are copied to `~/.config/opencode/skills/unity-6000/`
- AND `state.json` is written to `~/.unity-expert/state.json`
- AND a backup of the original `opencode.json` is created

#### Scenario: Install with existing agents (no force)

- GIVEN `opencode.json` already contains 5 `unity-*` agents
- WHEN user runs `unity-expert install --provider gemini` (without `--force`)
- THEN the command prints: `[INFO] Unity agents already installed (5 found). Use --force to reinstall.`
- AND exits with code 0 without modifying any files

#### Scenario: Force reinstall

- GIVEN `opencode.json` already contains 5 `unity-*` agents
- WHEN user runs `unity-expert install --provider gpt --force`
- THEN a backup of the current `opencode.json` is created
- AND all 18 agents are re-merged with `openai/gpt-4o` / `openai/gpt-4o-mini` models
- AND prompts and skills are re-copied (overwriting existing files)
- AND `state.json` is updated

#### Scenario: Dry-run preview

- GIVEN `opencode.json` exists with no `unity-*` agents
- WHEN user runs `unity-expert install --provider claude --dry-run`
- THEN no files are modified
- AND output shows: `[DRY RUN] Would install 18 agent(s) with claude preset`
- AND output shows: `[DRY RUN] Would copy 18 prompt(s) to ~/.config/opencode/prompts/unity/`
- AND output shows: `[DRY RUN] Would copy 19 skill(s) to ~/.config/opencode/skills/unity-6000/`
- AND exit code is 0

#### Scenario: Install specific agents only

- GIVEN `opencode.json` exists with no `unity-*` agents
- WHEN user runs `unity-expert install --provider claude --agents unity-6000-expert,unity-graphics-expert`
- THEN only `unity-6000-expert` and `unity-graphics-expert` are added to `opencode.json`
- AND all prompts and skills are still copied (full set)
- AND `state.json` records only the 2 installed agents

#### Scenario: Missing opencode.json precondition

- GIVEN `~/.config/opencode/opencode.json` does NOT exist
- WHEN user runs `unity-expert install --provider claude`
- THEN the command prints: `[ERROR] OpenCode config not found at: ~/.config/opencode/opencode.json`
- AND prints: `Please ensure OpenCode is installed. Visit: https://opencode.ai`
- AND exits with code 1

---

### Requirement: Sync Command

The system MUST provide a `sync` command that idempotently re-applies configuration without reinstalling binaries or re-checking preconditions.

**Flags:**

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--provider` | string | preserved from state | Update model provider preset |
| `--dry-run` | bool | `false` | Preview changes without applying |
| `--force` | bool | `false` | Force overwrite even if no drift detected |

**Behavior:**
1. Load `state.json` to discover previously installed agents
2. If no state found, scan `opencode.json` for existing `unity-*` agents
3. If no agents found at all → print `[INFO] Nothing to sync. Run 'install' first.` and exit 0
4. Create backup snapshot
5. Re-merge agent configs (preserving user model assignments unless `--provider` is specified)
6. Re-copy prompts (only files that differ from embedded versions)
7. Re-copy skills (only directories that differ from embedded versions)
8. Update `state.json` with current state
9. Verify and report

**Diff from install:**
- Sync does NOT check `opencode.json` precondition (assumes it exists from prior install)
- Sync does NOT install agents that are not already in state/opencode.json
- Sync preserves user model assignments unless `--provider` overrides
- Sync detects drift and reports "no changes needed" if everything matches

#### Scenario: Sync with no drift

- GIVEN `state.json` shows 18 agents installed with claude preset
- AND `opencode.json` matches the expected state (no user modifications)
- WHEN user runs `unity-expert sync`
- THEN backup is created
- AND merge detects no changes needed
- AND output shows: `All managed assets are up to date. No files changed.`
- AND exit code is 0

#### Scenario: Sync with provider change

- GIVEN `state.json` shows 18 agents installed with `opencode` preset
- WHEN user runs `unity-expert sync --provider claude`
- THEN all 18 agents have their models updated to `anthropic/claude-sonnet-4-20250514`
- AND `state.json` is updated with new provider and model assignments
- AND output shows: `[OK] Updated 18 agent model(s) to claude preset`

#### Scenario: Sync with user-modified models (no provider flag)

- GIVEN `state.json` shows `unity-6000-expert` was installed with `opencode` preset
- AND user manually changed `unity-6000-expert` model to `anthropic/claude-sonnet-4-20250514` in `opencode.json`
- WHEN user runs `unity-expert sync` (no `--provider`)
- THEN the user's manual model assignment is PRESERVED (not overwritten by preset)
- AND output shows: `Preserved user model assignments for 1 agent(s)`

#### Scenario: Sync dry-run

- GIVEN `state.json` shows 18 agents installed
- AND `opencode.json` has 3 agents with drifted model assignments
- WHEN user runs `unity-expert sync --dry-run`
- THEN no files are modified
- AND output shows: `[DRY RUN] Would update 3 agent model(s)`
- AND output shows: `[DRY RUN] Would copy 2 prompt(s) (changed from embedded)`
- AND exit code is 0

#### Scenario: Sync with nothing installed

- GIVEN no `state.json` exists and no `unity-*` agents in `opencode.json`
- WHEN user runs `unity-expert sync`
- THEN output shows: `[INFO] Nothing to sync. Run 'install' first.`
- AND exit code is 0

---

### Requirement: Doctor Command

The system MUST provide a `doctor` command that validates the health of the Unity Agent Expert installation.

**Flags:** None

**Behavior:**
1. Check `opencode.json` exists at `~/.config/opencode/opencode.json`
2. Check Unity agents are present in `opencode.json` (agents matching `unity-*`)
3. Check prompt files exist at `~/.config/opencode/prompts/unity/`
4. Check skill directories exist at `~/.config/opencode/skills/unity-6000/`
5. Check `state.json` exists at `~/.unity-expert/state.json`
6. Validate JSON syntax of `opencode.json`
7. Print PASS/FAIL for each check with details
8. Exit with code 0 if all checks pass, code 1 if any fail

**Output format:**
```
Unity Expert Doctor — Installation Health Check
================================================

[PASS] opencode.json exists at ~/.config/opencode/opencode.json
[PASS] 18 Unity agent(s) found in opencode.json
[PASS] 18 prompt file(s) present in ~/.config/opencode/prompts/unity/
[PASS] 19 skill directory(ies) present in ~/.config/opencode/skills/unity-6000/
[PASS] state.json found at ~/.unity-expert/state.json
[PASS] opencode.json JSON syntax is valid

All 6 checks passed. Installation is healthy.
```

#### Scenario: All checks pass

- GIVEN a complete, valid installation
- WHEN user runs `unity-expert doctor`
- THEN all 6 checks show `[PASS]`
- AND output shows: `All 6 checks passed. Installation is healthy.`
- AND exit code is 0

#### Scenario: Missing opencode.json

- GIVEN `~/.config/opencode/opencode.json` does NOT exist
- WHEN user runs `unity-expert doctor`
- THEN the opencode.json check shows: `[FAIL] opencode.json not found at ~/.config/opencode/opencode.json`
- AND agent check shows: `[FAIL] Cannot check agents — opencode.json missing`
- AND exit code is 1

#### Scenario: Partial installation (agents but no prompts)

- GIVEN `opencode.json` has 18 Unity agents but `prompts/unity/` directory is empty
- WHEN user runs `unity-expert doctor`
- THEN agent check shows `[PASS]`
- AND prompt check shows: `[FAIL] 0 prompt file(s) found (expected 18)`
- AND exit code is 1

#### Scenario: Corrupted opencode.json

- GIVEN `opencode.json` exists but contains invalid JSON
- WHEN user runs `unity-expert doctor`
- THEN the JSON validation check shows: `[FAIL] opencode.json has invalid JSON syntax: <error detail>`
- AND exit code is 1

#### Scenario: Missing state.json (non-fatal)

- GIVEN a valid installation but `~/.unity-expert/state.json` was deleted
- WHEN user runs `unity-expert doctor`
- THEN the state.json check shows: `[WARN] state.json not found (non-fatal, run 'sync' to regenerate)`
- AND all other checks pass
- AND exit code is 0 (WARN is not a failure)

---

### Requirement: Version Command

The system MUST provide a `version` command that displays the CLI version string.

**Flags:** None

**Behavior:**
1. Print version string in format: `unity-expert v{version}`
2. Version is injected at build time via `-ldflags` (defaults to `dev` if not set)
3. Exit with code 0

#### Scenario: Release version

- GIVEN the binary was built with `-ldflags "-X main.version=1.0.0"`
- WHEN user runs `unity-expert version`
- THEN output shows: `unity-expert v1.0.0`
- AND exit code is 0

#### Scenario: Development version

- GIVEN the binary was built without version injection
- WHEN user runs `unity-expert version`
- THEN output shows: `unity-expert vdev`
- AND exit code is 0

---

## Domain: PS1 Bootstrapper

### Requirement: Minimal PS1 Bootstrapper

The system MUST provide a PowerShell bootstrapper (`install.ps1`) of approximately 20 lines that downloads the Go binary and delegates all installation logic to it.

**Responsibilities:**
1. Check prerequisites: PowerShell 5.1+ available
2. Determine install location: `$env:LOCALAPPDATA\unity-expert\bin\`
3. Check if `unity-expert.exe` already exists at install location
4. If not present, download `unity-expert.exe` from GitHub Releases
5. Copy/download to `$env:LOCALAPPDATA\unity-expert\bin\unity-expert.exe`
6. Call `unity-expert.exe` with all passed arguments forwarded

**Script signature:**
```powershell
# Usage: irm <url>/install.ps1 | iex
# With args: irm <url>/install.ps1 | iex -ArgumentList "install --provider claude"
```

**Behavior:**
- Accepts a single string argument containing CLI flags (e.g., `"install --provider claude"`)
- If no argument provided, defaults to `"install"`
- Detects existing installation → defaults to `"sync"` instead of `"install"`
- Downloads binary for Windows AMD64 from GitHub Releases
- Validates download with SHA256 checksum (if available)
- Creates `$env:LOCALAPPDATA\unity-expert\bin\` directory if not exists
- Executes `unity-expert.exe` with forwarded arguments
- Propagates exit code from `unity-expert.exe`

#### Scenario: Fresh install via one-liner

- GIVEN user has PowerShell 5.1+ and no prior installation
- WHEN user runs `irm <url>/install.ps1 | iex`
- THEN bootstrapper downloads `unity-expert.exe` to `$env:LOCALAPPDATA\unity-expert\bin\`
- AND runs `unity-expert.exe install` (default command)
- AND displays install output from the Go CLI

#### Scenario: Install with provider flag

- GIVEN user has PowerShell 5.1+ and no prior installation
- WHEN user runs `irm <url>/install.ps1 | iex -ArgumentList "install --provider claude"`
- THEN bootstrapper downloads `unity-expert.exe`
- AND runs `unity-expert.exe install --provider claude`
- AND all 18 agents are installed with Claude models

#### Scenario: Re-run on existing installation

- GIVEN `unity-expert.exe` already exists at `$env:LOCALAPPDATA\unity-expert\bin\`
- WHEN user runs `irm <url>/install.ps1 | iex` (no arguments)
- THEN bootstrapper detects existing binary
- AND runs `unity-expert.exe sync` (default for existing installs)
- AND does NOT re-download the binary

#### Scenario: Download failure

- GIVEN user has no internet connection or GitHub is unreachable
- WHEN user runs `irm <url>/install.ps1 | iex`
- THEN bootstrapper prints: `[ERROR] Failed to download unity-expert.exe`
- AND prints: `Check your internet connection and try again.`
- AND exits with code 1

---

## Domain: Error Handling

### Requirement: Error Classification and Exit Codes

The system MUST classify errors and report them with appropriate exit codes.

**Exit codes:**

| Code | Meaning |
|------|---------|
| `0` | Success (all operations completed) |
| `1` | General error (precondition failure, invalid flags, I/O error) |
| `2` | Backup/rollback failure (critical — user should investigate) |

**Error categories:**

| Category | Examples | Exit Code | User Message |
|----------|----------|-----------|--------------|
| **Precondition failure** | `opencode.json` missing, invalid provider value | 1 | `[ERROR] <description>` + remediation hint |
| **Invalid flags** | Unknown `--provider` value, malformed `--agents` CSV | 1 | `[ERROR] Invalid flag: <flag>. Run 'unity-expert help' for usage.` |
| **I/O error** | Cannot read/write `opencode.json`, permission denied | 1 | `[ERROR] File operation failed: <path>. <detail>` |
| **JSON parse error** | `opencode.json` contains invalid JSON | 1 | `[ERROR] Invalid JSON in <path>: <detail>` |
| **Backup failure** | Cannot create backup directory, disk full | 2 | `[ERROR] Backup failed. Aborting to prevent data loss.` |
| **Rollback failure** | Cannot restore from backup after failure | 2 | `[CRITICAL] Rollback failed. Manual recovery may be needed. Backup at: <path>` |
| **Network error** | Download fails in PS1 bootstrapper | 1 | `[ERROR] Download failed. Check internet connection.` |
| **Verification failure** | Installed files missing after apply | 1 | `[ERROR] Verification failed. Restored from backup.` |

**Error output format:**
```
[ERROR] <category>: <description>
[ERROR] Detail: <technical detail>
[ERROR] Hint: <remediation suggestion>
```

**Dry-run error behavior:**
- In dry-run mode, errors during the "plan" phase (flag validation, preset lookup) MUST still be reported with appropriate exit codes
- Errors during the "apply" phase do NOT occur in dry-run (nothing is applied)

#### Scenario: Invalid provider value

- GIVEN user runs `unity-expert install --provider invalid-provider`
- THEN output shows: `[ERROR] Invalid flag: --provider. Value 'invalid-provider' is not a known preset.`
- AND output shows: `Known providers: opencode, claude, gpt, gemini, custom`
- AND exit code is 1

#### Scenario: Backup failure aborts install

- GIVEN disk is full and backup cannot be created
- WHEN user runs `unity-expert install --provider claude`
- THEN output shows: `[ERROR] Backup failed. Aborting to prevent data loss.`
- AND NO changes are made to `opencode.json`
- AND exit code is 2

#### Scenario: Rollback after partial failure

- GIVEN agent merge succeeds but prompt copy fails (permission denied)
- WHEN user runs `unity-expert install --provider claude`
- THEN rollback restores `opencode.json` from backup
- AND output shows: `[ERROR] Prompt copy failed: permission denied`
- AND output shows: `[OK] Restored opencode.json from backup`
- AND exit code is 1

---

## Domain: State Persistence

### Requirement: State File Schema

The system MUST persist installation state to `~/.unity-expert/state.json` after every successful `install` or `sync` operation.

**Schema:**
```json
{
  "version": 1,
  "installedAt": "2026-04-21T23:30:00Z",
  "updatedAt": "2026-04-21T23:30:00Z",
  "installedAgents": ["unity-6000-expert", "unity-graphics-expert", "..."],
  "provider": "claude",
  "modelAssignments": {
    "unity-6000-expert": "anthropic/claude-sonnet-4-20250514",
    "unity-graphics-expert": "anthropic/claude-sonnet-4-20250514"
  }
}
```

**Behavior:**
- `installedAt` is set on first install, never modified by sync
- `updatedAt` is set on every install and sync
- `installedAgents` reflects the agents currently managed by the CLI
- `provider` reflects the last used provider preset
- `modelAssignments` reflects the current model for each agent (may include user overrides after sync)
- State file is written atomically (write to temp file, then rename)
- If state file is corrupted or missing, the CLI recovers by scanning `opencode.json`

#### Scenario: State recovery from opencode.json

- GIVEN `state.json` was deleted but `opencode.json` has 10 `unity-*` agents
- WHEN user runs `unity-expert sync`
- THEN the CLI scans `opencode.json` and discovers the 10 agents
- AND performs sync on those 10 agents
- AND writes a new `state.json` with discovered agents

---

## Domain: Backup System

### Requirement: Backup Before Mutation

The system MUST create a backup of `opencode.json` before any `install` or `sync` operation that modifies files.

**Behavior:**
- Backup location: `~/.unity-expert/backups/<timestamp>/`
- Backup contents: `opencode.json` copy + `manifest.json`
- Manifest includes: timestamp, source command (`install`/`sync`), CLI version, file paths
- Backups are auto-pruned to keep the 5 most recent
- Backup is NOT created in dry-run mode
- If backup creation fails → abort the entire operation (exit code 2)

#### Scenario: Backup auto-pruning

- GIVEN 5 existing backups in `~/.unity-expert/backups/`
- WHEN user runs `unity-expert install --provider claude`
- THEN a 6th backup is created
- AND the oldest backup is deleted
- AND exactly 5 backups remain

---

## Coverage Summary

| Domain | Requirements | Scenarios |
|--------|-------------|-----------|
| CLI Commands — Install | 1 | 6 |
| CLI Commands — Sync | 1 | 5 |
| CLI Commands — Doctor | 1 | 5 |
| CLI Commands — Version | 1 | 2 |
| PS1 Bootstrapper | 1 | 4 |
| Error Handling | 1 | 3 |
| State Persistence | 1 | 1 |
| Backup System | 1 | 1 |
| **Total** | **8** | **27** |

### Coverage Assessment
- **Happy paths:** ✅ All commands have primary success scenarios
- **Edge cases:** ✅ Existing installs, missing files, corrupted state, partial failures
- **Error states:** ✅ Precondition failures, I/O errors, backup failures, rollback scenarios
- **Dry-run:** ✅ Covered for install and sync commands
