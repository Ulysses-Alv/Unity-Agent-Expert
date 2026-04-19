# Branch group reference

A branch group generates branches and fronds. The available properties change based on your choice between branches and fronds, and on whether the branch group serves as the trunk. Note that if you edit the branch group using its control points, the procedural properties are disabled.

## Using curves

You can use curves to control some properties: click the curve to open the **Curve** window. For more information, refer to [Editing curves](EditingCurves.md).

For many properties, the curve controls the details of the branch’s shape along its length. In that case, the first point of the curve is the base of the branch, and the second point is the tip. For example, you can use the **Crinkliness** property’s curve to create branches that are straight near their base and become more crinkly at the tip, like a corkscrew hazel tree.

The slider next to a curve is a multiplier on the curve values. Use the multiplier to add some variety when reusing the same curves for several branch groups, or to fine-tune a curve’s effect. If the slider is at `0`, the curve has no impact.

## Distribution

| **Property** | **Function** |
| --- | --- |
| **Group Seed** | The randomizing seed for the branch group. Change the seed to vary the group’s procedural generation. |
| **Frequency** | The number of branches (or, for the first group, the number of trunks). Note that the final number of branches might not equal the frequency: a branch group functioning as the trunk or the primary branch group always has the same number of branches as its frequency, but secondary branch groups have a varying number of branches to create a more natural look. |
| **Distribution** | Gives a shape to the distribution of branches around a parent branch. If the group is the trunk, and the frequency is more than 1 (for a reed bed, for example), this property controls the distribution of the trunks around the center point of the group. Use the curve to fine-tune the structure created by the combination of **Group Seed** and **Distribution** and to limit the branch group to portions of the trunk. For example, to prevent branches growing near the ground, move the first point of the curve to `(0.5, 0.0)`. |
| **Twirl** | Move all the branches around the parent branch. Not available when you select **Random** from the **Distribution** dropdown. |
| **Whorled Step** | When you select **Whorled** from the **Distribution** dropdown: how many branches grow out of every growth point along the parent branch. If the branch group has fewer branches than the Whorled Steps value, the branches all grow from the same point on the parent branch. For real plants, this is usually a Fibonacci number. |
| **Growth Scale** | Based on the **Length** property of the branches, scale each branch’s length relative to its position along its parent. For example, have smaller branches higher up the trunk than near the ground. Use the curves to define a scale. The slider is the curve multiplier. |
| **Growth Angle** | Set the branches’ angle of growth relative to the parent. Use the curve to adjust the angles according to the branch’s position along its parent. For example, for the branches growing out of the trunk, the first point is the lower branches; the second point is the higher branches. The slider is the curve multiplier. Note that **Growth Angle** and **Seek Sun** can limit each other. |

## Geometry

| **Property** | **Function** |
| --- | --- |
| **LOD Multiplier** | The tree root node sets the level of detail (LOD) for the whole tree. Use the LOD Multiplier to give the branch group a higher or lower LOD. |
| **Geometry Mode** | Type of geometry for this branch group: Branch Only, Branch + Fronds, Fronds Only. |
| **Branch Material** | Material for the branch bark. |
| **Break Material** | Material for the cross-section of broken branches. |
| **Frond Material** | Material for fronds. |

For more information about the materials, refer to [Design a tree](tree-FirstTree.md).

## Shape

### General

Some options aren’t available for fronds.

