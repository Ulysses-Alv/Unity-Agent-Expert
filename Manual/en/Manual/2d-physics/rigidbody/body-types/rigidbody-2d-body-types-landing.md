# Rigidbody 2D body types

There are three options for **Body Type** which define the behavior of the **Rigidbody** 2D. Any **Collider** 2D attached to that Rigidbody 2D inherits the Rigidbody 2D’s Body Type as well.

The selected Body Type defines the Rigidbody 2D’s movement behavior (position and rotation) and Collider interactions. When a Body Type changes, Unity recalculates various mass-related internal properties, and all existing contacts for the Collider 2Ds attached to the Rigidbody 2D need to be re-evaluated during the **GameObject**’s next FixedUpdate. Depending on how many contacts and Collider 2Ds are attached to the body, changing the Body Type can cause variations in performance.

The properties of the Rigidbody 2D component in its **Inspector** window differs depending on which Body Type you have selected. You can refer to the following pages for detailed information about the property settings for each Body Type:

| **Page** | **Description** |
| --- | --- |
| [Introduction to Rigidbody 2D body types](introduction-to-rigidbody-2d-body-types.md) | Understand the concepts behind the different types of Rigidbody 2D body types. |
| [Dynamic Body Type](dynamic/dynamic-body-type-landing.md) | Techniques and resources for working with a Dynamic Body Type Rigidbody 2D. |
| [Kinematic Body Type](kinematic/kinematic-body-type-landing.md) | Techniques and resources for working with a Kinematic Body Type Rigidbody 2D. |
| [Static Body Type](static/static-body-type-landing.md) | Techniques and resources for working with a Static Body Type Rigidbody 2D. |

## Additional resources

* [API documentation on Rigidbody2D](../../../../ScriptReference/Rigidbody2D.md)