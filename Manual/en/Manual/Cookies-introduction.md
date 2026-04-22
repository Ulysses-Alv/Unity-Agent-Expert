# Introduction to cookies

![An example of baked fake caustics achieved using baked light cookies.](../uploads/Main/updated-cookie.png)

An example of baked fake caustics achieved using baked light cookies.

A cookie is a mask that you place on a light to create a shadow with a specific shape or color, which changes the appearance and intensity of the light.

Cookies are an efficient way of simulating complex lighting effects with minimal runtime performance impact. Effects you can simulate with cookies include caustics, soft shadows, and light shapes.

For example, the following illustration shows how you can simulate a real-time point light shadow with a light cookie. If there are no dynamic objects that can enter the area with a light cookie shadow, a user is unlikely to notice a difference between the cookie and a real-time shadow.

![Point light with a light cookie emulating shadows (left), and a point light without shadows (right)](../uploads/urp/shadows/shadows-with-light-cookie.png)  
Point light with a light cookie emulating shadows (left), and a point light without shadows (right)

Spot lights with cookies can create the effect of light coming in from a window.

If you create a Texture that contains an alpha channel and assign it to the **Cookie** variable of the light, the cookie is projected from the light. The cookie’s alpha mask modulates the light’s brightness, creating light and dark spots on surfaces. This is a great way to add complexity or atmosphere to a **scene**.

## Render pipeline compatibility

See [render pipeline feature comparison](render-pipelines-feature-comparison.md) for more information about support for cookies across **render pipelines**.

## Additional resources

* [Unity forum: New Feature - Baked Light Cookies](https://forum.unity.com/threads/2020-1-new-feature-baked-light-cookies.848269/)