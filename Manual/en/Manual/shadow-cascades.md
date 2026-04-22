# Introduction to shadow cascades

Shadow cascades let you improve the visual fidelity of shadows without increasing the shadow map resolution.

For example, the following illustration shows shadow rendering with different numbers of shadow cascades. The shadow resolution is 2048 in all cases.

![Shadow rendering with different cascade numbers. A: 1 cascade, B: 2 cascades, C: 3 cascades, D: 4 cascades](../uploads/urp/shadows/shadow-cascades-comparison.png)  
Shadow rendering with different cascade numbers. A: 1 (no cascades), B: 2 cascades, C: 3 cascades, D: 4 cascades.

Shadow cascades only work with directional lights.

From a technical perspective, shadow cascades mitigate a problem called [perspective aliasing](shadow-cascades-implementation-details.md#perspective-aliasing), where real-time shadows from [directional lights](Lighting.md) appear pixelated when they are near the **camera**.

## Additional resources

* [Configure shadow cascades](shadow-cascades-use.md)
* [Performance impact of shadow cascades](shadow-cascades-performance.md)
* [Technical implementation details](shadow-cascades-implementation-details.md)