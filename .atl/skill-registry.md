# Skill Registry

**Delegator use only.** Any agent that launches sub-agents reads this registry to resolve compact rules, then injects them directly into sub-agent prompts. Sub-agents do NOT read this registry or individual SKILL.md files.

See `_shared/skill-resolver.md` for the full resolution protocol.

## User Skills

| Trigger | Skill | Path |
|---------|-------|------|
| When writing physics code in Unity, configuring colliders/rigidbodies, implementing joints or articulations, doing raycasts or physics queries, optimizing physics performance, or working with 2D physics | unity-physics | C:\Users\PC\.config\opencode\skills\unity-6000\unity-physics\SKILL.md |
| When writing C# scripts, MonoBehaviour, coroutines, Jobs, or any Unity scripting task in Unity 6000 | unity-scripting | C:\Users\PC\.config\opencode\skills\unity-6000\unity-scripting\SKILL.md |
| When writing shader code, configuring URP, working with materials, lighting scenes, or implementing rendering features in Unity 6000 | unity-3d-graphics | C:\Users\PC\.config\opencode\skills\unity-6000\unity-3d-graphics\SKILL.md |
| When writing C# code for Unity UI Toolkit, handling UI events, implementing data binding, creating editor windows, runtime UI controllers, or working with VisualElement lifecycle | unity-uitk-csharp | C:\Users\PC\.config\opencode\skills\unity-6000\unity-uitk-csharp\SKILL.md |
| When creating editor tools, custom inspectors, or extending the Unity Editor in Unity 6000 | unity-editor | C:\Users\PC\.config\opencode\skills\unity-6000\unity-editor\SKILL.md |
| When working with 2D sprites, tilemaps, 2D lighting, or any 2D game development task in Unity 6000 | unity-2d | C:\Users\PC\.config\opencode\skills\unity-6000\unity-2d\SKILL.md |
| When creating or editing .uss files, styling Unity UI Toolkit elements, working with USS selectors, transitions, filters, or responsive UI layouts | unity-uss | C:\Users\PC\.config\opencode\skills\unity-6000\unity-uss\SKILL.md |
| When creating or editing .uxml files, defining UI structure in Unity UI Toolkit, working with UXML templates, custom UXML elements, or declarative data binding | unity-uxml | C:\Users\PC\.config\opencode\skills\unity-6000\unity-uxml\SKILL.md |
| When working with cameras, cinematics, Timeline, or virtual camera systems in Unity 6000 | unity-cinemachine | C:\Users\PC\.config\opencode\skills\unity-6000\unity-cinemachine\SKILL.md |
| When loading assets at runtime, using Addressables, or managing asset bundles in Unity 6000 | unity-addressables | C:\Users\PC\.config\opencode\skills\unity-6000\unity-addressables\SKILL.md |
| When creating particle effects, VFX graphs, shader effects, or visual effects in Unity 6000 | unity-vfx | C:\Users\PC\.config\opencode\skills\unity-6000\unity-vfx\SKILL.md |
| When developing AR/VR applications, using XR frameworks, implementing XR interactions, or handling XR devices in Unity 6000 | unity-xr | C:\Users\PC\.config\opencode\skills\unity-6000\unity-xr\SKILL.md |
| When handling player input, using Input System, or migrating from legacy input in Unity 6000 | unity-input | C:\Users\PC\.config\opencode\skills\unity-6000\unity-input\SKILL.md |
| When building players, configuring build settings, or deploying to platforms in Unity 6000 | unity-build-deploy | C:\Users\PC\.config\opencode\skills\unity-6000\unity-build-deploy\SKILL.md |
| When working with character animation, state machines, Timeline sequences, or Playables API in Unity 6000 | unity-animation | C:\Users\PC\.config\opencode\skills\unity-6000\unity-animation\SKILL.md |
| When working with audio playback, mixing, spatial audio, or audio effects in Unity 6000 | unity-audio | C:\Users\PC\.config\opencode\skills\unity-6000\unity-audio\SKILL.md |
| When optimizing performance, profiling, reducing memory usage, or improving frame rate in Unity 6000 | unity-performance | C:\Users\PC\.config\opencode\skills\unity-6000\unity-performance\SKILL.md |
| When working with Unity packages, creating custom packages, or managing dependencies in Unity 6000 | unity-packages | C:\Users\PC\.config\opencode\skills\unity-6000\unity-packages\SKILL.md |
| When creating a new skill, add agent instructions, or document patterns for AI | skill-creator | C:\Users\PC\.config\opencode\skills\skill-creator\SKILL.md |
| When writing Go tests, using teatest, or adding test coverage | go-testing | C:\Users\PC\.config\opencode\skills\go-testing\SKILL.md |
| When creating a pull request, opening a PR, or preparing changes for review | branch-pr | C:\Users\PC\.config\opencode\skills\branch-pr\SKILL.md |
| When creating a GitHub issue, reporting a bug, or requesting a feature | issue-creation | C:\Users\PC\.config\opencode\skills\issue-creation\SKILL.md |
| When user says "judgment day", "review adversarial", "dual review", "doble review", "juzgar", "que lo juzguen" | judgment-day | C:\Users\PC\.config\opencode\skills\judgment-day\SKILL.md |

