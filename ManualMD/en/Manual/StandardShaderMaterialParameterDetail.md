# Secondary Maps (Detail Maps) and Detail Mask

Secondary Maps (or Detail maps) allow you to overlay a second set of textures on top of the main textures listed above. You can apply a second Albedo colour map, and a second **Normal map**. Typically, these would be mapped on a much smaller scale repeated many times across the object’s surface, compared with the main Albedo and Detail maps.

The reason for this is to allow the material to have sharp detail when viewed up close, while also having a normal **level of detail** when viewed from further away, without having to use a single extremely high texture map to achieve both goals.

Typical uses for detail textures would be:

* Adding skin detail, such as pores and hairs, to a character’s skin
* Adding tiny cracks and lichen growth to a brick wall
* adding small scratches and scuffs to a large metal container

![This character has a skin texture map, but no detail texture yet. We will add skin pores as a detail texture.](../uploads/Main/StandardShaderDetailNotAppliedYet.jpg)

This character has a skin texture map, but no detail texture yet. We will add skin pores as a detail texture.

![The Albedo skin pore detail texture](../uploads/Main/StandardShaderDetailSkin.jpg)

The Albedo skin pore detail texture

![The normal map for the skin pore detail](../uploads/Main/StandardShaderDetailSkinNormal.jpg)

The normal map for the skin pore detail

![The end result, the character now has subtle skin pore detail across her skin, at a much higher resolution than the base Albedo or Normal map layer would have allowed.](../uploads/Main/StandardShaderDetailApplied.jpg)

The end result, the character now has subtle skin pore detail across her skin, at a much higher resolution than the base Albedo or Normal map layer would have allowed.

![Detail textures can have a subtle but striking effect on the way light hits a surface. This is the same character in a different lighting context.](../uploads/Main/StandardShaderDetailDifferentContext.jpg)

Detail textures can have a subtle but striking effect on the way light hits a surface. This is the same character in a different lighting context.

If you use a single normal map do ALWAYS plug it into the primary channel. The Secondary normal map channel is more expensive than the primary one but has the exact same effect.

### Detail Mask

The detail mask texture allows you to control where the detail texture is applied on your model. This means you can decide which areas should display the detail texture and which should not. For instance, if your model has skin pores, you might mask the lips and eyebrows to prevent the pores from appearing in those areas.