# Introduction to lens flare effects

Understand how Unity manages **lens flares**, which simulate the effect of lights refracting inside a **camera** lens. Use them to represent bright lights or to add a bit more atmosphere to your **Scene**.

## Flares

A Flare asset allows you to create and configure the appearance of lens flares.

Flares work by containing several Flare **Elements** on a single **Texture**. Within the Flare, you pick and choose which **Elements** you want to include from any of the Textures.

## Displaying flares

A Lens Flare component displays a lens flare that is configured by a [Flare asset](class-Flare.md).

You can display a Flare asset with a [Light component](class-Light.md). If you do this, Unity automatically tracks the position and direction of the Light and uses those values to configure the appearance of the lens flare.

Use this component instead to configure the values of the lens flare yourself, which gives you more precise control.