## Compact Rules

Pre-digested rules per skill. Delegators copy matching blocks into sub-agent prompts as `## Project Standards (auto-resolved)`.

### unity-physics
- Rigidbody movement: use MovePosition/MoveRotation, NEVER set transform.position directly
- Physics calls (raycasts, velocity) go in FixedUpdate, not Update
- 2D physics uses Rigidbody2D/Collider2D with separate layer system
- Joints: Configure anchor/connected anchor precisely; spring forces require MinDistance >= MaxDistance
- Character controllers: Use Physics.SphereCast or CapsuleCast for ground detection, not raycasts
- Physics materials: Combine on colliders for friction/bounce, not on rigidbodies
- Layer overrides in Unity 6000: Use PhysicsLayerPair in URP 3D settings for per-material collision

### unity-scripting
- Never call Unity API from serialization callbacks or field initializers (runs on loading thread)
- Use [SerializeField] for private fields that need Inspector exposure
- Coroutines: store in variable and stop with StopCoroutine on OnDisable
- Jobs: always Schedule and Complete; forget Complete = memory leak
- Burst: only on blittable types; avoid managed strings in job structs
- MonoBehaviour lifecycle: Awake→OnEnable→Start→FixedUpdate→Update→LateUpdate→OnDisable→OnDestroy

### unity-3d-graphics
- URP uses RendererFeatures for post-processing; built-in uses PostProcessing package
- ShaderLab for structure, HLSL for actual shader code (CG is legacy)
- Master Stack in Shader Graph outputs to PBR or Unlit only
- Lightmapping: mark static, use Contribute Global Illumination on lightmapped objects
- GPU Resident Drawer in Unity 6000 reduces draw call overhead for skinned meshes
- URP Render Graph: use RenderGraph API to track resource lifetime automatically

### unity-uitk-csharp
- Runtime UI: UIDocument with VisualTreeAsset (UXML); Editor UI: UXML loaded via UXML Lourve
- Element query: Query with<T>() after element is part of tree; cached query is faster
- Data binding: SerializedObject for inspector, Bind() method for runtime
- Custom controls: [UxmlElement] + partial class + UXML factory
- ListView: handle item activation in makeItem, bind data in bindItem
- Dispose VisualElements to prevent memory leaks in long-lived UIs

### unity-editor
- EditorWindow: Use CreateInstance, not new; get layout rect via position
- Custom Inspectors: use [CustomEditor(typeof(MyComponent))] + Editor class
- PropertyDrawer: [CustomPropertyDrawer(typeof(MyType))] for attribute-based editing
- Gizmos: draw in OnDrawGizmos (painted in scene) vs OnDrawGizmosSelected (when selected)
- UI Toolkit preferred over IMGUI in Unity 6000; use IMGUI only for quick debug tools

### unity-2d
- Sprites: use SpriteRenderer with SpriteMask for selective visibility
- Tilemaps: use TilemapCollider2D + Rigidbody2D (Static/Kinematic as needed)
- 2D lights: URP 2D renderer uses Light2D components; shape types: freeform, polygon, sprite
- Sorting: use SortingGroup for complex objects, Order in Layer for simple case
- 9-slice scaling: Sprite type = Sliced in SpriteRenderer
- 2D physics queries: Physics2D.Raycast, BoxCast, CircleCast, Overlap — all have 2D suffix

### unity-uss
- Flexbox layout: use style.flexDirection, style.alignItems, style.justifyContent
- USS variables: define in theme file, reference with var(--my-variable)
- Transitions: style.transition with timing functions; use transition-delay for sequences
- Responsive: combine media queries with USS; use -unity-device-density for screen adaptation
- Custom filters: implement USS filter interface, apply via style.filter

### unity-uxml
- Template root must have xmlns:xe="unityEngine/Editor/Experimental/UIElements"
- Custom elements: [UxmlElement] attribute, partial class, optional [UxmlAttributeDescription]
- Data binding: use bind- attribute with bind path; two-way binding via SerializedObject
- VisualElement hierarchy: parent-child via Add(child) or USS <层级>
- Instantiate template: UIDocument.rootVisualElement.Add(template.Instantiate())

### unity-cinemachine
- Virtual cameras: set Body for follow, Aim for look-at; use Framing Transposer for 2D
- Dolly tracks: use CinemachinePath, animate position along path with CinemachineDollyCart
- Priority: higher priority camera wins; use for cutscene switching
- Noise: add noise profile to simulate handheld, shake; configure amplitude and frequency
- Blend: configure durations in Brain; custom blend via Animation curves

### unity-addressables
- Address.LoadAssetAsync<T>() for single asset; LoadAssetAsync for multiple
- Release: call Release on the handle, not the asset object; tracks reference count
- Groups: create groups per feature/asset type; use Play mode: Use Asset Database for dev
- Labels: use for conditional loading; Addressables.GetDownloadSizeAsync for预估
- Async: always await/callback; never access asset before load completes

