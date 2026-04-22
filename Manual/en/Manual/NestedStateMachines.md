# Sub-State Machines

It is common for a character to have complex actions that consist of a number of stages. Rather than handle the entire action with a single state, it makes sense to identify the separate stages and use a separate state for each. For example, a character may have an action called “Trickshot” where it crouches to take a steady aim, shoots and then stands up again.

![The sequence of states in a Trickshot action](../uploads/Main/MecanimStateExcerpt.png)

The sequence of states in a “Trickshot” action

Although this is useful for control purposes, the downside is that the **state machine** will become large and unwieldy as more of these complex actions are added. You can simplify things somewhat just by separating the groups of states visually with empty space in the editor. However, Mecanim goes a step further than this by allowing you to collapse a group of states into a single named item in the state machine diagram. These collapsed groups of states are called **Sub-state machines**.

You can create a sub-state machine by right clicking on an empty space within the **Animator Controller** window and selecting **Create Sub-State Machine** from the context menu. A sub-state machine is represented in the editor by an elongated hexagon to distinguish it from normal states.

![A sub-state machine](../uploads/Main/MecanimSubStateMachine.png)

A sub-state machine

When you double-click the hexagon, the editor is cleared to let you edit the sub-state machine as though it were a completely separate state machine in its own right. The bar at the top of the window shows a “breadcrumb trail” to show which sub-state machine is currently being edited (and note that you can create sub-state machines within other sub-state machines, and so on). Clicking an item in the trail will focus the editor on that particular sub-state machine.

![The breadcrumb trail](../uploads/Main/MecanimStateMachineCrumbs.png)

The “breadcrumb trail”

## External transitions

As noted above, a sub-state machine is just a way of visually collapsing a group of states in the editor, so when you make a transition to a sub-state machine, you have to choose which of its states you want to connect to.

![Choosing a target state within the Trickshot sub-state machine](../uploads/Main/MecanimSelectSubState.png)

Choosing a target state within the “Trickshot” sub-state machine

You will notice an extra state in the sub-state machine whose name begins with *Up*.

![The Up state](../uploads/Main/MecanimSubStateUp.png)

The “Up” state

The *Up* state represents the “outside world”, the state machine that encloses the sub-state machine in the view. If you add a transition from a state in sub-state machine to the *Up* state, you will be prompted to choose one of the states of the enclosing machine to connect to.

![Connecting to a state in the enclosing machine](../uploads/Main/MecanimSubStateExternal.png)

Connecting to a state in the enclosing machine