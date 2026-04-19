# Prevent Unity stripping GPU instancing shaders in the Built-In Render Pipeline

Unity generates Surface **shaders** with instancing [variants](SL-MultipleProgramVariants.md) by default, unless you specify `noinstancing` in the `#pragma` directive. Unity ignores uses of #pragma multi\_compile\_instancing in a **surface shader**.

Unity’s Standard and StandardSpecular shaders have instancing support by default, but with no per-instance properties other than the transform.

If your **scene** contains no **GameObjects** with GPU instancing enabled, then Unity strips instancing shader variants. To override the stripping behaviour:

1. Open Project Settings (menu: **Edit** > **Project Settings**).
2. Go to **Graphics**.
3. In the **Shader Stripping** section, set **Instancing Variants** to **Keep All**.