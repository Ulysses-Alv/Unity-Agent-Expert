# Work with blend shapes

## Prepare the artwork

Once you have set up blend shapes in your 3D modeling application, do the following:

1. In your 3D modeling application, enable these export settings:
   * Enable exporting animation.
   * Enable exporting blend shapes for deformed models.
2. Export your selection to an .fbx file.
3. [Import your FBX file](ImportingModelFiles.md) into Unity.
4. Select the newly imported Model in the Hierarchy window. The **Inspector** window displays the **BlendShapes** section containing all the blend shapes under the [SkinnedMeshRenderer](class-SkinnedMeshRenderer.md) component.
5. For each listed blend shape, you can change its influence (weighting) to the default shape, where:
6. `0` means the blend shape has no influence.
7. `100` means the blend shape has full influence.

## Create animations In Unity

To create a blend animation:

1. Open the Animation window (from the main Unity menu: **Window** > **Animation** > **Animation**).
2. Select the **Animation Clip** where you want to add the blend shape animation. You can select **Create New Clip** to create a new Animation Clip.
3. Click **Add Property** and select **[Your Character]** > **[Your Character Body]** > **Skinned Mesh Renderer** and select the blend shape you want to animate.
4. Select the added blend shape property, adjusting the **keyframes** and blend weights, and set keyframes.

To preview your animation, click **Play** in the Editor window or the Animation window.

## Scripting access

You can also set the blend weights through scripting using functions like [GetBlendShapeWeight](../ScriptReference/SkinnedMeshRenderer.GetBlendShapeWeight.md) and [SetBlendShapeWeight](../ScriptReference/SkinnedMeshRenderer.SetBlendShapeWeight.md).

To check how many blend shapes a **Mesh** has, use the [blendShapeCount](../ScriptReference/Mesh-blendShapeCount.md) variable.

This code example demonstrates how to blend a default shape into two other blend shapes over time when attached to a **GameObject** with three or more blend shapes:

```
using UnityEngine;
using System.Collections;

public class BlendShapeExample : MonoBehaviour
{
    int blendShapeCount;
    SkinnedMeshRenderer skinnedMeshRenderer;
    Mesh skinnedMesh;
    float blendOne = 0f;
    float blendTwo = 0f;
    float blendSpeed = 1f;
    bool blendOneFinished = false;

    void Awake ()
    {
        skinnedMeshRenderer = GetComponent<SkinnedMeshRenderer> ();
        skinnedMesh = GetComponent<SkinnedMeshRenderer> ().sharedMesh;
    }

    void Start ()
    {
        blendShapeCount = skinnedMesh.blendShapeCount;
    }

    void Update ()
    {
        if (blendShapeCount > 2) {
            if (blendOne < 100f) {
                skinnedMeshRenderer.SetBlendShapeWeight (0, blendOne);
                blendOne += blendSpeed;
            } else {
                blendOneFinished = true;
            }

            if (blendOneFinished == true && blendTwo < 100f) {
                skinnedMeshRenderer.SetBlendShapeWeight (1, blendTwo);
                blendTwo += blendSpeed;
            }
        }
    }
}
```