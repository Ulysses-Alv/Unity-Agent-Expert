# HDR Output debug views in URP

URP offers three debug views for **HDR** rendering.

These views are:

* [Gamut View](#gamut-view)
* [Gamut Clip](#gamut-clip)
* [Values exceeding Paper White](#values-exceeding-paper-white)

To access them, navigate to **Window** > **Analysis** > **Render Pipeline Debugger** > **Lighting** > **HDR Debug Mode**.

## Gamut View

![Gamut Debug View](../../../uploads/urp/post-proc/hdr/HDR-Output-GamutView.png)

Gamut Debug View

The triangles in this debug view indicate which parts of three specific color gamuts this **scene** covers. The small triangle displays the [Rec709](https://en.wikipedia.org/wiki/Rec._709) gamut values, the medium triangle displays the [P3-D65](https://en.wikipedia.org/wiki/DCI-P3) gamut values, and the large triangle displays the [Rec2020](https://en.wikipedia.org/wiki/Rec._2020) gamut values. This enables you to check color plot changes while color grading. It can also help you ensure that you benefit from the wider color gamut available in HDR.

## Gamut Clip

![Gamut Clip Debug View](../../../uploads/urp/post-proc/hdr/HDR-Output-GamutClip.png)

Gamut Clip Debug View

This debug view indicates the relationship between scene values and specific color gamuts. Areas of the screen with values within the Rec709 gamut are green, areas outside of the Rec709 gamut but inside the P3-D65 gamut are blue, and areas outside of both are red.

## Values exceeding Paper White

![Values Exceeding Paper White Debug View](../../../uploads/urp/post-proc/hdr/HDR-Output-OverPaperWhite.png)

Values Exceeding Paper White Debug View

This debug view uses a color coded gradient to indicate parts of the Scene that exceed the Paper White value and Max Nits. The gradient ranges from yellow to red and blue. Yellow corresponds to **Paper White** +1, red corresponds to **Max Nits**, and blue corresponds to **Max Nits**+1.