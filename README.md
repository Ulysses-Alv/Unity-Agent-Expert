# Unity 6000.3 LTS Expert Agent System

AI agents specialized in Unity 6000.3 LTS development. Built on OpenCode with an orchestrator + specialized sub-agents architecture.

## Quick Installation

### One-Liner (recommended)

```bash
# Linux / macOS
curl -fsSL https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.sh | bash
```
```bash
# Windows PowerShell
irm https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/install.ps1 | iex
```

This installs the `unity-agent-expert` CLI to `~/.local/bin` (Linux/macOS) or `%LOCALAPPDATA%\unity-agent-expert\bin` (Windows).

### Add to PATH

After installation, add the binary to your PATH:

**Linux/macOS:**
```bash
export PATH="$PATH:$HOME/.local/bin"
# Add to ~/.bashrc or ~/.zshrc to persist
```

**Windows PowerShell:**
```powershell
[Environment]::SetEnvironmentVariable('PATH', $env:PATH + ';$env:LOCALAPPDATA\unity-agent-expert\bin', 'User')
```

### First Run

```bash
unity-agent-expert --help
```

## CLI Commands

### `unity-agent-expert install`

Installs Unity agents, prompts, and skills into the OpenCode configuration.

```bash
unity-agent-expert install --provider gpt         # Install with GPT models
unity-agent-expert install --provider claude     # Install with Claude models
unity-agent-expert install --dry-run             # Preview changes
unity-agent-expert install --force               # Overwrite existing
```

### `unity-agent-expert sync`

Idempotently re-applies configuration (updates model assignments, etc.).

```bash
unity-agent-expert sync                            # Re-apply last provider
unity-agent-expert sync --provider claude         # Change to Claude models
```

### `unity-agent-expert doctor`

Checks installation health.

```bash
unity-agent-expert doctor
```

### `unity-agent-expert version`

Shows version information.

```bash
unity-agent-expert version
```

## Provider Model Mapping

| Provider | Models Used |
|----------|-------------|
| `opencode` | MiniMax M2.7 for general tasks; MiMo-V2-Pro for graphics/animation/VFX |
| `claude` | Claude Sonnet 4 for all agents |
| `gpt` | GPT-4o for complex tasks; GPT-4o-mini for lighter tasks |
| `gemini` | Gemini 2.5 Pro for complex tasks; Gemini 2.5 Flash for faster responses |
| `custom` | User-defined models in presets.json |

## Repository Structure

```
.
├── unity-expert/                  # Go CLI source
│   ├── cmd/unity-expert/          # Main entry point
│   ├── internal/                  # Internal packages
│   │   ├── cli/                   # CLI commands (install, sync, doctor)
│   │   ├── config/                # OpenCode config handling
│   │   ├── install/               # Install pipeline
│   │   ├── backup/                # Backup system (tar.gz)
│   │   └── embed/                 # Embedded assets
│   ├── prompts/                   # Embedded prompts
│   ├── skills/                    # Embedded skills
│   └── .goreleaser.yaml           # Multi-platform build config
│
├── install.ps1                    # Bootstrapper installer (root)
├── scripts/
│   ├── install.sh                 # Linux/macOS bootstrapper
│   ├── install.ps1                # Legacy installer (full logic)
│   └── install-legacy.ps1         # Legacy installer (full logic, backup)
│
├── opencode/                      # OpenCode configuration
│   ├── config/
│   │   └── agents.json            # Agent definitions
│   └── prompts/unity/             # Prompts for each agent
│
└── skills/                        # Knowledge skills
    ├── unity-3d-graphics/         # URP, shaders, lighting
    ├── unity-2d/                  # Sprites, tilemaps, 2D physics
    ├── unity-physics/             # Colliders, joints, rigidbody
    ├── unity-scripting/           # MonoBehaviour, C#, Jobs
    ├── unity-animation/           # Animator, Timeline, Playables
    ├── unity-audio/               # AudioSource, AudioMixer
    ├── unity-editor/               # EditorWindow, Custom Inspectors
    ├── unity-performance/         # Profiling, optimization
    ├── unity-build-deploy/         # BuildPlayerOptions, IL2CPP
    ├── unity-vfx/                 # Particle Systems, VFX Graph
    ├── unity-input/                # Input System (not legacy)
    ├── unity-xr/                  # AR Foundation, VR
    ├── unity-ui-*                  # UI Toolkit (UXML, USS, C#)
    ├── unity-cinemachine/          # Virtual cameras, Timeline
    ├── unity-addressables/        # Async loading, AssetBundles
    ├── unity-packages/             # Package Manager
    └── unity-shadergraph/          # Shader Graph, nodes, Master Stack
```

