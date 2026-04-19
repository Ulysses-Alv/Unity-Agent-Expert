# Change how lights fade to match the Built-In Render Pipeline

After converting the necessary **project settings** and materials from the Built-In **Render Pipeline** to URP, the overall look of the **scene** might still not match the look of the original project.

One reason for this is that the light falloff functions in the Built-In Render Pipeline and URP are different. URP implements inverse square light falloff, while the Built-In Render Pipeline uses quadratic falloff. Changes in quality or light component settings might not be enough to achieve the same look in URP.

To change the light falloff function in URP and make it look similar to the Built-In Render Pipeline falloff, refer to [Change the light falloff function in URP](../lighting/custom-lighting-change-light-falloff.md).

## Additional resources

* [Custom lighting in URP](../lighting/custom-lighting-landing.md)