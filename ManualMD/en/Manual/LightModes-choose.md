# Choose a Light Mode

You can choose whether lighting is pre-calculated (baked), updated in real-time, or a combination (mixed), depending on the performance and visual quality you aim for:

* Baked: Provides essential illumination without significant performance impact.
* Mixed: Ensures a stable lighting setup and preserves shadows.
* Realtime: Adds dynamic effects and realism, often in more advanced or interactive environments.

The following table compares **light modes**:

| **Light Mode** | **Baked** | **Mixed** | **Realtime** |
| --- | --- | --- | --- |
| Typical use case | Background and **ambient lights** | Fixed lights | Moving, destructible, or flickering lights |
| Performance impact at runtime | None | High | Highest |
| Iteration time | Slowest, due to Unity baking both direct and indirect lighting | Slow, due to Unity baking indirect lighting | Fast |
| Direct lighting | Baked | Real-time | Real-time |
| Indirect lighting | Baked | Baked | No, unless you use **Enlighten** Realtime **Global Illumination**. |
| Specular highlights | No | Real-time | Real-time |
| Shadows from dynamic **GameObjects** | Baked | Real-time | Real-time |
| Shadows from static GameObjects, up to Shadow Distance | Baked | Real-time, or baked into shadow mask textures | Real-time |
| Casts shadows from static GameObjects, beyond Shadow Distance | Baked | No shadows, or baked into shadow maps if you use the optional [Shadowmask Lighting Mode](lighting-mode.md). | No shadow |
| Contributes to Enlighten Realtime Global Illumination | Yes | Yes | Yes |

## Additional resources

* [Enable shadows](shadow-configuration.md)