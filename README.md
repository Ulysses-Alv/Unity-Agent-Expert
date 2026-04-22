# Unity 6000.3 LTS Expert Agent System

AI agents specialized in Unity 6000.3 LTS development. Built on OpenCode with an orchestrator + specialized sub-agents architecture.

## Quick Installation

### One-Liner (recommended)

```bash
# Linux / macOS
curl -fsSL https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.sh | bash

# Windows PowerShell
irm https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.ps1 | iex
```

### Installation Flags

| Flag | Description |
|------|-------------|
| `--dry-run` / `-DryRun` | Preview what will be installed without making changes |
| `--force` / `-Force` | Overwrite existing agents (default: skip if already installed) |
| `--dev` / `-Dev` | Install development dependencies |
| `--provider <name>` / `-Provider <name>` | Set model provider. Values: `opencode`, `claude`, `gpt`, `gemini`, `custom` |

#### Provider Model Mapping

| Provider | Models Used |
|----------|-------------|
| `opencode` | MiniMax M2.7 for general tasks; MiMo-V2-Pro for graphics/animation/VFX |
| `claude` | Claude Sonnet 4 for all agents |
| `gpt` | GPT-4o for complex tasks; GPT-4o-mini for lighter tasks |
| `gemini` | Gemini 2.5 Pro for complex tasks; Gemini 2.5 Flash for faster responses |
| `custom` | User-defined models in presets.json |

### Examples

```bash
# Linux/macOS
curl -fsSL https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.sh | bash -s -- --dry-run
curl -fsSL https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.sh | bash -s -- --provider claude
```

```powershell
# Windows PowerShell - parameters BEFORE the pipe
irm https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.ps1 -DryRun | iex
irm https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.ps1 -Force | iex
irm https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.ps1 -Provider gpt | iex
```

## Repository Structure

```
.
├── opencode/                    # OpenCode configuration
│   ├── config/
│   │   └── agents.json         # Agent definitions
│   ├── prompts/unity/         # Prompts for each agent
│   │   ├── unity-6000-expert.md # Orchestrator (delegates to sub-agents)
│   │   └── unity-*-expert.md   # 18 specialized sub-agents
│   └── install.ps1             # Local installation script
│
├── scripts/
│   ├── install.sh              # One-liner for Linux/macOS
│   └── install.ps1             # One-liner for Windows
│
├── skills/                      # Knowledge skills
│   ├── unity-3d-graphics/       # URP, shaders, lighting
│   ├── unity-2d/                # Sprites, tilemaps, 2D physics
│   ├── unity-physics/           # Colliders, joints, rigidbody
│   ├── unity-scripting/         # MonoBehaviour, C#, Jobs
│   ├── unity-animation/          # Animator, Timeline, Playables
│   ├── unity-audio/             # AudioSource, AudioMixer
│   ├── unity-editor/            # EditorWindow, Custom Inspectors
│   ├── unity-performance/       # Profiling, optimization
│   ├── unity-build-deploy/      # BuildPlayerOptions, IL2CPP
│   ├── unity-vfx/               # Particle Systems, VFX Graph
│   ├── unity-input/             # Input System (not legacy)
│   ├── unity-xr/                # AR Foundation, VR
│   ├── unity-ui-*               # UI Toolkit (UXML, USS, C#)
│   ├── unity-cinemachine/       # Virtual cameras, Timeline
│   ├── unity-addressables/      # Async loading, AssetBundles
│   ├── unity-packages/          # Package Manager
│   └── unity-shadergraph/       # Shader Graph, nodes, Master Stack
│
├── agents/                      # Standalone agents (backup)
├── Manual/                      # Unity Manual converted to MD (3435 files)
└── scripts/                     # One-liner installers
```

## Local Installation

```bash
# Clone the repo
git clone https://github.com/Ulysses-Alv/Unity-Agent-Expert.git
cd Unity-Agent-Expert

# Linux/macOS
chmod +x install.sh
./install.sh --dry-run   # Preview
./install.sh             # Install

# Windows PowerShell
.\install.ps1 -DryRun    # Preview
.\install.ps1            # Install
```

## Usage

### Activate Unity Orchestrator

```
/agent unity-6000-expert
```

The orchestrator detects the problem domain and delegates to the appropriate specialized sub-agent.

### Available Domains

| Domain | Agent |
|--------|-------|
| UI Toolkit | unity-ui-expert |
| Shader Graph | unity-shadergraph-expert |
| 3D Graphics | unity-graphics-expert |
| 2D Graphics | unity-2d-expert |
| Physics | unity-physics-expert |
| Scripting | unity-scripting-expert |
| Animation | unity-animation-expert |
| Audio | unity-audio-expert |
| Editor | unity-editor-expert |
| Performance | unity-performance-expert |
| Build & Deploy | unity-build-expert |
| VFX | unity-vfx-expert |
| Input | unity-input-expert |
| XR (AR/VR) | unity-xr-expert |
| Cinemachine | unity-cinemachine-expert |
| Addressables | unity-addressables-expert |
| Packages | unity-packages-expert |

## Architecture

```
┌─────────────────────────────────────────────────────┐
│                   USER                               │
└─────────────────────┬───────────────────────────────┘
                      │
                      ▼
           ┌────────────────────────┐
           │  unity-6000-expert     │  ← Orchestrator
           │  (detects domain)       │
           └────────────┬────────────┘
                        │ async delegation
     ┌─────────────────┼─────────────────┐
     │                 │                 │
     ▼                 ▼                 ▼
┌─────────┐     ┌───────────┐    ┌──────────┐
│unity-   │     │unity-    │    │unity-    │
│graphics │     │physics   │    │scripting │
│-expert  │     │-expert   │    │-expert   │
└─────────┘     └───────────┘    └──────────┘
```

## Key Principles

1. **Async delegation** — The orchestrator NEVER blocks. Uses `delegate` (async), not `task` (sync).
2. **Updatable skills** — Sub-agents read skills from the filesystem; no hardcoded knowledge.
3. **Resilience** — If config gets overwritten, just run `install.ps1` to restore it.
4. **Idempotency** — Install scripts can run multiple times without duplicating agents.

## Updating

```bash
# One-liner (simple)
curl -fsSL https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.sh | bash

# Local (after git pull)
cd Unity-Agent-Expert
./install.sh
# or
.\install.ps1
```

## Backup Files

If the install script modifies opencode.json, it automatically creates a backup:

```
~/.config/opencode/opencode.json.backup.YYYYMMDDXXXXXX
```

## Requirements

- OpenCode installed
- PowerShell 5.1+ (Windows) or bash (Linux/macOS)
- Access to `~/.config/opencode/`

## License

MIT Licence
