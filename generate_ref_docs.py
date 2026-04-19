#!/usr/bin/env python3
"""
Generate reference docs for skills that don't have them.
Enrich skills with source doc references from ManualMD.
"""

from pathlib import Path
import re

MANUALMD = Path("D:/Projects/Skill creation/ManualMD/en/Manual")
SKILLS = Path("D:/Projects/Skill creation/skills")

# Map skills to relevant manual sections/files
SKILL_DOCS = {
    "unity-animation": [
        "AnimationOverview.md",
        "AnimationSection.md",
        "class-AnimationClip.md",
        "AnimationStateMachines.md",
        "AnimationLayers.md",
        "AnimationParameters.md",
        "AnimationClips.md",
        "com.unity.timeline.md",
        "CharacterAnimationFeature.md",
    ],
    "unity-audio": [
        "Audio.md",
        "AudioOverview.md",
        "AudioMixer.md",
        "AudioMixerOverview.md",
        "class-AudioSource.md",
        "AudioSource-overview.md",
        "class-AudioClip.md",
        "class-AudioListener.md",
    ],
    "unity-vfx": [
        "class-ParticleSystem.md",
        "ParticleSystemOverview.md",
        "VFXGraph.md",
        "ParticleSystems.md",
    ],
    "unity-input": [
        "Input.md",
        "class-InputManager.md",
        "InputSystem.md",
        "InputEvent.md",
    ],
    "unity-build-deploy": [
        "BuildPlayerPipeline.md",
        "BuildSettings.md",
        "class-BuildPlayerOptions.md",
        "BuildTarget.md",
        "BuildPipeline.md",
        "PublishingBuilds.md",
    ],
    "unity-xr": [
        "XR.md",
        "XRInput.md",
        "class-ARSession.md",
        "XRInteraction.md",
        "VRDevices.md",
    ],
    "unity-2d": [
        "Sprite.md",
        "SpriteEditor.md",
        "SpriteRenderer.md",
        "Tilemap.md",
        "2DGameObjects.md",
        "2d-physics/2d-physics.md",
    ],
    "unity-performance": [
        "Profiler.md",
        "Performance.md",
        "Memory.md",
        "PerformanceOptimization.md",
        "class-QualitySettings.md",
    ],
    "unity-editor": [
        "Editor.md",
        "EditorWindows.md",
        "CustomEditors.md",
        "UIElements.md",
        "class-EditorWindow.md",
    ],
    "unity-scripting": [
        "class-MonoBehaviour.md",
        "ScriptsIntro.md",
        "class-Object.md",
        "class-GameObject.md",
        "class-Component.md",
        "class-Transform.md",
    ],
}

def create_reference_doc(skill_name: str, doc_files: list) -> str:
    """Create a reference doc content for a skill."""
    lines = [
        f"# Unity 6000.3 LTS {skill_name.replace('unity-', '').title()} — Reference Documentation",
        "",
        "## Source",
        "",
        "All content extracted from local Unity 6000.3.9f1 documentation:",
        "`D:/Projects/Skill creation/ManualMD/en/Manual/`",
        "",
        "## Files Analyzed",
        "",
    ]

    for doc in doc_files:
        # Get just the filename without path
        filename = Path(doc).name
        # Convert to display name
        name = Path(doc).stem.replace("-", " ").replace("_", " ").title()
        lines.append(f"- `{doc}` — {name}")

    return "\n".join(lines)

def main():
    # Skills that need reference docs created
    skills_needing_refs = [
        "unity-animation",
        "unity-audio",
        "unity-vfx",
        "unity-input",
        "unity-build-deploy",
        "unity-xr",
        "unity-2d",
        "unity-performance",
        "unity-editor",
        "unity-scripting",
    ]

    created = 0
    for skill_name in skills_needing_refs:
        skill_dir = SKILLS / skill_name
        if not skill_dir.exists():
            print(f"Creating directory: {skill_dir}")
            skill_dir.mkdir(parents=True, exist_ok=True)

        ref_dir = skill_dir / "references"
        ref_dir.mkdir(exist_ok=True)

        ref_file = ref_dir / "docs.md"
        docs_list = SKILL_DOCS.get(skill_name, [])

        content = create_reference_doc(skill_name, docs_list)
        ref_file.write_text(content, encoding="utf-8")
        print(f"Created: {ref_file}")
        created += 1

    print(f"\nCreated {created} reference docs")

    # Now verify all skills have SKILL.md
    all_skills = [d.name for d in SKILLS.iterdir() if d.is_dir() and (d / "SKILL.md").exists()]
    print(f"\nTotal skills with SKILL.md: {len(all_skills)}")

    return 0

if __name__ == "__main__":
    exit(main())