| **Property** | **Function** |
| --- | --- |
| **Length** | Range for branch length. Move the range up or down to increase or decrease the size. Widen the range for more variety. Note that changing the length of a parent branch changes the distribution of its children. Also refer to the scaling from the **Growth Scale** property in the **Distribution** section. |
| **Relative Length** | Keep the length and width of branches proportional, so that longer branches are also wider. |
| **Radius** | Adjust the radius of the branches. Use the curve to change the width along the length of the branch: the first curve point is the base of the branch; the second curve point is the tip. Note that a parent group’s radius impacts the radius of its children. The slider is a multiplier for both the curve and the radius inherited from a parent branch group. |
| **Cap Smoothing** | Defines how a branch ends: `1` tapers off to a round edge (like a cactus), and `0` ends on a sudden flat surface (like a sawed-off branch). |
| **Crinkliness** | Define how straight a branch is: `0` is completely straight; `1` is crinkly. Use the curve to control crinkliness along the length of the branch: the first point controls crinkliness near the base of the branch; the second point controls crinkliness near the tip. The slider is a multiplier. |
| **Seek Sun** | Sun seeking branches grow up rather than away from the branch or toward the ground. This value has no impact on a trunk. Use the curve to control sun seeking along the length of the branch: the first point is the base of the branch; the second point is the direction of the tip starting from the middle of the branch. For example, to create branches that touch the ground and then grow upwards, like an old oak tree, put the first point at `(0,0)` and the second point at `(1,1)`. Add points in the middle to give the tip’s growth a more or less pronounced angle. The slider is a multiplier. Note that **Seek Sun** and **Growth Angle** can limit each other. |
| **Noise** | Overall noise factor for the procedural shape generation. A higher value gives the branch an uneven surface and the appearance of damaged bark or uneven growth. If the value is less than `1`, it limits the possible noise from the U and V noise scale values. Use the curve to create a precise noise pattern. The slider is a multiplier. |
| **Noise Scale U** | Scale of the noise around the branch, as a multiplier of **Noise**. |
| **Noise Scale V** | Scale of the noise along the branch, as a multiplier of **Noise**. |
| **Break Chance** | The chance a branch in the group is broken. If the value is `1`, all branches are broken. The broken branches aren’t impacted by **Cap Smoothing**. Note that the cross-section visible on a broken branch has its own material: **Break Material**. |
| **Break Location** | Range of branch breaking points relative to the length of the branches. Move the range up or down to increase or decrease the length of broken branches. Widen the range for more variety. Set a low range and low location for the appearance of burls or burrs. |

### Trunk

For the trunk, you can specify the way the trunk flares out at the bottom:

| **Property** | **Function** |
| --- | --- |
| **Flare Radius** | Adds a flare at the bottom of the tree trunk. Use the flare and a radius curve to create a tree that tapers as it grows. Use no flare and a flat radius curve for tree without much taper. |
| **Flare Height** | The stopping point of the flare. A low flare height limits to the flare to the area of the roots. |
| **Flare Noise** | A higher noise value creates an intermittent flare, like the trunk of a giant sequoia. A low value creates a smooth flare, like a pine tree. |

### Branch groups other than trunk

For branch groups that aren’t the trunk, you can define the way the branch looks where it connects to the trunk or parent branch:

| **Property** | **Function** |
| --- | --- |
| **Weld Length** | How far along the branch the weld stops. The higher the value, the more tapered the weld. This value has no impact if the top and bottom spreads are `0`. |
| **Spread Top** | The thickness of the weld on the top of the branch. |
| **Spread Bottom** | The thickness of the weld on the bottom of the branch. |

## Fronds

This tab is available only if the branch group’s **Geometry** type is **Fronds**.

The settings in this table are relative to the branches the fronds are built around. For example, the overall branch count comes from the **Frequency** property. The **Frond Count** refers to how many fronds each branch has.

| **Property** | **Function** |
| --- | --- |
| **Frond Count** | The number of fronds for each branch. By default, frond are evenly spaced around the branch; use **Frond Crease** to adjust the spacing. |
| **Frond Width** | The width of the fronds. Use the curve to adjust the shape along the length of the branch: the first point controls the start of the frond; the second point controls its tip. The slider is a multiplier. |
| **Frond Range** | Defines the start and end point of the fronds relative to their central point. Note that this allows the fronds to be detached from their branch or the ground. The fronds obey the **Frond Width** curve even when they’re not full length. |
| **Frond Rotation** | Rotate the fronds around their parent branch. When combined with **Frond Crease parameter**, this creates fronds that face only one way, rather than growing at 360 degrees. |
| **Frond Crease** | At `0,` the fronds fan out of the parent branch in all directions. Give a negative or positive value to fold the fronds toward one side of the branch. |

## Wind

Increase or decrease a branch group’s response to the **Wind Zone**. To create a Wind Zone, refer to [Animate trees with Wind Zones](class-WindZone.md).

Note that wind preview is only available in Play mode.

| **Property** | **Function** |
| --- | --- |
| **Main Wind** | Increases the impact of the Wind Zone. |
| **Main Turbulence** | Increases the impact of the Wind Zone’s turbulence on the entire object. Not available for the trunk. |
| **Edge Turbulence** | Increases the impact of the Wind Zone’s turbulence along the edge of the object. Not available for branches. |
| **Create Wind Zone** | Create a [Wind Zone](class-WindZone.md) **GameObject**. Note that the Wind Zone impacts all trees in the **scene**, not just trees created with this model. |

## Additional resources

* [Mesh Renderer component](class-MeshRenderer.md)
* [Animate trees with Wind Zones](class-WindZone.md)
* [Create trees and leaves from meshes](terrain-Tree-From-Mesh.md)