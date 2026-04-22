# Access the Particle System from the Animation system

All **particle** properties are accessible by the Animation system, meaning you can **keyframe** them in and control them from your animations.

To access the **Particle System**’s properties, follow these steps:

1. Attach an **Animator component** to the Particle System’s **GameObject**. An Animation Controller and an Animation are also required.
2. To animate a Particle System property, open the **Animation Window** with the GameObject containing the Animator and Particle System selected. Click **Add Property** to add properties.
3. Scroll to the right to reveal the **add controls**.

Note that for curves, you can only keyframe the overall **curve multiplier**, which can be found next to the curve editor in the **Inspector**.