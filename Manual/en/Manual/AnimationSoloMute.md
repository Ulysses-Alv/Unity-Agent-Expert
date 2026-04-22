# State Machine Solo and Mute

In complex **state machines**, it is useful to preview the operation of some parts of the machine separately. For this, you can use the **Mute** and **Solo** functionality:

* **Mute** disables a transition.
* **Solo** plays only that transition.

You can set up multiple **Solo** transitions to play only those that have Solo enabled. If one transition has **Solo** enabled, Unity enables **Mute** on the other transitions. If both **Solo** and **Mute** are enabled, then **Mute** takes precedence.

You can set up **Mute** and **Solo** states either from the **Transition Inspector**, or the **State Inspector**, which provides an overview of all transitions from that state.

![In the State Inspector, transitions from State 0 to State A and State B are Solo. Transitions from State 0 to State C and State D are muted.](../uploads/Main/MecanimSoloMuteInspector.png)

In the State Inspector, transitions from State 0 to State A and State B are Solo. Transitions from State 0 to State C and State D are muted.

![In the Animator window, green transition arrows indicate solo states. Red transition arrows indicate muted states.](../uploads/Main/MecanimSoloMuteGraph.png)

In the Animator window, green transition arrows indicate solo states. Red transition arrows indicate muted states.

In the example above, if you are in `State 0`, only transitions to `State A` and `State B` are available. The muted transitions, `State C` and `State D`, are disabled.

Known issues:

* The Controller graph doesn’t always reflect the internal Mute states of the engine.