### unity-vfx
- Particle Systems (Shuriken): use modules (Emission, Shape, Velocity over Lifetime) efficiently
- VFX Graph: use GPU events for complex spawn patterns; CPU events for small counts
- HDRP/URP: VFX Graph renders correctly in both; Shuriken needs renderer module configured
- LOD: particle systems rarely need LOD; consider budget in systems with many particles
- Optimize: set max particles cap; use AutoRandomSeed; avoid in-world space for mobile

### unity-xr
- AR Foundation: use XRPlaneManager, XRAnchorManager, XRRayInteractor for AR features
- XR Interaction Toolkit: use interactors (ray, direct, teleport) + interactables (grab, select)
- Meta Quest: use OXR plugin; recommend using interaction profiles
- VisionOS: spatial video, eye tracking via XRKit
- Hand tracking: XRHandTrackingInteractor for articulated hands

### unity-input
- Input System: use InputAction with callbacks (started/performed/completed), not polling
- InputBindings: use <Keyboard>/<Mouse>/<Gamepad> paths; composite bindings for analog+direction
- PlayerInput: use InvokeUnityEvents for callbacks; ConfigureAwait ensures player context
- Migration: legacy Input.GetAxis/GetButton still works; Input System offers consistency
- Actions maps: switch maps for different game states (e.g., walking vs driving)

### unity-build-deploy
- BuildPlayerOptions: set scenes, locationPathName, target, options
- IL2CPP: enables AOT compilation; some .NET features unavailable (Reflection.Emit)
- Platform: switch BuildTarget before building; handle platform-specific #if UNITY_EDITOR
- AssetBundles: build with BuildAssetBundleOptions; load via AssetBundle.LoadFromFileAsync

### unity-animation
- Animator Controller: use layers for additive animation (e.g., full-body + upper-body)
- Blend trees: 1D/2D freeform Cartesian; check thresholds visually in inspector
- Animator: use triggers for one-shot animations, bools for sustained states
- Timeline: use RuntimeAnimatorController on GameObject; tracks drive animation directly
- Playables API: for custom mixer logic; AnimationMixerPlayable for blending multiple clips

### unity-audio
- AudioSource: 3D sound uses spatialBlend, rolloff; 2D uses spatialBlend = 0
- AudioMixer: use snapshot transitions for smooth mix changes; send levels for real-time mixing
- Spatial audio: enable Spatialize on AudioSource + AudioMixer; use Sonть spacing for early reflections
- WebGL: audio requires audio context Resume() on user interaction

### unity-performance
- Profiler: use CPUProfilerModule to find GC allocations; GPUProfilerModule for render bottlenecks
- Memory: use Profiler.LiveObjectCount to track object counts; profile memory by category
- GPU instancing: enable on materials for draw call reduction; use Graphics.DrawMeshInstanced
- SRP Batcher: keep materials consistent across frames; SRP Batcher handles rest automatically
- Script optimization: avoid allocations in Update; use struct for frequently-called data

### unity-packages
- Package Manager: scoped registries for third-party packages; Unity registry is default
- Custom package: create package.json with name, version, dependencies; place in Assets or git
- Dependency: specify exact versions for stability; check Unity version compatibility

### skill-creator
- Skill structure: SKILL.md required, assets/ optional, references/ optional
- SKILL.md frontmatter: name, description (with Trigger:), license, metadata
- Compact rules: 5-15 lines of actionable patterns, not documentation
- Project skills: place in skills/ subdirectory of project root

### go-testing
- Table-driven tests: use struct with name, input, expected, wantErr fields
- teatest: for Bubbletea TUI testing; use ft.DeepEqual for assertion
- Golden files: use golden.Read or golden.Write in testdata/ subdirectory
- Test coverage: run go test -coverprofile=cover.out && go tool cover

### branch-pr
- Every PR MUST link an approved issue — no exceptions
- Every PR MUST have exactly one type:* label (type/feat, type/fix, type/docs, etc.)
- Automated checks must pass before merge
- Use conventional commits: feat:, fix:, docs:, chore:, refactor:

### issue-creation
- Bug reports: include reproduction steps, expected vs actual behavior, Unity version
- Feature requests: describe problem solved, not just the solution
- Use labels to categorize: bug, enhancement, documentation, question

### judgment-day
- Launch two independent blind judge sub-agents simultaneously
- Synthesize findings, apply fixes, re-judge until both pass or escalate after 2 iterations

## Project Conventions

| File | Path | Notes |
|------|------|-------|
| README | D:\Projects\Skill creation\README.md | Unity-Agent-Expert project - skill repository for Unity 6000.3 LTS |
| Agents Config | D:\Projects\Skill creation\opencode\config\agents.json | 18 specialized sub-agents + orchestrator |
| .atl Registry | D:\Projects\Skill creation\.atl\skill-registry.md | This file |

Read the convention files listed above for project-specific patterns and rules. All referenced paths have been extracted — no need to read index files to discover more.
