# Introduction to GameObjects

The **GameObject** is the most important concept in the Unity Editor.

Every object in your game is a **GameObject**, from characters and collectible items to lights, **cameras** and special effects. However, a GameObject can’t do anything on its own; you need to give it properties before it can become a character, an environment, or a special effect.

![Four different types of GameObject: an animated character, a light, a tree, and an audio source](../uploads/Main/GameObjectsExamples.jpg)

Four different types of GameObject: an animated character, a light, a tree, and an audio source

**GameObjects** are the fundamental objects in Unity that represent characters, props and scenery. They do not accomplish much in themselves but they act as containers for **Components**, which implement the functionality.

To give a GameObject the properties it needs to become a light, or a tree, or a camera, you need to add [components](Components.md) to it. Depending on what kind of object you want to create, you add different combinations of components to a GameObject.

Unity has lots of different built-in component types, and you can also make your own components using the [Unity Scripting API](CreatingComponents.md).

For example, a Light object is created by attaching a [Light](class-Light.md) component to a GameObject.

![A simple Light GameObject with several Components](../uploads/Main/GameObjectLightExample1.png)

A simple Light GameObject with several Components

A solid cube object has a **Mesh** Filter and **Mesh Renderer** component, to draw the surface of the cube, and a Box **Collider** component to represent the object’s solid volume in terms of physics.

![A simple Cube GameObject with several Components](../uploads/Main/GameObjectCubeExample1.png)

A simple Cube GameObject with several Components

## Details

A GameObject always has a [Transform](class-Transform.md) component attached (to represent position and orientation) and it is not possible to remove this. The other components that give the object its functionality can be added from the editor’s **Component** menu or from a script. There are also many useful pre-constructed objects (primitive shapes, Cameras, etc) available on the **GameObject > 3D Object** menu, see [Primitive Objects](PrimitiveObjects.md).

Because GameObjects are an important part of Unity, there is a lot of manual content with extensive detail about them. See the following sections for more information on using GameObjects in Unity:

* [Transforms](class-Transform.md)
* [Introduction to components](Components.md)
* [Using Components](UsingComponents.md)
* [Primitive and placeholder objects](PrimitiveObjects.md)
* [Creating components with scripting](CreatingComponents.md)
* [Deactivating GameObjects](DeactivatingGameObjects.md)
* [Tags](Tags.md)
* [Static GameObjects](StaticObjects.md)
* [Saving your work](Saving.md)

You can find out more about controlling GameObjects from **scripts** on the [GameObject scripting reference page](../ScriptReference/GameObject.md).

---

* 2017–08–01 Page amended