## Architecture

The system uses a **bootstrapper + CLI** architecture:

```
┌─────────────────────────────────────────────────────────────┐
│  install.ps1 / install.sh (bootstrapper)                   │
│  Downloads unity-agent-expert binary                        │
│  Installs to ~/.local/bin or %LOCALAPPDATA%\                 │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────┐
│  unity-agent-expert (CLI)                                    │
│  ├── install  → Installs agents, prompts, skills            │
│  ├── sync     → Idempotent re-apply                         │
│  ├── doctor   → Health checks                               │
│  └── version  → Version info                                │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────┐
│  ~/.config/opencode/                                         │
│  ├── opencode.json           # Modified with Unity agents    │
│  ├── prompts/unity/          # Copied prompts               │
│  └── skills/unity-6000/     # Copied skills                │
└─────────────────────────────────────────────────────────────┘
```

## Updating

```bash
# Update the CLI
unity-agent-expert update

# Or re-run the bootstrapper
irm https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.ps1 | iex
```

## Backup & Rollback

Every `install` automatically creates a backup:

- Location: `~/.unity-expert/backups/`
- Format: `tar.gz` with `manifest.json`
- Auto-prune: keeps 5 most recent

To rollback, use the restore command or restore from a specific backup.

## Requirements

- OpenCode installed
- PowerShell 5.1+ (Windows) or bash (Linux/macOS)
- Access to `~/.config/opencode/`

## Contributing

Contributions are welcome! Here's how you can help:

### Ways to Contribute

- **Report bugs** — Open an issue with reproduction steps
- **Suggest features** — Open an issue with your proposal
- **Improve documentation** — Skills, prompts, or README improvements
- **Add new skills** — Extend Unity domain coverage
- **Submit code** — PRs for bug fixes or new features

### Development Setup

```bash
# Clone the repo
git clone https://github.com/Ulysses-Alv/Unity-Agent-Expert.git
cd Unity-Agent-Expert

# Prerequisites
- Go 1.21+
- goreleaser (for releases)

# Build locally
cd unity-expert
go build -ldflags "-X main.version=1.0.0" -o unity-agent-expert.exe ./cmd/unity-expert

# Run tests
go test ./...

# Build release
goreleaser build --clean
```

### Creating Releases

Releases are automated with goreleaser. To create a release:

```bash
# Ensure GITHUB_TOKEN is set
export GITHUB_TOKEN="your-github-token"

# Create and push version tag
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0

# Run goreleaser
cd unity-expert
goreleaser release --clean
```

### Code Standards

- Write clear commit messages (conventional commits preferred)
- Add tests for new functionality
- Run `go fmt` before committing
- Update this README if you add new features

## License

MIT License

---

**⚠️ Disclaimer: Everything was vibecodeado**

This project was built with AI-assisted coding (vibe coding). Every line, every decision, every architectural choice was made with AI agents — from the initial exploration to the final implementation.

**Philosophy:** If it works, ship it. If it's ugly, make it work first. Beauty comes later.

The goal was to solve a real problem (Unity agent installation) not to write perfect code. The result is pragmatic, functional, and documented.

If you want to learn how this was built, explore the git history. If you want to build something similar, use the same approach: define what you want, let AI do the work, verify the result.

**Recommended reading:**
- [gentle-ai/gentle-ai](https://github.com/Gentleman-Programming/gentle-ai) — inspiration for the bootstrapper + CLI architecture
- [SDD Workflow](https://github.com/gentle-ai/gentle-ai/blob/main/AGENTS.md) — how to use AI agents for software development
- [Agent Teams Lite](https://github.com/gentle-ai/agent-teams-lite) — the orchestration system used
