# Per-frame updates

Unity performs some updates once per frame. Your MonoBehaviour **scripts** can hook into this update loop through the [`MonoBehaviour.Update`](../ScriptReference/MonoBehaviour.Update.md) method. For example, in your game character’s `Update` method, you might read the user input from a joypad, and move the character forward a certain amount. It’s important to remember when handling time-based actions like this that the game’s frame rate can vary and so the length of time between `Update` calls also varies.

The variable frame rate of a game is often expressed in **frames per second**, or **FPS**. Frame rate varies according to factors like the capabilities of the host device and the complexity of the graphics and computation required to draw each frame. For example, your game may run at a slower frame rate when there are 100 active, on-screen characters than when there is only one.

Unless otherwise constrained by your [quality settings](class-QualitySettings.md) or by [Adaptive Performance](adaptive-performance/adaptive-performance.md), Unity tries to run your game at the fastest frame rate possible. You can see more details of what occurs each frame in the **Game Logic** section of the [execution order diagram](execution-order.md).

Consider the task of moving an object forward gradually, one frame at a time. It might seem at first that you could just translate the object by a fixed distance each frame:

```
//C# script example
using UnityEngine;
using System.Collections;

public class ExampleScript : MonoBehaviour {
    public float distancePerFrame;
    
    void Update() {
        transform.Translate(0, 0, distancePerFrame); // this is incorrect
    }
}
```

However with this code, as the frame rate varies, the object’s apparent speed also varies. If the game is running at 100 FPS, the object moves `distancePerFrame` 100 times per second. But if the frame rate slows to 60 FPS (due to CPU load, for example) then it only steps forward 60 times per second and therefore covers a shorter distance over the same amount of time.

In most cases this is undesirable, particularly with games and animation. It’s much more common to want your in-game objects to move at steady and predictable speeds regardless of the frame rate. The solution is to scale the amount of movement each frame by the amount of time elapsed each frame, which you can read from the [Time.deltaTime](../ScriptReference/Time-deltaTime.md) property:

```
//C# script example
using UnityEngine;
using System.Collections;

public class ExampleScript : MonoBehaviour {
    public float distancePerSecond;
    
    void Update() {
        transform.Translate(0, 0, distancePerSecond * Time.deltaTime);
    }
}
```

Note that the movement is now given as `distancePerSecond` rather than `distancePerFrame`. As the frame rate varies, the size of the movement step will vary accordingly and so the object’s speed will be constant.

Depending on your target platform, use either [Application.targetFrameRate](../ScriptReference/Application-targetFrameRate.md) or [QualitySettings.vSyncCount](../ScriptReference/QualitySettings-vSyncCount.md) to set the frame rate of your application. For more information, see the [Application.targetFrameRate API documentation](../ScriptReference/Application-targetFrameRate.md).

## Additional resources

* [Fixed updates](fixed-updates.md)
* [Handling variations in time](time-handling-variations.md)
* [In-game versus real time](time-scale.md)
* [Capture frame rate](time-capture-frame-rate.md)