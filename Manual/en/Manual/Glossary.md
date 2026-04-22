# Glossary

| [0–9](#glossary-1) | [A](#glossary-a) | [B](#glossary-b) | [C](#glossary-c) | [D](#glossary-d) | [E](#glossary-e) | [F](#glossary-f) | [G](#glossary-g) |
| [H](#glossary-h) | [I-J](#glossary-ij) | [K](#glossary-k) | [L](#glossary-l) | [M](#glossary-m) | [N](#glossary-n) | [O](#glossary-o) | [P](#glossary-p) |
| [Q](#glossary-q) | [R](#glossary-r) | [S](#glossary-s) | [T](#glossary-t) | [U](#glossary-u) | [V](#glossary-v) | [W](#glossary-w) | [X-Z](#glossary-xz) |

## 0–9

#### 1D Blend Tree:

A Blend Tree for 1D blending, which blends motion according to a single Animation Parameter. [More info](BlendTree-1DBlending.md)

#### 2D Blend Tree:

A Blend Tree for 2D blending, which blends motion according to two **Animation Parameters**. [More info](BlendTree-2DBlending.md)

#### 2D Object:

A 2D **GameObject** such as a **tilemap** or **sprite**. [More info](Unity2D.md)

#### 3D Object:

A 3D **GameObject** such as a cube, **terrain** or ragdoll. [More info](GameObjects.md)

## A

#### Accelerator:

The Unity **Accelerator** is an external tool that provides an asset cache that keeps copies of a team’s imported assets. The goal of the **Accelerator** is to speed up teamwork and reduce iteration time by coordinating asset sharing so that you don’t need to reimport portions of your project. [More info](UnityAccelerator.md)

#### active users:

Players who recently played your game. Unity **Analytics** defines an active player as someone who has played within the last 90 calendar days. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### Ad ARP:

(Average Revenue Per User) Average Unity Ads revenue per player. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### ad revenue:

Total Unity Ads revenue. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### ad starts:

The number of video ads that started playing. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### ADB:

An Android Debug Bridge (ADB). You can use an **ADB** to deploy an Android package (APK) manually after building. [More info](https://developer.android.com/studio/command-line/adb.html)

#### ads per DAU:

The number of ads started per active player on a given day. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### age 14 and under:

By default, Unity does not breakout **analytics** data for players under the age of 14. See **COPPA** Compliance. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### all spenders:

Players who have made any verified or unverified in-app purchases in their lifetime. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### ambient light:

Light that doesn’t come from any specific direction, and contributes equal light in all directions to the **Scene**. [More info](lighting-window.md)

#### ambient occlusion:

A method to approximate how much **ambient light** (light not coming from a specific direction) can hit a point on a surface.

#### analytics events:

Events dispatched to the **Analytics** Service by instances of your applications. **Analytics** events contain the data that is processed and aggregated to provide insights into player behavior. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/events)

#### Analytics:

Abbreviation of **Unity Analytics**

#### anchor:

A UI layout tool that fixes a UI element to a parent element. Anchors are shown as four small triangular handles in the **Scene** view and anchor information is also shown in the **Inspector**. [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/UIBasicLayout.html#anchors)

#### Android Keystore:

An Android system that lets you store cryptographic key entries for enhanced device security. [More info](class-PlayerSettingsAndroid.md#projectkeystore)

#### Animation Blend Shape:

Enables you to make an object change its form by blending between two separate meshes. [More info](BlendShapes.md)

#### Animation Blend Tree:

Used for continuous blending between similar **Animation Clips** based on float **Animation Parameters**. [More info](class-BlendTree.md)

#### Animation Clip:

Animation data that can be used for animated characters or simple animations. It is a simple “unit” piece of motion, such as (one specific instance of) “Idle”, “Walk” or “Run”. [More info](class-AnimationClip.md)

#### Animation Clip Node:

A node in a Blend Tree graph that contains an **animation clip**, such as a run or walk animation. [More info](class-BlendTree.md)

#### animation compression:

The method of compressing animation data to significantly reduce file sizes without causing a noticeable reduction in motion quality. Animation **compression** is a trade off between saving on memory and image quality. [More info](class-AnimationClip.md#AssetProperties)

#### Animation Curves:

Allows you to add data to an imported clip so you can animate the timings of other items based on the state of an animator. For example, for a game set in icy conditions, you could use an extra animation curve to control the emission rate of a **particle** system to show the player’s condensing breath in the cold air. [More info](AnimationCurvesOnImportedClips.md)

#### Animation Event:

Allows you to add data to an imported clip which determines when certain actions should occur in time with the animation. For example, for an animated character you might want to add events to walk and run cycles to indicate when the footstep sounds should play. [More info](AnimationEventsOnImportedClips.md)

#### Animation Key:

The value of an animatable property, set at a specific point in time. Setting at least two keys for the same property creates an animation. [More info](animeditor-AnimationCurves.md)

#### Animation Layer:

An **Animation Layer** contains an Animation **State Machine** that controls animations of a model or part of it. An example of this is if you have a full-body layer for walking or jumping and a higher layer for upper-body motions such as throwing an object or shooting. The higher layers take precedence for the body parts they control. [More info](AnimationLayers.md)

#### Animation Parameters:

Used to communicate between scripting and the **Animator Controller**. Some parameters can be set in scripting and used by the controller, while other parameters are based on Custom Curves in **Animation Clips** and can be sampled using the scripting API. [More info](AnimationParameters.md)

#### Animation State Machine:

A graph within an **Animator Controller** that controls the interaction of Animation States. Each state references an **Animation Blend Tree** or a single **Animation Clip**. [More info](AnimationStateMachines.md)

#### Animation Transition:

Allows a **state machine** to switch or blend from one animation state to another. Transitions define how long a blend between states should take, and the conditions that activate them. [More info](StateMachineTransitions.md)

#### Animator Component:

A component on a model that animates that model using the Animation system. The component has a reference to an **Animator Controller** asset that controls the animation. [More info](class-AnimatorController.md)

#### Animator Controller:

Controls animation through **Animation Layers** with Animation **State Machines** and **Animation Blend Trees**, controlled by **Animation Parameters**. The same **Animator Controller** can be referenced by multiple models with **Animator components**. [More info](class-AnimatorController.md)

#### Animator Override Controller:

Allows you to create multiple variants of an **Animator Controller**, with each variant using a different set of animations, while retaining the original Controller’s structure, parameters and logic. [More info](AnimatorOverrideController.md)

#### Animator Window:

The window where the **Animator Controller** is visualized and edited. [More info](AnimatorWindow.md)

#### Aniso Level:

The anisotropic filtering (AF) level of a texture. Allows you to increase texture quality when viewing a texture at a steep angle. Good for floor and ground textures. [More info](class-TextureImporter.md)

#### antialiasing:

A technique for decreasing artifacts, like jagged lines (jaggies), in images to make them appear smoother.

#### AOT Compilation:

Ahead of Time (AOT) compilation is an optimization method used by all platforms except iOS for optimizing the size of the built player. . [More info](iphone-playerSizeOptimization.md)

#### APK:

The Android Package format output by Unity. An APK is automatically deployed to your device when you select File > Build & Run. [More info](android-BuildProcess.md)

#### application version:

Player segments based on application version or bundleid. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### AR:

Augmented Reality [More info](AROverview.md)

#### ARPDAU:

(Average Revenue Per Daily Active User) The average revenue per user who played on a given day. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### ARPPU:

(Average Revenue Per Paying User) Average **verified IAP revenue** per user who completed a verified IAP transaction. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### aspect ratio:

The relationship of an image’s proportional dimensions, such as its width and height.

#### asset package:

A collection of files and data from Unity projects, or elements of projects, which are compressed and stored in one file, similar to Zip files, with the `.unitypackage` extension. **Asset packages** are a handy way of sharing and re-using Unity projects and collections of assets. [More info](AssetPackages.md)

#### Asset Server:

Legacy - An asset and **version control** system with a graphical user interface integrated into Unity. Enables team members to work together on a project on different computers.

#### asset:

Any media or data that can be used in your game or project. An asset may come from a file created outside of Unity, such as a 3D Model, an audio file or an image. You can also create some asset types in Unity, such as an **Animator Controller**, an Audio Mixer or a **Render Texture**. [More info](AssetWorkflow.md)

#### Asset Store:

A growing library of free and commercial assets created by Unity and members of the community. Offers a wide variety of assets, from textures, models and animations to whole project examples, tutorials and Editor extensions. [More info](AssetStore.md)

#### Asset Store package:

A bundled collection of assets or tools available for purchase or download on the Unity **Asset Store**, compressed and stored as an **asset package** (`.unitypackage` extension) or a **UPM package**. You can manage your packages from the **Asset Store** either on the online store or through the Package Manager window. [More info](AssetStorePackages.md)

#### Audio Clip:

A container for audio data in Unity. Unity supports mono, stereo and multichannel audio assets (up to eight channels). Unity can import .aif, .wav, .mp3, and .ogg audio file format, and .xm, .mod, .it, and .s3m tracker module formats. [More info](class-AudioClip.md)

#### Audio Distortion Filter:

An **audio filter** that distorts the sound from an **Audio Source** or sounds reaching the **Audio Listener** by simulating the sound of a low quality radio transmission. [More info](class-AudioDistortionFilter.md)

#### Audio Effect:

Any effect that can modify the output of Audio Mixer components, such as filtering frequency ranges of a sound or applying reverb. [More info](class-AudioEffectMixer.md)

#### Audio Filter:

Any **audio filter** that distorts the sound from an **Audio Source** or sounds reaching an **Audio Listener**. [More info](class-AudioEffect.md)

#### Audio High Pass Filter:

An **audio filter** that passes high frequencies of an AudioSource and cuts off signals with frequencies lower than the Cutoff Frequency. [More info](class-AudioHighPassFilter.md)

#### Audio Listener:

A component that acts like a microphone, receiving sound from **Audio Sources** in the **scene** and outputting to the computer speakers. [More info](class-AudioListener.md)

#### Audio Low Pass Filter:

An **audio filter** that passes low frequencies of an **Audio Source** or all sound reaching an **Audio Listener** while removing frequencies higher than the Cutoff Frequency. [More info](class-AudioLowPassFilter.md)

#### Audio Source:

A component which plays back an **Audio Clip** in the **scene** to an **audio listener** or through an audio mixer. [More info](class-AudioSource.md)

#### Audio Spatializer:

A **plug-in** that changes the way audio is transmitted from an **audio source** into the surrounding space. It takes the source and regulates the gains of the left and right ear contributions based on the distance and angle between the AudioListener and the AudioSource. [More info](AudioSpatializerSDK.md)

#### Augmented Reality:

Augmented Reality (AR) uses computer graphics or video composited on top of a live video feed to augment the view and create interaction with real and virtual objects. [More info](AROverview.md)

#### Avatar Mask:

A specification for which body parts to include or exclude for an animation rig. Used in **Animation Layers** and in the importer. [More info](class-AvatarMask.md)

#### Avatar:

An interface for **retargeting** animation from one rig to another. [More info](ConfiguringtheAvatar.md)

## B

#### baked lights:

Light components whose Mode property is set to Baked. Unity pre-calculates the illumination from **Baked Lights** before runtime, and does not include them in any runtime lighting calculations. [More info](LightModes-introduction.md#baked)

#### billboard:

A textured **2D object** that rotates so that it always faces the **Camera**. [More info](class-BillboardRenderer.md)

#### Bind-pose:

The pose at which the character was modeled.

#### Blend Node:

A node in a Blend Tree graph that blends **animation clip** nodes. [More info](class-BlendTree.md)

#### blend:

Transition from one animation to another smoothly and seamlessly, such as blending a character’s walking and running animations according to the character’s speed.

#### blit:

A shorthand term for “bit block transfer”. A **blit** operation is the process of transferring blocks of data from one place in memory to another.

#### Bloom:

A **post-processing** effect used to reproduce an imaging artifact of real-world **cameras**. The effect produces fringes of light extending from the borders of bright areas in an image, contributing to the illusion of an extremely bright light overwhelming the **camera** or eye capturing the **scene**.

#### Body Transform:

The mass center of the character. It is used in for animation **retargeting** and provides the most stable displacement model. [More info](RootMotion.md)

#### Body Type:

Defines a fixed behavior for a 2D **Rigidbody**. Can be Dynamic (the body moves under simulation and is affected by forces like gravity), Kinematic (the body moves under simulation, but isn’t affected by forces like gravity) or Static (the body doesn’t move under simulation). [More info](2d-physics/rigidbody/introduction-to-rigidbody-2d.md)

#### bounding volume:

A closed shape representing the edges and faces of a **collider** or trigger.

#### Box Collider:

A cube-shaped **collider** component that handles **collisions** for **GameObjects** like dice and ice cubes. [More info](class-BoxCollider.md)

#### broad phase collision detection:

A **collision** detection phase that computes pairs of potentially overlapping objects by judging only their respective **bounding volumes**. [More info](http://planning.cs.uiuc.edu/node214.html)

#### Build Automation:

A continuous integration service for Unity projects that automates the process of creating builds on Unity’s servers. [More info](https://docs.unity.com/devops/en/manual/unity-build-automation)

#### build profile:

A set of customizable configuration settings to use when creating a build for your target platform. [More info](build-profiles.md)

#### build:

The process of compiling your project into a format that is ready to run on a specific platform or platforms. [More info](building-introduction.md)

#### built-in package:

*Built-in* packages allow users to toggle Unity features on or off through the Package Manager. Enabling or disabling a package reduces the run-time build size. For example, most projects don’t use the legacy **Particle** System. By removing the abstracted package of this feature, the related code and resources are not part of the final built product. Typically, these packages contain only the **package manifest** and are bundled with Unity (rather than available on the package registry).

#### bump map:

An image texture used to represent geometric detail across the surface of a **mesh**, for example bumps and grooves. Can be represented as a **heightmap** or a **normal map**. [More info](StandardShaderMaterialParameterNormalMap.md)

#### bundled package:

Unity stores *bundled* packages in the [global cache](upm-cache.md) when you install Unity. You can install these packages in your project even if you are completely offline (not currently connected to the internet or a local network).

## C

#### Cache API:

A Javascript API to store network request and response pairs in the browser cache. [More info](https://developer.mozilla.org/en-US/docs/Web/API/Cache)

#### call stack:

A list of methods that were called at run time, organized as a last-in-first-out stack.

#### Camera:

A component which creates an image of a particular viewpoint in your **scene**. The output is either drawn to the screen or captured as a texture. [More info](CamerasOverview.md)

#### Canvas Group:

A group of UI elements within a Canvas. Use a **Canvas Group** to control a group of UI elements collectively without needing to handle them each individually. [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/class-CanvasGroup.html)

#### Canvas Renderer:

Renders a graphical UI object contained within a Canvas. [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/class-CanvasRenderer.html)

#### Canvas Scaler:

Controls the overall scale and **pixel** density of all UI elements in the Canvas, including font sizes and image borders. [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/script-CanvasScaler.html)

#### Canvas:

The area that contains all UI elements in a **scene**. The Canvas area is shown as a rectangle in the **Scene** View. [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/UICanvas.html)

#### Capsule Collider:

A capsule-shaped **collider** component that handles **collisions** for **GameObjects** like barrels and character limbs. [More info](class-CapsuleCollider.md)

#### category:

A **Profiler category** identifies the workload data for a Unity subsystem (for example, Rendering, Scripting and Animation categories). Unity applies color-coding to categories to help visually distinguish the types of data in the Profiler window. [More info](ProfilerWindow.md)

#### center of mass:

Represents the average position of all mass in a **Rigidbody** for the purposes of physics calculations. By default it is computed from all **colliders** belonging to the **Rigidbody**, but can be modified via script. [More info](../ScriptReference/Rigidbody-centerOfMass.md)

#### Character Controller:

A simple, capsule-shaped **collider** component with specialized features for behaving as a character in a game. Unlike true **collider** components, a **Rigidbody** is not needed and the momentum effects are not realistic. [More info](class-CharacterController.md)

#### Character Joint:

An extended ball-socket **joint** which allows a **joint** to be limited on each axis. Mainly used for Ragdoll effects. [More info](class-CharacterJoint.md)

#### churn:

The rate at which users are leaving your game during a specified period. Your user churn is important in estimating the lifetime value of your users. Mathematically, churn is the complement of retention (in other words: Churn + Retention = 100%). [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### clipping plane:

A plane that limits how far or close a **camera** can see from its current position. A **camera**’s viewable range is between the far and near **clipping planes**. See far **clipping plane** and near **clipping plane**. [More info](class-Camera.md)

#### closed platform:

Includes platforms that require confidentiality and legal agreements with the platform provider for using their developer tools and hardware. These platforms aren’t open to development unless you have an established relationship with the provider. For example PlayStation®, Game Core for Xbox®, and Nintendo®.

#### Cloth:

A component that works with the Skinned **Mesh** Renderer to provide a physics-based solution for simulating fabrics. [More info](class-Cloth.md)

#### cohort:

A group of players with at least one similar characteristic. You can define and analyze different cohorts of your user base with segments. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### collider:

An invisible shape that is used to handle physical **collisions** for an object. A **collider** doesn’t need to be exactly the same shape as the object’s **mesh** - a rough approximation is often more efficient and indistinguishable in gameplay. [More info](CollidersOverview.md)

#### collision detection:

An automatic process performed by Unity which determines whether a moving **GameObject** with a **Rigidbody** and **collider** component has come into contact with any other **colliders**. [More info](CollidersOverview.md)

#### collision:

A **collision** occurs when the **physics engine** detects that the **colliders** of two **GameObjects** make contact or overlap, when at least one has a **Rigidbody** component and is in motion. [More info](CollidersOverview.md)

#### component:

A functional part of a **GameObject**. A **GameObject** can contain any number of components. Unity has many built-in components, and you can create your own by writing **scripts** that inherit from MonoBehaviour. [More info](UsingComponents.md)

#### compression:

A method of storing data that reduces the amount of storage space it requires. See [Texture Compression](texture-choose-format-by-platform.md), [Animation Compression](class-AnimationClip.md#AssetProperties), [Audio Compression](class-AudioClip.md), [Build Compression](ReducingFilesize.md).

#### Configurable Joint:

An extremely customizable **joint** that other **joint** types are derived from. It can be used to create anything from adapted versions of existing **joints** to custom designed and highly specialized **joints**. [More info](class-ConfigurableJoint.md)

#### Console window:

A Unity Editor window that shows errors, warnings and other messages generated by Unity, or your own **scripts**. [More info](Console.md)

#### console:

Abbreviation of **game console**

#### Constant Force:

A simple component for adding a **constant force** or torque to game objects with a **Rigidbody**. [More info](class-ConstantForce.md)

#### constraints:

Settings on **Joint** components which limit movement or rotation. The type and number of constraints vary depending on the type of **Joint**. [More info](2d-physics/joints/2d-joints-landing.md)

#### contact distance:

A **joint** limit property that sets the minimum distance tolerance between the **joint** position and the limit at which the limit will be enforced. [More info](class-CharacterJoint.md)

#### continuous collision detection:

A **collision** detection method that calculates and resolves **collisions** over the entire physics simulation step. This can prevent fast-moving objects from tunnelling through walls during a simulation step. [More info](ContinuousCollisionDetection.md)

#### conversion rate:

The percentage of users who complete an action or sequence of actions. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### COPPA:

(Children’s Online Privacy Protection Act) **COPPA** is a US law that applies to apps that collect personal information and are targeted to children under the age of 13. [More info](https://docs.unity.com/ugs/manual/analytics/manual/privacy-overview)

#### core events:

Core events are the basic events dispatched by the Unity **Analytics** code in your game. These events, and the **analytics** based on them, become available by turning on Unity **Analytics** for a project. Core events include: app running, app start, and device info. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/events)

#### CTR:

(Click Through Rate) The percentage of players who click a link in an ad displayed in your game. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### Cubemap:

A collection of six square textures that can represent the reflections in an environment or the **skybox** drawn behind your geometry. The six squares form the faces of an imaginary cube that surrounds an object; each face represents the view along the directions of the world axes (up, down, left, right, forward and back). [More info](class-Cubemap-landing.md)

#### Culling Mask:

Allows you to include or omit objects to be rendered by a **Camera**, by Layer.

#### custom events:

Custom events are freeform events that you can dispatch when an appropriate **standard event** is not available. Custom events can have any name and up to ten parameters. Use **standard events** in preference to custom events where possible. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/custom-event)

## D

#### damping ratio:

A **joint** setting to control spring oscillation. A higher **damping ratio** means the spring will come to rest faster. [More info](2d-physics/joints/fixed-joint-2d-reference.md)

#### Data Explorer:

A Unity **Analytics** Dashboard page that allows you to build, view and export reports on your **Analytics** metrics and events. You can also see how metrics and custom events change over time. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer)

#### DAU per MAU:

(DAU/MAU) The percentage of monthly active users who play on a given day. Also known as **Sticky Factor** in the **analytics** and game industries, this metric is often used as one estimate of player engagement. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### DAU:

(Daily Active Users) The number of different players who started a session on a given day. **DAU** includes both new and returning players. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### day 1 retention:

The percentage of players who returned to your game one day after playing the first time. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### day 30 retention:

The percentage of players who returned to your game thirty days after playing the first time. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### day 7 retention:

The percentage of players who returned to your game seven days after playing the first time. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### default package:

Unity automatically pre-installs a select number of *default* packages (for example, the Analytics Library, Unity Timeline, etc.) when you create a new project. This differs from a **bundled package** because you don’t need to install it and it differs from a **built-in package** because it extends Unity’s features rather than being able to enable or disable them.

#### deferred shading:

A **rendering path** in the Built-in **Render Pipeline** that places no limit on the number of Lights that can affect a **GameObject**. All Lights are evaluated per-pixel, which means that they all interact correctly with **normal maps** and so on. Additionally, all Lights can have cookies and shadows. [More info](RenderTech-DeferredShading.md)

#### demographics:

Player segments based on reported demographics. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### depth buffer:

A memory store that holds the z-value depth of each **pixel** in an image, where the z-value is the depth for each rendered **pixel** from the projection plane. [More info](class-RenderTexture.md)

#### depth of field:

A **post-processing** effect that simulates the focus properties of a **camera** lens. [More info](PostProcessingOverview.md)

#### development build:

A **development build** includes debug symbols and enables the Profiler. [More info](building-introduction.md)

#### dimetric projection:

A form of parallel projection where the dimensions of a **3D object** are projected onto a 2D plane, and only two of the three angles between the axes are equal to each other. This form of projection is commonly used in isometric video games to simulate three-dimensional depth. [More info](tilemaps/work-with-tilemaps/isometric-tilemaps/isometric-tilemap-landing.md)

#### Direct Blend Tree:

A Blend Tree that allows you to map animator parameters directly to the weight of a Blend Tree child. This is useful when you want to have exact control over the animations that are being blended rather than blend them indirectly using one or two parameters (in the case of 1D and 2D blend trees). [More info](BlendTree-DirectBlending.md)

#### Direct dependency:

A **direct** dependency occurs when your project “requests” a specific package version. To create a **direct dependency**, you add that package and version to the **dependencies** property in your **project manifest** (expressed in the form `package_name@package_version`). [More info](upm-dependencies.md)

#### discrete collision detection:

A **collision** detection method that calculates and resolves **collisions** based on the pose of objects at the end of each physics simulation step. [More info](class-Rigidbody.md)

#### Distance Shadowmask:

A version of the **Shadowmask** lighting mode that includes high quality shadows cast from static **GameObjects** onto dynamic **GameObjects**. [More info](lighting-mode.md#shadowmask)

#### distortion effect:

An **audio effect** that modifies the sound by squashing and clipping the waveform to produce a rough, harsh result. [More info](class-AudioDistortionFilter.md)

#### dolphins:

Players who have spent between $5 and $19.99. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### doppler factor:

An audio setting that allows you to control how much the velocity of an object (relative to the audio listener) affects the pitch of any **audio sources** attached to it. [More info](class-AudioManager.md)

#### dry level:

An audio setting that allows you to set the mix level of unprocessed signal in output in mB.

#### dry mix:

An audio setting that allows you to set the volume of the original signal to pass to output.

#### dynamic batching:

An automatic Unity process which attempts to render multiple meshes as if they were a single **mesh** for optimized graphics performance. The technique transforms all of the **GameObject** vertices on the CPU and groups many similar vertices together. [More info](DrawCallBatching.md)

#### dynamic friction:

A **Physics Material** property that defines the friction for a **Rigidbody** when it’s in motion. Lower values mean less friction, so a setting of zero represents slipping on ice. [More info](class-PhysicsMaterial.md)

#### dynamic resolution:

A **Camera** setting that allows you to dynamically scale individual render targets to reduce workload on the GPU. [More info](DynamicResolution-landing.md)

## E

#### eCPM:

(estimated Cost Per Mille) The estimated revenue for 1000 ad impressions for your app. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### Editor scripts:

C# source files composed entirely of code that runs in the Unity Editor only and not in the runtime Player build. Keep such **scripts** in dedicated Editor assemblies either by placing them in a parent folder called Editor or creating an Editor-only assembly definition.

#### embedded package:

An *embedded* package is a **mutable** package that you store under the `Packages` directory at the root of a Unity project. This differs from most packages which you download from a package server and are **immutable**. [More info](upm-concepts.md#Embedded)

#### Emscripten:

The toolchain that Unity uses to convert from C and C++ to WebAssembly. [More info](https://emscripten.org)

#### engagement:

Engagement is a broad measure of how players enjoy, or are otherwise invested, in your game. Impossible to measure directly, the following metrics are frequently used to estimate engagement: Retention, **DAU**, **MAU**, **DAU**/**MAU**, number of sessions, and session length. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### Enlighten:

A lighting system by Geomerics used in Unity for **Enlighten** Realtime **Global Illumination**. [More info](https://www.siliconstudio.co.jp/en/products-service/enlighten/)

#### ETC:

(Ericsson Texture Compression) A block-based **texture format** that compresses textures to significantly reduce file sizes without causing a noticeable reduction in image quality. [More info](texture-choose-format-by-platform.md)

#### Event System:

A way of sending events to objects in the application based on input, be it keyboard, mouse, touch, or custom input. The **Event System** consists of a few components that work together to send events. [More info](UIE-Runtime-Event-System.md)

#### exponential fog:

A fog model that emulates realistic fog behavior by simulating light absorption over distance by a certain attenuation factor.

#### exposure value:

A value that represents a combination of a **camera**’s shutter speed and f-number. It is essentially a measurement of exposure such that all combinations of shutter speed and f-number that yield the same level of exposure have the same EV.

#### extrapolate:

See **Extrapolation**

#### extrapolation:

The process of storing the last few known values and using them to predict future values. Used in animation, physics and multiplayer.

#### Extrude Edges:

A Texture property that enables you to define how much area to leave around a **sprite** in a generated **mesh**.

## F

#### F2P:

(Free to Play) A business model that offers users free access to a fully functional game and a significant portion of app content. Monetization strategies for these titles generally include microtransactions that allow users to access premium features and virtual goods. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### far clipping plane:

The maximum draw distance for a **camera**. Geometry beyond the plane defined by this value is not rendered. The plane is perpendicular to the **camera**’s forward (Z) direction.

#### FBX:

Autodesk’s proprietary format that Unity uses to import and export Models, animation, and more. [More info](ImportingAssets.md)

#### feature set:

A feature set is a collection of related packages that you can use to achieve specific results in the Unity Editor. You can manage feature sets directly in Unity’s Package Manager. [More info](FeatureSets.md)

#### fill rate:

The rate at which ads are available when you request one. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### first person shooter:

A common game genre, featuring a first-person view of a 3D world, and gun-based combat with other players or NPCs.

#### Fixed Joint:

A **joint** type that is completely constrained, allowing two objects to be held together. Implemented as a spring so some motion may still occur. [More info](class-FixedJoint.md)

#### Fixed Joint 2D:

A 2D **joint** type which is completely constrained, allowing two objects to be held together. Implemented as a spring so some small motion may still occur. [More info](2d-physics/joints/fixed-joint-2d-reference.md)

#### fixed timestep:

A customizable frame-rate-independent interval that dictates when physics calculations and FixedUpdate() events are performed. [More info](class-TimeManager.md)

#### flythrough mode:

A **Scene** view navigation mode that allows you to fly around the **scene** in first-person, similar to how you would navigate in many games. [More info](SceneViewNavigation.md#flythrough)

#### FMOD:

Audio in Unity is built on top of a middleware called FMOD. FMOD is integrated with the Unity engine for creating and playing back interactive audio.

#### fog:

A **post-processing** effect that overlays a color onto objects depending on the distance from the **camera**. Use this to simulate fog or mist in outdoor environments, or to hide clipping of objects near the **camera**’s far clip plane. [More info](PostProcessingOverview.md)

#### forward kinematics:

A method of posing a skeleton for animation by rotating the **joint** angles to predetermined values. The position of a child **joint** changes according to the rotation of its parent and so the end point of a chain of **joints** can be determined from the angles and relative positions of the individual **joints** it contains.

#### Forward Rendering:

A **rendering path** that renders each object in one or more passes, depending on lights that affect the object. Lights themselves are also treated differently by **Forward Rendering**, depending on their settings and intensity. [More info](RenderTech-ForwardRendering.md)

#### FPS:

See **first person shooter**, **frames per second**.

#### frame:

A single image from a sequence of images that represent moving graphics. While your game is running, the **camera** in your game renders frames to the screen as fast as possible. May also refer to a frame from a video clip, or **sprite** animation frames. See **frames per second** (FPS).

#### frames per second:

The frequency at which consecutive frames are displayed in a running game. [More info](RenderingStatistics.md)

#### Fresnel Effect:

An effect representing the increase in reflectivity on objects when light hits at grazing angles.

#### frustum:

The region of 3D space that a perspective **camera** can see and render. In the **Scene** view, the frustum is represented by a truncated rectangular pyramid with its top at the **camera\_\_’s near** clipping plane\_\_ and its base at the **camera\_\_’s far** clipping plane\_\_. [More info](UnderstandingFrustum.md)

#### funnel:

In **Analytics**, a funnel is a linear sequence of standard or custom events that you expect a player to complete in order. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/introducing-funnels)

## G

#### game console:

A device that runs and displays video games.

#### game controller:

A device to control objects and characters in a game.

#### GameObject:

The fundamental object in Unity **scenes**, which can represent characters, props, scenery, **cameras**, waypoints, and more. A **GameObject**’s functionality is defined by the Components attached to it. [More info](class-GameObject.md)

#### geography:

Player segments based on country. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### GI Cache:

The cached intermediate files used when Unity precomputes lighting data. Unity keeps this cache to speed up computation. [More info](GICache.md)

#### Git dependency:

The Package Manager retrieves **Git dependencies** from a Git repository directly rather than from a package registry. **Git dependencies** use a Git URL reference instead of a version, and there’s no guarantee about the package quality, stability, validity, or even whether the version stated in its `package.json` file respects Semantic Versioning rules with regards to officially published releases of this package. [More info](upm-concepts.md#Git)

#### Gizmo:

A graphic overlay associated with a **GameObject** in a **Scene**, and displayed in the **Scene** View. Built-in **scene** tools such as the move tool are **Gizmos**, and you can create custom **Gizmos** using textures or scripting. Some **Gizmos** are only drawn when the **GameObject** is selected, while other **Gizmos** are drawn by the Editor regardless of which **GameObjects** are selected. [More info](GizmosMenu.md#GizmosIcons)

#### global illumination:

A group of techniques that model both direct and indirect lighting to provide realistic lighting results.

#### Gradle:

An Android build system that automates several build processes. This automation means that many common build errors are less likely to occur. [More info](android-gradle-overview.md)

#### Gravity Modifier:

A **Particle** System property that scales the gravity value set in the physics manager. A value of zero switches gravity off. [More info](PartSysMainModule.md)

## H

#### Halo:

The glowing light areas around light sources, used to give the impression of small dust **particles** in the air. [More info](class-Halo.md)

#### hard shadows:

A shadow property that produces shadows with a sharp edge. Hard shadows are not particularly realistic compared to Soft Shadows but they involve less processing, and are acceptable for many purposes.

#### HDR:

high dynamic range

#### HDRI:

high dynamic range image

#### heightmap:

A greyscale Texture that stores height data for an object. Each **pixel** stores the height difference perpendicular to the face that **pixel** represents.

#### high twist limit:

The higher limit of a Character **Joint**. [More info](class-CharacterJoint.md)

#### Hinge Joint:

A **joint** that groups together two **Rigidbody** components, constraining them to move like they are connected by a hinge. It is perfect for doors, but can also be used to model chains, pendulums and so on. [More info](class-HingeJoint.md)

#### HLAPI:

Abbreviation of **High Level API**.

#### host:

In a multiplayer network game without a dedicated server, one of the peers in the game acts as the center of authority for the game. This peer is called the “host”. It runs a server and a “local client”, while the other peers each run a “remote client”. [More info](multiplayer.md)

#### Human template:

A pre-defined bone-mapping. Used for matching bones from FBX files to the **Avatar**. [More info](class-HumanTemplate.md)

#### Humanoid animation:

An animation using humanoid skeletons. Humanoid models generally have the same basic structure, representing the major articulate parts of the body, head and limbs. This makes it easy to map animations from one humanoid skeleton to another, allowing **retargeting** and inverse **kinematics**. [More info](ConfiguringtheAvatar.md)

## I-J

#### IAP:

See **In App Purchase**

#### ignore file:

A special file used in many **Version Control** Systems which specifies files to be excluded from being placed under **version control**. In Unity projects there are a number of files which could be excluded from **version control**, and using an Ignore File is the best way to achieve this.

#### IL2CPP:

A Unity-developed scripting back-end which you can use as an alternative to Mono when building projects for some platforms. [More info](./scripting-backends-il2cpp.md)

#### Image control:

An Image control displays a non-interactive image to the user, such as a decoration and icon. You can change the image from a script to reflect changes in other controls. [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/script-Image.html)

#### immutable:

You cannot change the contents of an **immutable** (read-only) package. This is the opposite of **mutable**. Most packages are **immutable**, including packages downloaded from the package registry or by Git URL.

#### impressions:

The number of times ads are seen in your game. An impression is counted even if the ad is not completed. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### in app purchase:

Revenue from “micro-transactions” within a game. [More info](UnityIAP)

#### indirect dependency:

An **indirect**, or transitive dependency occurs when your project requests a package which itself “depends on” another package. For example, if your project depends on the `alembic@1.0.7` package which in turn depends on the `timeline@1.0.0` package, then your project has an **direct dependency** on Alembic and an **indirect dependency** on Timeline. [More info](upm-dependencies.md)

#### input key:

A key on a keyboard relating to the Input class. [More info](../ScriptReference/KeyCode.md)

#### Input Manager:

Settings where you can define all the different input axes, buttons and controls for your project. [More info](class-InputManager.md)

#### Inspector:

A Unity window that displays information about the currently selected **GameObject**, asset or **project settings**, allowing you to inspect and edit the values. [More info](UsingTheInspector.md)

#### Interactable:

A UI component property that determines whether the component can accept input. [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/script-Selectable.html)

#### interpolate:

See **Interpolation**

#### interpolation:

The process of calculating values in-between two defined values. Used in animation (between keyframes), physics (between physics time-steps), and multiplayer (between network updates)

#### inverse kinematics (IK):

The automatic calculation of **joint** angles (eg. the shoulder and elbow joint of an arm) so that the end point (eg. the hand) reaches a desired point in space. In contrast to **Forward Kinematics** [More info](InverseKinematics.md)

#### Irradiance Budget:

A **lightmap** property that determines the precision of the incoming light data used to light each texel in the **lightmap** [More info](class-LightmapParameters.md)

#### Irradiance Quality:

A **lightmap** property that sets the number of rays that are cast and used to compute which clusters affect a given output **lightmap** texel. [More info](class-LightmapParameters.md)

#### isometric projection:

A form of parallel projection where the dimensions of a **3D object** are projected onto a 2D plane, and the angles between all three axes are equal to each other. This form of projection is commonly used in isometric video games to simulate three-dimensional depth. [More info](tilemaps/work-with-tilemaps/isometric-tilemaps/isometric-tilemap-landing.md)

#### joint:

A physics component allowing a dynamic connection between **Rigidbody** components, usually allowing some degree of movement such as a hinge. [More info](Joints.md)

#### Joy Num:

An **Input Manager** property that defines which joystick will be used. [More info](class-InputManager.md)

## K

#### Keyframe Reduction:

A process that removes redundant **keyframes**. [More info](class-AnimationClip.md#AssetProperties)

#### keyframe:

A frame that marks the start or end point of a transition in an animation. Frames in between the **keyframes** are called inbetweens.

#### kinematics:

The geometry that describes the position and orientation of a character’s **joints** and bodies. Used by inverse **kinematics** to control character movement.

## L

#### Layer Mask:

A value defining which layers to include or exclude from an operation, such as rendering, **collision** or your own code. [More info](Layers.md)

#### layer:

Layers in Unity can be used to selectively opt groups of **GameObjects** in or out of certain processes or calculations. This includes **camera** rendering, lighting, physics **collisions**, or custom calculations in your own code. [More info](Layers.md)

#### Lens Flare:

A component that simulates the effect of lights refracting inside a **camera** lens. Use a **Lens Flare** to represent very bright lights or add atmosphere to your **scene**. [More info](class-LensFlare.md)

#### level of detail:

The *Level Of Detail* (LOD) technique is an optimization that reduces the number of triangles that Unity has to render for a **GameObject** when its distance from the **Camera** increases. [More info](LevelOfDetail.md)

#### Light Mode:

A Light property that defines the use of the Light. Can be set to Realtime, Baked and Mixed. [More info](LightModes.md)

#### Light Probe Group:

A component that enables you to add **Light Probes** to **GameObjects** in your **scene**. [More info](class-LightProbeGroup.md)

#### Light Probe Proxy Volume:

A component that allows you to use more lighting information for large dynamic **GameObjects** that cannot use baked **lightmaps** (for example, large Particle Systems or skinned Meshes). [More info](class-LightProbeProxyVolume.md)

#### Light Probe:

Light probes store information about how light passes through space in your **scene**. A collection of **light probes** arranged within a given space can improve lighting on moving objects and static **LOD** scenery within that space. [More info](LightProbes.md)

#### lightmap texture:

Final rendered texture containing lighting data. **Lightmap** textures store baked lighting information such as indirect illumination and shadowing, and are applied to meshes via their UVs.

#### Lightmap:

A pre-rendered texture that contains the effects of light sources on static objects in the **scene**. **Lightmaps** are overlaid on top of **scene** geometry to create the effect of lighting. [More info](Lightmapping.md)

#### Lightmapper:

A tool in Unity that bakes **lightmaps** according to the arrangement of lights and geometry in your **scene**. [More info](Lightmapping.md)

#### Line Renderer:

A component that takes an array of two or more points in 3D space and draws a straight line between each one. You can use a single **Line Renderer** component to draw anything from a simple straight line to a complex spiral. [More info](class-LineRenderer.md)

#### local package:

A *local* package already exists on the file system, usually outside the project folder. To install the package, notify the Package Manager of its location through the **Packages** window. [More info](upm-ui-local.md)

#### LOD Group:

A component to manage **level of detail** (LOD) for **GameObjects**. [More info](class-LODGroup.md)

#### LOD:

See **level of detail**

#### Loop Pose:

An **animation clip** setting that blends the end and start of an animation to create a seamless join. [More info](class-AnimationClip.md)

#### Low Twist Limit:

A **joint** property that sets the lower limit of a **joint**. [More info](class-CharacterJoint.md)

#### LTV:

(Lifetime Value) The estimated value of an average player over their lifetime with your application or game. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

## M

#### Main Editor Player:

The Unity Player that exists in the main Unity Editor.

#### managed plug-in:

A managed .NET assembly that is created with tools like Visual Studio for use in Unity. [More info](./plug-ins.md)

#### manifest:

There are two types of manifest files: **project manifest****s** and **package manifest****s**.

#### margin:

Minimum distance between UV charts within the same layout. Margin helps prevent texture bleeding between charts during filtering and mipmapping, and is usually specified in texels at the target resolution.

#### marker:

A Unity Profiler API structure that describes a CPU or GPU event, such as a button click. Each event marker appears as a vertical lines or label in the Profiler window. [More info](profiler-markers.md)

#### mask:

Can refer to a **Sprite** Mask, a **UI Mask**, or a **Layer Mask** [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/script-Mask.html)

#### Material:

An asset that defines how a surface should be rendered. [More info](class-Material.md)

#### MAU:

(Monthly Active Users) The number of players who started a session within the last 30 days. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### mcs:

The Mono C# compiler file format. [More info](platform-dependent-compilation.md)

#### Mesh Collider:

A free-form **collider** component which accepts a **mesh** reference to define its **collision** surface shape. [More info](class-MeshCollider.md)

#### Mesh Filter:

A **mesh** component that takes a **mesh** from your assets and passes it to the **Mesh** Renderer for rendering on the screen. [More info](class-MeshFilter.md)

#### Mesh Renderer:

A **mesh** component that takes the geometry from the **Mesh** Filter and renders it at the position defined by the object’s **Transform component**. [More info](class-MeshRenderer.md)

#### Mesh:

The main graphics primitive of Unity. Meshes make up a large part of your 3D worlds. Unity supports triangulated or Quadrangulated polygon meshes. Nurbs, Nurms, Subdiv surfaces must be converted to polygons. [More info](mesh.md)

#### minnow:

A player who has spent less than $5 in their lifetime. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### Mixed Lights:

Light components whose Mode property is set to Mixed. Some calculations for Mixed Lights are performed in advance, and some calculations for Mixed Lights are performed at runtime. The behavior of all Mixed Lights in a **Scene** is determined by the **Scene**’s Lighting Mode. [More info](LightModes-landing.md)

#### mixed reality:

Mixed Reality (MR) combines its own virtual environment with the user’s real-world environment and allows them to interact with each other.

#### model file:

A file containing a 3D data, which may include definitions for meshes, bones, animation, materials and textures. [More info](3D-formats.md)

#### model:

A 3D model representation of an object, such as a character, a building, or a piece of furniture. [More info](3D-formats.md)

#### Mono:

A **scripting backend** used in Unity. [More info](./scripting-backends-il2cpp.md)

#### MonoDevelop:

An integrated development environment (IDE) supplied with Unity 2017.3 and previous versions. From Unity 2018.1 onwards, **MonoDevelop** is replaced by Visual Studio. [More info](https://www.monodevelop.com/)

#### MR:

Mixed Reality

#### muscle definition:

This allows you to have more intuitive control over the character’s skeleton. When an **Avatar** is in place, the Animation system works in muscle space, which is more intuitive than bone space. [More info](MuscleDefinitions.md)

#### mutable:

You can change the contents of a **mutable** package. This is the opposite of **immutable**. Only **Local package****s** and **Embedded package****s** are **mutable**.

## N

#### narrow phase collision detection:

A **collision** detection phase that determines whether the pairs of objects found in the broad phase will actually collide. It then computes the contact points for those pairs and gives them to the solver to use when solving **collisions**. [More info](http://planning.cs.uiuc.edu/node214.html)

#### native plug-in:

A platform-specific native code library that is created outside of Unity for use in Unity. Allows you can access features like OS calls and third-party code libraries that would otherwise not be available to Unity. [More info](./plug-ins.md)

#### NavMesh:

A **mesh** that Unity generates to approximate the walkable areas and obstacles in your environment for path finding and AI-controlled navigation. [More info](https://docs.unity3d.com/Packages/com.unity.ai.navigation@latest/index.html?subfolder=/manual/NavInnerWorkings.html%23walkable-areas)

#### near clipping plane:

A plane that limits how close a **camera** can see from its current position. The plane is perpendicular to the **camera**’s forward (Z) direction. [More info](CamerasOverview.md)

#### networking:

The Unity system that enables multiplayer gaming across a computer network. [More info](multiplayer.md)

#### never monetized:

Players who have never spent real currency. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### new users:

Users who played your game for the first time. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### normal map:

A type of **Bump Map** texture that allows you to add surface detail such as bumps, grooves, and scratches to a model which catch the light as if they are represented by real geometry.

#### normal:

The direction perpendicular to the surface of a **mesh**, represented by a Vector. Unity uses normals to determine object orientation and apply shading. [More info](scripting-vectors.md)

#### number of unverified transactions:

The total number of IAP transactions, whether or not they have been verified. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### number of users:

The cumulative number of unique players over the last 90 days. Users who have not played in more than 90 days are removed from the count. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### number of verified transactions:

IAP transactions that have been verified through the appropriate app store. IAP verification is currently supported by the Apple App Store and the Google Play Store. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

## O

#### object:

See **GameObject**.

#### occlusion culling:

A process that disables rendering **GameObjects** that are hidden (occluded) from the view of the **camera**. [More info](OcclusionCulling.md)

#### ODR:

On-demand resources (ODR) is a feature available for the iOS and tvOS platforms, from version 9.0 of iOS and tvOS onwards. It allows you to reduce the size of your application by separating the core assets (those that are needed from application startup) from assets which may be optional, or which appear in later levels of your game. [More info](AppThinning.md)

#### orthographic 3D:

A common type of **camera** view used in games that give you a bird’s-eye view of the action, and is sometimes called “2.5D”. [More info](2Dor3D.md)

#### overhead:

The amount of supporting code that the Profiler needs to run. This overhead might affect the performance of your application.

## P

#### package manifest:

Each package has a *manifest*, which provides information about the package to the Package Manager. The manifest contains information such as the name of the package, its version, a description for users, dependencies on other packages (if any), and other details. [More info](upm-manifestPkg.md)

#### package:

A container that stores various types of features and assets for Unity, including Editor or Runtime tools and libraries, Asset collections, and project templates. Packages are self-contained units that the Unity Package Manager can share across Unity projects. Most of the time these are called *packages*, but occasionally they are called **Unity Package Manager (UPM) packages**. The [Unity Package Manager](upm-ui.md) (UPM) can display, add, and remove packages from your project. These packages are native to the Unity Package Manager and provide a fundamental method of delivering Unity functionality. However, the Unity Package Manager can also display [Asset Store packages](AssetStorePackages.md) that you downloaded from the **Asset Store**. [More info](Packages.md)

#### padding:

Minimum distance between different UV layouts in the atlas, specified in texels. Padding prevents cross-object bleeding within a UV atlas to accommodate filtering and mip levels.

#### parent:

An object that contains child objects in a hierarchy. When a **GameObject** is a Parent of another **GameObject**, the Child **GameObject** will move, rotate, and scale exactly as its Parent does. You can think of parenting as being like the relationship between your arms and your body; whenever your body moves, your arms also move along with it. [More info](class-Transform.md#parent)

#### particle system:

A component that simulates fluid entities such as liquids, clouds and flames by generating and animating large numbers of small 2D images in the **scene**. [More info](class-ParticleSystem.md)

#### particle:

A small, simple image or **mesh** that is emitted by a **particle** system. A **particle** system can display and move **particles** in great numbers to represent a fluid or amorphous entity. The effect of all the **particles** together creates the impression of the complete entity, such as smoke. [More info](class-ParticleSystem.md)

#### percentage of population:

Your player population as a percentage. Typically only useful when combined with a segment. Calculated as the percentage of the Number of Users metric who are members of a specified segment. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### Perforce:

A **version control** system for file change management. [More info](perForceIntegration.md)

#### PhotoCapture:

An API that enables you to take photos from a HoloLens web **camera** and store them in memory or on disk. [More info](../ScriptReference/Windows.WebCam.PhotoCapture.md)

#### physically based shading:

An advanced lighting model that simulates the interactions between materials and light in a way that mimics reality. [More info](shader-StandardShader.md)

#### physics engine:

A system that simulates aspects of physical systems so that objects can accelerate correctly and be affected by **collisions**, gravity and other forces. [More info](PhysicsSection.md)

#### Physics Material 2D:

Use this to adjust the friction and bounce that occurs between 2D physics objects when they collide [More info](2d-physics/physics-material-2d-reference.md)

#### Physics Material:

A physics asset for adjusting the friction and bouncing effects of colliding objects. [More info](class-PhysicsMaterial.md)

#### ping pong:

To repeatedly play an animation to the end, then in reverse back to the beginning, in a loop.

#### pixel:

The smallest unit in a computer image. **Pixel** size depends on your screen resolution. **Pixel** lighting is calculated at every screen **pixel**. [More info](ShadowPerformance.md)

#### Play On Awake:

Set this to true to make an **Audio Source** start playing on awake [More info](class-AudioClip.md)

#### Playable Graph:

An API for controlling **Playables**. **Playable Graphs** allow you to create, connect and destroy **playables**. [More info](Playables-Graph.md)

#### Playables:

An API that provides a way to create tools, effects or other gameplay mechanisms by organizing and evaluating data sources in a tree-like structure known as the PlayableGraph. [More info](Playables.md)

#### player log:

The .log file created by a Standalone Player that contains a record of events, such as script execution times, the compiler version, and AssetImport time. Log files can help diagnose problems. [More info](log-files.md#player)

#### Player Settings:

Settings that let you set various player-specific options for the final game built by Unity. [More info](class-PlayerSettings.md)

#### plug-in:

A set of code created outside of Unity that creates functionality in Unity. There are two kinds of **plug-ins** you can use in Unity: Managed **plug-ins** (managed .NET assemblies created with tools like Visual Studio) and Native **plug-ins** (platform-specific native code libraries). [More info](./plug-ins.md)

#### post-processing:

A process that improves product visuals by applying filters and effects before the image appears on screen. You can use post-processing effects to simulate physical camera and film properties, for example Bloom and Depth of Field. [More info](PostProcessingOverview.md) post processing, postprocessing, postprocess

#### prefab:

An asset type that allows you to store a **GameObject** complete with components and properties. The **prefab** acts as a template from which you can create new object instances in the **scene**. [More info](Prefabs.md)

#### preview package:

A *preview* package is in development and not yet ready for production. A package in preview might be at any stage of development, from the initial stages to near completion.

#### Profiler category:

Identifies the workload data for a Unity subsystem (for example, Rendering, Scripting and Animation categories). Unity applies color-coding to categories to visually distinguish between the types of data in the **Profiler** window.

#### profiler counter:

Placed in code with the ProfilerCounter API to track metrics, such as the number of enemies spawned in your game. [More info](https://docs.unity3d.com/Packages/com.unity.profiling.core@latest/index.html?subfolder=/manual/profilercounter-guide.html)

#### profiler marker:

Placed in code to describe a CPU or GPU event that is then displayed in the Unity Profiler window. Added to Unity code by default, or you can use [ProfilerMarker API](https://docs.unity3d.com/Packages/com.unity.profiling.core@latest/index.html?subfolder=/manual/profilermarker-guide.html) to add your own custom markers. [More info](profiler-markers.md)

#### profiler sample:

A set of data associated with a **Profiler marker**, that the Profiler has recorded and collected.

#### Profiler:

A window that helps you to optimize your game. It shows how much time is spent in the various areas of your game. For example, it can report the percentage of time spent rendering, animating, or in your game logic. [More info](Profiler.md)

#### Progressive Web App:

A software application that’s delivered through the web. It uses certain browser features to create a user experience on par with a native application. [More info](https://developer.mozilla.org/en-US/docs/Web/Progressive_web_apps)

#### project manifest:

Each Unity project has a *project manifest*, which acts as an entry point for the Package Manager. This file must be available in the `<project>/Packages` directory. The Package Manager uses it to configure many things, including a list of dependencies for that project, as well as any package repository to query for packages. [More info](upm-manifestPrj.md)

#### Project Settings:

A broad collection of settings which allow you to configure how Physics, Audio, **Networking**, Graphics, Input and many other areas of your project behave. [More info](comp-ManagerGroup.md)

#### Project window:

A window that shows the contents of your `Assets` folder (Project tab) [More info](ProjectView.md)

#### project:

In Unity, you use a project to design and develop a game. A project stores all of the files that are related to a game, such as the asset and **Scene** files. [More info](2Dor3D.md)

#### Property Drawer:

A Unity feature that allows you to customize the look of certain controls in the **Inspector** window by using attributes on your **scripts**, or by controlling how a specific Serializable class should look [More info](editor-PropertyDrawers.md)

#### pseudo-depth:

A visual simulation of three-dimensional characteristics on a two-dimensional object or environment. This effect is sometimes called “2.5D”. [More info](tilemaps/work-with-tilemaps/isometric-tilemaps/isometric-tilemap-landing.md)

## Q

#### Quad:

A primitive object that resembles a plane but its edges are only one unit long, it uses only 4 vertices, and the surface is oriented in the XY plane of the local coordinate space. [More info](PrimitiveObjects.md)

#### Quaternion:

Unity’s standard way of representing rotations as data. When writing code that deals with rotations, you should usually use the **Quaternion** class and its methods. [More info](QuaternionAndEulerRotationsInUnity.md)

## R

#### rasterization:

The process of generating an image by calculating **pixels** for each polygon or triangle in the geometry. This is an alternative to **ray tracing**.

#### Raw Image:

A Visual UI Component that displays a non-interactive image to the user. This can be used for decoration, icons, and so on, and the image can also be changed from a script to reflect changes in other controls. [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/script-RawImage.html)

#### ray tracing:

The process of generating an image by tracing out rays from the Camera through each pixel and recording the color contribution at the hit point. This is an alternative to rasterization. raytracing

#### Realtime Lights:

Light components whose Mode property is set to Realtime. Unity calculates and updates the lighting of **Realtime Lights** every frame at runtime. No **Realtime Lights** are precomputed. [More info](LightModes-introduction.md#realtime)

#### Reflection Probe:

A rendering component that captures a spherical view of its surroundings in all directions, rather like a **camera**. The captured image is then stored as a **Cubemap** that can be used by objects with reflective materials. [More info](class-ReflectionProbe.md)

#### Relative Joint 2D:

A 2D **joint** that allows two game objects controlled by **Rigidbody** physics to maintain in a position based on each other’s location. Use this **joint** to keep two objects offset from each other, at a position and angle you decide [More info](2d-physics/joints/relative-joint-2d-reference.md)

#### Remote Settings:

Remote settings are game variables that you can set remotely on your **Analytics** Dashboard. [More info](https://docs.unity.com/ugs/en-us/manual/remote-config/manual/WhatsRemoteConfig)

#### render pipeline:

A series of operations that take the contents of a **Scene**, and displays them on a screen. Unity lets you choose from pre-built **render pipelines**, or write your own. [More info](render-pipelines.md)

#### Render Texture:

A special type of Texture that is created and updated at runtime. To use them, first create a new **Render Texture** and designate one of your **Cameras** to render into it. Then you can use the **Render Texture** in a Material just like a regular Texture. [More info](class-RenderTexture.md)

#### rendering layer:

A layer that defines how specific effects are applied across different objects. Rendering layers don’t define draw order. They’re selection groups you assign objects to. Rendering layers let lights, decals, shadows, and custom passes include or ignore specific objects.

#### rendering layer mask:

A bitmask that aggregates multiple rendering layers. By assigning one or more rendering layers to an object’s MeshRenderer or an effect such as light, decals, or APVs, you can control how and where Unity applies the effect in the **scene**.

#### rendering path:

The technique that a **render pipeline** uses to render graphics. Choosing a different **rendering path** affects how lighting and shading are calculated. Some **rendering paths** are more suited to different platforms and hardware than others. [More info](RenderingPaths.md)

#### retargeting:

Applying animations created for one model to another. [More info](Retargeting.md)

#### rig:

A skeletal hierarchy of **joints** for a **mesh**. [More info](FBXImporter-Rig.md)

#### rigging:

The process of building a skeleton hierarchy of **joints** for your **mesh**. Performed with an external tool, such as Blender or Autodesk Maya. [More info](UsingHumanoidChars.md)

#### Rigidbody:

A component that allows a **GameObject** to be affected by simulated gravity and other forces. [More info](class-Rigidbody.md)

#### Root Motion:

Motion of character’s **root node**, whether it’s controlled by the animation itself or externally. [More info](RootMotion.md)

#### Root node:

A transform in an animation hierarchy that allows Unity to establish consistency between **Animation clips** for a generic model. It also enables Unity to properly blend between Animations that have not been authored “in place” (that is, where the whole Model moves its world position while animating). [More info](class-AnimationClip.md#AllClipProperties)

#### Root Transform:

The Transform at the top of a hierarchy of Transforms. In a **Prefab**, the **Root Transform** is the topmost Transform in the **Prefab**. In an animated humanoid character, the **Root Transform** is a projection on the Y plane of the Body Transform and is computed at run time. At every frame, a change in the **Root Transform** is computed, and then this is applied to the **GameObject** to make it move. [More info](RootMotion.md)

## S

#### Scene View:

An interactive view into the world you are creating. You use the **Scene** View to select and position scenery, characters, **cameras**, lights, and all other types of Game Object. [More info](UsingTheSceneView.md)

#### Scene:

A **Scene** contains the environments and menus of your game. Think of each unique **Scene** file as a unique level. In each **Scene**, you place your environments, obstacles, and decorations, essentially designing and building your game in pieces. [More info](CreatingScenes.md)

#### Scripting Backend:

A framework that powers scripting in Unity. Unity supports three different **scripting backends** depending on target platform: Mono, .NET and **IL2CPP**. Universal Windows Platform, however, supports only two: .NET and **IL2CPP**. [More info](scripting-backends.md)

#### Scripting Event:

A way of allowing a user-driven callback to persist from edit time to run time without the need for additional programming and script configuration [More info](UIE-Runtime-Event-System.md)

#### Scripts:

A piece of code that allows you to create your own Components, trigger game events, modify Component properties over time and respond to user input in any way you like. [More info](creating-scripts.md)

#### ScrollView:

A UI Control which displays a large set of Controls in a viewable area that you can see by using a scrollbar. [More info](UIE-uxml-element-ScrollView.md)

#### segment:

Segments are subsets of your player base, split apart by key differentiators. Viewing metrics and events by segment can reveal differences in-game behavior between different groups. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### Self Collision:

A cloth property that prevents cloth from penetrating itself. [More info](class-Cloth.md)

#### Services:

A Unity facility that provides a growing range of complimentary services to help you make games and engage, retain and monetize audiences. [More info](UnityServices.md)

#### session:

A single play or usage period. A new session is counted when a player launches your game or brings a suspended game to the foreground after 30 minutes of inactivity. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### sessions per user:

The average number of sessions per person playing on a given day. Also known as Average Number of Sessions per **DAU**. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### Shader object:

An instance of the **Shader** class, a **Shader** object is container for **shader** programs and GPU instructions, and information that tells Unity how to use them. Use them with materials to determine the appearance of your **scene**. [More info](shader-objects.md)

#### shader variant:

A verion of a **shader** program that Unity generates according to a specific combination of **shader** keywords and their status. A **Shader** object can contain multiple **shader** variants. [More info](shader-variants.md)

#### shader:

A program that runs on the GPU. [More info](Shaders.md)

#### ShaderLab:

Unity’s language for defining the structure of **Shader** objects. [More info](SL-Shader.md)

#### Shadow:

A UI component that adds a simple outline effect to graphic components such as Text or Image. It must be on the same **GameObject** as the graphic component. [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/script-Shadow.html)

#### Shadowmask:

A **shadowmask** texture uses the same UV layout and resolution as its corresponding **lightmap** texture. [More info](lighting-mode.md#shadowmask)

#### Skinning:

The process of binding bone **joints** to the vertices of a character’s **mesh** or ‘skin’. Performed with an external tool, such as Blender or Autodesk Maya. [More info](UsingHumanoidChars.md)

#### Skybox:

A special type of Material used to represent skies. Usually six-sided. [More info](sky-landing.md)

#### Soft Particles:

Particles that create semi-transparent effects like smoke, fog or fire. Soft **particles** fade out as they approach an opaque object, to prevent intersections with the geometry. [More info](shader-StandardParticleShaders.md)

#### Soft Shadows:

A shadow property that produces shadows with a soft edge. Soft shadows are more realistic compared to Hard Shadows and tend to reduce the “blocky” aliasing effect from the shadow map, but they require more processing.

#### Spatial Mapping:

The process of mapping real-world surfaces into the virtual world.

#### specular color:

The color of a specular highlight.

#### specular highlight

A bright spot of light that appears on the surface of shiny objects when illuminated.

#### speculative continuous collision detection:

A **collision** detection method that inflates broad phase AABB of moving objects according to their velocities. This enables support for effects like rotations. [More info](ContinuousCollisionDetection.md)

#### Sphere Collider:

A sphere-shaped **collider** component that handles **collisions** for **GameObjects** like balls or other things that can be roughly approximated as a sphere for the purposes of physics. [More info](class-SphereCollider.md)

#### Spring Joint:

A **joint** type that connects two **Rigidbody** components together but allows the distance between them to change as though they were connected by a spring. [More info](class-SpringJoint.md)

#### Sprite Atlas:

**Graphics:** A utility that packs several **sprite** textures tightly together within a single texture known as an atlas. [More info](sprite/atlas/v2/v2-landing). **2D:** A texture that is composed of several smaller textures. Also referred to as a texture atlas, image **sprite**, **sprite** sheet or packed texture. [More info](sprite/atlas/atlas-landing.md).

#### Sprite Mask:

A texture which defines which areas of an underlying image to reveal or hide. [More info](sprite/mask/mask-landing.md)

#### Sprite Renderer:

A component that lets you display images as **Sprites** for use in both 2D and 3D **scenes**. [More info](sprite/renderer/sprite-renderer-reference.md)

#### Sprite:

A 2D graphic objects. If you are used to working in 3D, **Sprites** are essentially just standard textures but there are special techniques for combining and managing **sprite** textures for efficiency and convenience during development. [More info](sprite/sprite-landing.md)

#### standard event:

Standard events are application-specific events that you dispatch in response to important player actions or milestones. **Standard events** have standardized names and defined parameter lists. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/events)

#### State Machine Behaviour:

A script that attaches to a state within a **state machine** to control what happens when the **state machine** enters, exits or remains within a state, such as play sounds as states are entered. [More info](StateMachineBehaviours.md)

#### State Machine:

The set of states in an **Animator Controller** that a character or animated **GameObject** can be in, along with a set of transitions between those states and a variable to remember the current state. The states available will depend on the type of gameplay, but typical states include things like idling, walking, running and jumping. [More info](StateMachineBasics.md)

#### Static Batching:

A technique Unity uses to draw **GameObjects** on the screen that combines static (non-moving) **GameObjects** into big Meshes, and renders them in a faster way. [More info](DrawCallBatching.md)

#### static receiver:

A static **GameObject** that is receiving a shadow from another static or dynamic **GameObject** [More info](StaticObjects.md)

#### stencil buffer:

A memory store that holds an 8-bit per-pixel value. In Unity, you can use a **stencil buffer** to flag **pixels**, and then only render to **pixels** that pass the stencil operation. [More info](class-RenderTexture.md)

#### stencil masking meshes:

Overflow hidden with either rounded corners or vector image background.

#### Sticky Factor:

An estimate of how compelling a game is to its players. A high “**sticky factor**” means that players stick with an app over time. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### Surface Shader:

A streamlined way of writing **shaders** for the Built-in **Render Pipeline**. [More info](SL-SurfaceShaders.md)

#### Swing Axis:

A **joint** property that defines the axis around which the **joint** can swing. [More info](class-CharacterJoint.md)

#### Swing Limit:

A **joint** property that limits the rotation around one element of the defined **Swing Axis**. [More info](class-CharacterJoint.md)

## T

#### T-pose:

The pose in which the character has their arms straight out to the sides, forming a “T”. The required pose for the character to be in, in order to make an **Avatar**.

#### tag:

**Multiplayer:** A reference word that you can assign to a **Virtual Player** or the **main Editor Player** to use in a script. For example, you could define a “Red team” ot “Blue team” tag to assign Players to each team. [More info](https://docs.unity3d.com/Packages/com.unity.multiplayer.playmode@latest?subfolder=/manual/player-tags/player-tags.html). **Scripting:** A reference word which you can assign to one or more **GameObjects** to help you identify **GameObjects** for scripting purposes. For example, you might define and “Edible” Tag for any item the player can eat in your game. [More info](Tags.md)

#### target matching:

A scripting function that allows you to move characters in such a way that a hand or foot lands in a certain place at a certain time. For example, the character may need to jump across stepping stones or jump and grab an overhead beam. [More info](TargetMatching.md)

#### target position:

A **joint** property to set the **target position** that the **joint**’s drive force should move it to. [More info](class-ConfigurableJoint.md)

#### target velocity:

A **joint** property to set the desired velocity with which the **joint** should move to the **Target Position** under the drive force. [More info](class-ConfigurableJoint.md)

#### Terrain Collider:

A terrain-shaped **collider** component that handles **collisions** for **collision** surface with the same shape as the **Terrain** object it is attached to. [More info](class-TerrainCollider.md)

#### Terrain:

The landscape in your **scene**. A **Terrain** **GameObject** adds a large flat plane to your **scene** and you can use the **Terrain**’s **Inspector** window to create a detailed landscape. [More info](terrain-UsingTerrains.md)

#### Test Framework,Test Runner:

The Test Framework package (formerly called the Test Runner) is a Unity tool that tests your code in both Edit mode and Play mode, and also on target platforms such as Standalone, Android, or iOS. [More info](https://docs.unity3d.com/Packages/com.unity.test-framework@latest)

#### Text Input Field:

A field that allows the user to input a Text string [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/script-InputField.html)

#### Text Mesh:

A **Mesh** component that displays a Text string [More info](class-TextMesh.md)

#### Text:

A non-interactive piece of text to the user. This can be used to provide captions or labels for other GUI controls or to display instructions or other text. [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/script-Text.html)

#### TextField control:

A **TextField control** displays a non-interactive piece of text to the user, such as a caption, label for other GUI controls, or instruction. [More info](gui-Controls.md)

#### texture compression:

3D Graphics hardware requires Textures to be compressed in specialized formats which are optimized for fast Texture sampling. [More info](texture-choose-format-by-platform.md)

#### texture format:

A file format for handling textures during real-time rendering by 3D graphics hardware, such as a graphics card or mobile device. [More info](texture-choose-format-by-platform.md)

#### Texture Import Inspector:

An **Inspector** that allows you to define how your images are imported from your project’s `Assets` folder into the Unity Editor. [More info](class-TextureImporter.md)

#### texture overrides:

Platform-specific settings that allow you to set the resolution, file size with associated memory size requirements, **pixel** dimensions, and quality of your textures for each target platform. [More info](texture-choose-format-by-platform.md)

#### texture space:

The coordinate system that defines a texture’s bounds.

#### texture:

An image used when rendering a **GameObject**, **Sprite**, or UI element. Textures are often applied to the surface of a **mesh** to give it visual detail. [More info](class-TextureImporter.md)

#### tile:

A simple class that allows a **sprite** to be rendered on a **Tilemap**. [More info](tilemaps/tiles-for-tilemaps/scriptable-tiles/scriptable-tiles-landing.md)

#### Tilemap:

A **GameObject** that allows you to quickly create 2D levels using tiles and a grid overlay. [More info](tilemaps/work-with-tilemaps/tilemap-reference.md)

#### Time Manager:

A Unity Settings Manager that lets you set a number of properties that control timing within your game. [More info](class-TimeManager.md)

#### Toggle:

A checkbox that allows the user to switch an option on or off. [More info](UIE-uxml-element-Toggle.md)

#### tonemapping:

The process of remapping **HDR** values of an image into a range suitable to be displayed on screen. [More info](PostProcessingOverview.md)

#### toolbar:

A row of buttons and basic controls at the top of the Unity Editor that allows you to interact with the Editor in various ways (e.g. scaling, translation). [More info](Toolbar.md)

#### total daily playing time per active user:

The average playing time of people playing on a given day. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### total daily playing time:

The cumulative playing time of all people playing on a given day. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### total IAP revenue:

The **total IAP revenue**, including revenue from both verified and unverified transactions. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### total sessions today:

The total number of sessions by all people playing on a given day. Also known as Total Sessions. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### total verified revenue:

Revenue from Unity Ads and verified IAP transactions. IAP verification is currently supported by the Apple App Store and the Google Play Store. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### Trail Renderer:

A visual effect that lets you to make trails behind **GameObjects** in the **Scene** as they move. [More info](class-TrailRenderer.md)

#### Transform component:

A **Transform component** determines the Position, Rotation, and Scale of each object in the **scene**. Every **GameObject** has a Transform. [More info](class-Transform.md)

#### transition:

The blend from one state to another in a **state machine**, such as transitioning a character from a walk to a jog animation. Transitions define how long the blend between states should take, and the conditions that activate the blend. [More info](class-Transition.md)

#### Translate DoF:

The three degrees-of-freedom associated with translation (movement in X,Y & Z) as opposed to rotation.

#### Tree:

A **GameObject** and associated Component that allows you to add tree assets to your **Scene**. You can add branch levels and leaves to trees in the Tree **Inspector** window. [More info](class-Tree.md)

## U

#### UI Mask:

A **visual component** that lets you restrict images from view to only a small section of an image. For instance, you can apply a Mask to a Panel UI element to block all child images from view. [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/script-Mask.html)

#### Unity Analytics:

A data platform that provides **analytics** for your Unity game. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/overview)

#### Unity Build Automation:

See **Build Automation**

#### Unity Cloud Diagnostics:

A suite of cloud-enabled tools that help you collect and identify issues that users encounter with your apps. [More info](https://docs.unity.com/ugs/manual/cloud-diagnostics/manual/CloudDiagnostics/WelcometoCloudDiagnostics)

#### Unity IAP:

Abbreviation of Unity **In App Purchase**

#### Unity Remote:

A downloadable app designed to help with Android, iOS and tvOS development. The app connects with Unity while you are running your project in Play Mode from the Unity Editor. [More info](UnityRemote5.md)

#### Unity unit:

The unit size used in Unity projects. By default, 1 **Unity unit** is 1 meter. To use a different scale, set the Scale Factor in the Import Settings when importing assets.

#### unknown gender:

Players to whom you have assigned Gender.Unknown. (Players whose gender has not been reported at all are not included in this segment.) [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### unverified IAP revenue:

IAP revenue from sources that do not support verification and from transactions that failed verification. Transactions can fail verification because they are fraudulent or because of missing or malformed information. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### UPM package:

A **Package** managed by the **Unity Package Manager**. Refer to **Packages**.

#### UV atlas:

Collection of multiple UV layouts packed together. A UV atlas combines several meshes’ layouts into a single texture space to reduce draw calls and optimize baking workflows like lightmapping.

#### UV chart:

Individual unwrapped section within a UV layout. UV charts are contiguous islands of faces that share a continuous mapping. Their placement determines seams, distortion, and packing efficiency. Also known as a UV island.

#### UV layout:

The 2D representation of a single **mesh**’s UV coordinates. A UV layout defines how the **mesh**’s surface maps to texture space and influences **lightmap** unwrapping, texture sampling, and packing.

## V

#### VBlank:

Vertical blanking interval (VBlank) is the time between the end of the final visible line of a frame and the beginning of the first visible line of the next frame. This is the refresh interval as defined by a screen’s refresh rate.

#### vector field:

A 3D texture, where each value represents a directional force that is applied to **particles** as they move through the field. A vector field is created in 3D animation software, such as Houdini. [More info](https://en.wikipedia.org/wiki/Vector_field)

#### velocity:

A vector that defines the speed and direction of motion of a Rigidbody

#### verified IAP revenue:

Revenue from verified IAP transactions. IAP verification is currently supported by the Apple App Store and the Google Play Store. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### verified package:

When a package passes release cycle testing for a specific version of Unity, it receives the *Verified For* designation. This means that these packages are guaranteed to work with the designated version of Unity.

#### verified paying users:

Players who made verified IAP purchases. IAP verification is currently supported by the Apple App Store and the Google Play Store. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### version control:

A system for managing file changes. You can use Unity in conjunction with most common **version control** tools, including **Perforce**, Git, Mercurial and PlasticSCM. [More info](VersionControl.md)

#### vertex shader:

A program that runs on each vertex of a 3D model when the model is being rendered. [More info](writing-shader-writing-shader-programs-hlsl.md)

#### VideoCapture:

A Unity API that allows you to record videos directly to the file system in the MP4 format. [More info](../ScriptReference/Windows.WebCam.VideoCapture.md)

#### viewport:

The user’s visible area of an app on their screen.

#### Virtual Player:

A Unity Player that exists separately from the **main Editor Player**. Use a **Virtual Player** to test multiplayer gameplay on the same device without the need to create a build. [More info](https://docs.unity3d.com/Packages/com.unity.multiplayer.playmode@latest?subfolder=/manual/virtual-players/virtual-players.html)

#### virtual reality:

Virtual Reality (VR) immerses users in an artificial 3D world of realistic images and sounds, using a headset and motion tracking. [More info](VROverview.md)

#### visual component:

A component that enables you to easily create GUI-specific functionality. [More info](https://docs.unity3d.com/Packages/com.unity.ugui@latest/index.html?subfolder=/manual/UIVisualComponents.html)

#### visual element:

A node of a **visual tree** that instantiates or derives from the C# [`VisualElement`](../ScriptReference/UIElements.VisualElement.md) class. You can style the look, define the behaviour, and display it on screen as part of the UI. [More info](UIE-VisualTree.md)

#### visual tree:

An object graph, made of lightweight nodes, that holds all the elements in a window or panel. It defines every UI you build with the UI Toolkit.

#### VR:

Virtual Reality [More info](VROverview.md)

#### VSync:

Vertical synchronization (VSync) is a display setting that caps a game’s frame rate to match the refresh rate of a monitor, to prevent image tearing. [More info](https://en.wikipedia.org/wiki/Screen_tearing#V-sync)

## W

#### WebGL:

A JavaScript API that renders 2D and 3D graphics in a web browser. The Unity Web build option allows Unity to publish content as JavaScript programs which use HTML5 technologies and the **WebGL** rendering API to run Unity content in a web browser. [More info](webgl.md)

#### whales:

Players who have spent at least $20 in their lifetime. [More info](https://docs.unity.com/ugs/en-us/manual/analytics/manual/data-explorer-v2)

#### Wheel Collider:

A special **collider** for grounded vehicles. It has built-in **collision** detection, wheel physics, and a slip-based tire friction model. It can be used for objects other than wheels, but it is specifically designed for vehicles with wheels. [More info](class-WheelCollider.md)

#### wind zone:

A **GameObject** that adds the effect of wind to your **terrain**. For instance, Trees within a **wind zone** will bend in a realistic animated fashion and the wind itself will move in pulses to create natural patterns of movement among the tree. [More info](class-WindZone.md)

#### world:

The area in your **scene** in which all objects reside. Often used to specify that coordinates are world-relative, as opposed to object-relative.

#### WYSIWYG:

What You See Is What You Get. A term used to describe a system where the user interface closely resembles the final output.

## X-Z

#### XR:

An umbrella term encompassing **Virtual Reality** (VR), **Augmented Reality** (AR) and **Mixed Reality** (MR) applications. Devices supporting these forms of interactive applications can be referred to as **XR** devices. [More info](XR.md)

#### zoom:

A **camera** control that lets you scale the view on your screen. To zoom a **camera** in the Unity Editor, press Alt + right click and drag. [More info](SceneViewNavigation.md)