# Animation state machine

It’s common for a character or a **GameObject** to have several animations for the different actions it performs in a game. For example, a character might breath and sway slightly when idle, walk when commanded, and raise their arms when they fall from a platform. A sliding door might open, close, or jam.

Mecanim uses a **state machine** to arrange these actions. A state machine is a graph of nodes and connecting lines that resembles a flowchart. A state machine plays the animation linked to the current action and determines the next action. You can create a state machine for each character and GameObject in your **scene**.

The following topics provide more details on Mecanim’s state machine:

* [State Machine Basics](StateMachineBasics.md)
* [Animation States](class-State.md)
* [Animation Parameters](AnimationParameters.md)
* [State Machine Transitions](StateMachineTransitions.md)
* [Animation Transitions](class-Transition.md)
* [Animation Blend Trees](class-BlendTree.md)
* [State Machine Behaviours](StateMachineBehaviours.md)
* [Sub-State Machines](NestedStateMachines.md)
* [Animation Layers](AnimationLayers.md)
* [State Machine Solo and Mute](AnimationSoloMute.md)
* [Target Matching](TargetMatching.md)