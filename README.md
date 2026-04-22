# Unity 6000.3 LTS Expert Agent System

Sistema de agentes IA especializados en Unity 6000.3 LTS. Construido sobre OpenCode con arquitectura de orquestador + sub-agentes especializados.

## Quick Installation

### One-Liner (recommended)

```bash
# Linux / macOS
curl -fsSL https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.sh | bash

# Windows PowerShell
irm https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.ps1 | iex
```

### With flags

```bash
# Preview only (don't install)
curl -fsSL https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.sh | bash -s -- --dry-run

# Force reinstall
curl -fsSL https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.sh | bash -s -- --force

# Use Claude models
irm https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.ps1 | iex -ArgumentList "-Provider claude"
```

## Estructura

```
.
├── opencode/                    # Configuración OpenCode
│   ├── config/
│   │   └── agents.json         # Definiciones de agentes
│   ├── prompts/unity/         # Prompts para cada agente
│   │   ├── unity-6000-expert.md # Orquestador (delega a sub-agentes)
│   │   └── unity-*-expert.md   # 16 sub-agentes especializados
│   └── install.ps1             # Script de instalación (local)
│
├── scripts/
│   ├── install.sh              # One-liner para Linux/macOS
│   └── install.ps1             # One-liner para Windows
│
├── skills/                      # Skills de conocimiento
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
│   ├── unity-input/             # Input System (no legacy)
│   ├── unity-xr/                # AR Foundation, VR
│   ├── unity-ui-*               # UI Toolkit (UXML, USS, C#)
│   ├── unity-cinemachine/       # Virtual cameras, Timeline
│   ├── unity-addressables/      # Async loading, AssetBundles
│   ├── unity-packages/          # Package Manager
│   └── unity-shadergraph/       # Shader Graph, nodes, Master Stack
│
├── agents/                      # Agentes standalone (backup)
├── Manual/                      # Manual convertido a MD (3435 archivos)
└── scripts/                     # Instaladores one-liner
```

## Instalación Local

```powershell
# Clonar el repo
git clone https://github.com/Ulysses-Alv/Unity-Agent-Expert.git
cd Unity-Agent-Expert

# Linux/macOS
chmod +x install.sh
./install.sh --dry-run   # Preview
./install.sh            # Instalar

# Windows PowerShell
.\install.ps1 -DryRun    # Preview
.\install.ps1            # Instalar
```

## Uso

### Activar el orquestador Unity

```
/agent unity-6000-expert
```

El orquestador detecta el dominio del problema y delega al sub-agente apropiado.

### Dominios disponibles

| Dominio | Agente | Modelo |
|---------|--------|--------|
| UI Toolkit | unity-ui-expert | MiMo-V2-Pro |
| Shader Graph | unity-shadergraph-expert | MiMo-V2-Pro |
| 3D Graphics | unity-graphics-expert | MiMo-V2-Pro |
| 2D Graphics | unity-2d-expert | MiniMax M2.7 |
| Physics | unity-physics-expert | Qwen 3.6 Plus |
| Scripting | unity-scripting-expert | Qwen 3.6 Plus |
| Animation | unity-animation-expert | MiMo-V2-Pro |
| Audio | unity-audio-expert | MiniMax M2.7 |
| Editor | unity-editor-expert | MiniMax M2.7 |
| Performance | unity-performance-expert | Qwen 3.6 Plus |
| Build | unity-build-expert | MiniMax M2.7 |
| VFX | unity-vfx-expert | MiMo-V2-Pro |
| Input | unity-input-expert | MiniMax M2.7 |
| XR | unity-xr-expert | Qwen 3.6 Plus |
| Cinemachine | unity-cinemachine-expert | MiMo-V2-Pro |
| Addressables | unity-addressables-expert | Qwen 3.5 Plus |
| Packages | unity-packages-expert | MiniMax M2.7 |

## Arquitectura

```
┌─────────────────────────────────────────────────────┐
│                   USER                               │
└─────────────────────┬───────────────────────────────┘
                      │
                      ▼
         ┌────────────────────────┐
         │  unity-6000-expert     │  ← Orquestador
         │  (detecta dominio)      │
         └────────────┬────────────┘
                      │ delegate (async)
     ┌────────────────┼────────────────┐
     │                │                │
     ▼                ▼                ▼
┌─────────┐     ┌───────────┐    ┌──────────┐
│unity-   │     │unity-    │    │unity-    │
│graphics │     │physics   │    │scripting │
│-expert  │     │-expert   │    │-expert   │
└─────────┘     └───────────┘    └──────────┘
```

## Principios

1. **Delegación async** — El orquestador NUNCA bloquea. Usa `delegate` (async), no `task` (sync).
2. **Skills actualizables** — Los sub-agentes leen skills del filesystem, no tienen conocimiento hardcodeado.
3. **Resiliencia** — Si el sdd-orchestrator actualiza y pisa `opencode.json`, solo corrés `install.ps1` para restaurarlo.
4. **Idempotencia** — Install.ps1 puede correr múltiples veces sin duplicar agentes.

## Actualización

```bash
# One-liner (simple)
curl -fsSL https://raw.githubusercontent.com/Ulysses-Alv/Unity-Agent-Expert/main/scripts/install.sh | bash

# Local (después de git pull)
cd Unity-Agent-Expert
./install.sh
# o
.\install.ps1
```

## Modelos asignados

| Tipo de modelo | Agentes |
|----------------|---------|
| MiniMax M2.7 | 2D, Audio, Editor, Build, Input, Packages |
| MiMo-V2-Pro | Graphics, Animation, VFX, UI, Cinemachine |
| Qwen 3.6 Plus | Physics, Scripting, Performance, XR |
| Qwen 3.5 Plus | Addressables |

## Archivos de respaldo

Si install.ps1 modifica opencode.json, crea backup automáticamente:
```
~/.config/opencode/opencode.json.backup.20260419XXXXXX
```

## Requisitos

- OpenCode instalado
- PowerShell 5.1+
- Acceso a `~/.config/opencode/`

## Licencia

Apache 2.0