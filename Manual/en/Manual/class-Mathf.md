# Using common math functions

[Switch to Scripting](../ScriptReference/Mathf.md "Go to Mathf page in the Scripting Reference")

Unity’s Mathf class provides a collection of common math functions, including trigonometric, logarithmic, and other functions commonly required in games and app development.

This page provides an overview of the Mathf class and its common uses when scripting with it. For an exhaustive reference of every member of the Mathf class, see the [Mathf script reference](../ScriptReference/Mathf.md).

## Trigonometric

All Unity’s trigonometry functions work in **radians**.

* [`Sin`](../ScriptReference/Mathf.Sin.md)
* [`Cos`](../ScriptReference/Mathf.Cos.md)
* [`Tan`](../ScriptReference/Mathf.Tan.md)
* [`Asin`](../ScriptReference/Mathf.Asin.md)
* [`Acos`](../ScriptReference/Mathf.Acos.md)
* [`Atan`](../ScriptReference/Mathf.Atan.md)
* [`Atan2`](../ScriptReference/Mathf.Atan2.md)

[`PI`](../ScriptReference/Mathf.PI.md) is available as a constant, and you can multiply by the static values [`Rad2Deg`](../ScriptReference/Mathf.Rad2Deg.md) or [`Deg2Rad`](../ScriptReference/Mathf.Deg2Rad.md) to convert between radians and degrees.

## Powers and Square Roots

Unity provides the common power and square root functions you would expect:
- [`Pow`](../ScriptReference/Mathf.Pow.md)
- [`Sqrt`](../ScriptReference/Mathf.Sqrt.md)
- [`Exp`](../ScriptReference/Mathf.Exp.md)

As well as some useful power-of-two related functions. These are useful when working with common binary data sizes, which are often constrained or optimized to power-of-two values (such as texture dimensions):

* [`ClosestPowerOfTwo`](../ScriptReference/Mathf.ClosestPowerOfTwo.md)
* [`NextPowerOfTwo`](../ScriptReference/Mathf.NextPowerOfTwo.md)
* [`IsPowerOfTwo`](../ScriptReference/Mathf.IsPowerOfTwo.md)

## Interpolation

Unity’s interpolation functions allows you to calculate a value that is some way between two given points. Each of these functions behaves in a different way, suitable for different situations. See the examples in each for more information:

* [`Lerp`](../ScriptReference/Mathf.Lerp.md)
* [`LerpAngle`](../ScriptReference/Mathf.LerpAngle.md)
* [`LerpUnclamped`](../ScriptReference/Mathf.LerpUnclamped.md)
* [`InverseLerp`](../ScriptReference/Mathf.InverseLerp.md)
* [`MoveTowards`](../ScriptReference/Mathf.MoveTowards.md)
* [`MoveTowardsAngle`](../ScriptReference/Mathf.MoveTowardsAngle.md)
* [`SmoothDamp`](../ScriptReference/Mathf.SmoothDamp.md)
* [`SmoothDampAngle`](../ScriptReference/Mathf.SmoothDampAngle.md)
* [`SmoothStep`](../ScriptReference/Mathf.SmoothStep.md)

Note that the [Vector classes](scripting-vectors.md) and the [`Quaternion`](class-Quaternion.md) class all have their own interpolation functions (such as Quaternion.Lerp) which allow you to interpolate positions, directions and rotations in multiple dimensions.

## Limiting and repeating values

These simple helper functions are often useful in games or apps and can save you time when you need to limit values to a certain range or repeat them within a certain range.

* [`Max`](../ScriptReference/Mathf.Max.md) and [`Min`](../ScriptReference/Mathf.Min.md)
* [`Repeat`](../ScriptReference/Mathf.Repeat.md) and [`PingPong`](../ScriptReference/Mathf.PingPong.md)
* [`Clamp`](../ScriptReference/Mathf.Clamp.md) and [`Clamp01`](../ScriptReference/Mathf.Clamp01.md)
* [`Ceil`](../ScriptReference/Mathf.Ceil.md) and [`Floor`](../ScriptReference/Mathf.Floor.md)

## Logarithmic

The [`Log`](../ScriptReference/Mathf.Log.md) function allows you to calculate the logarithm of a specified number, either the natural logarithm or in a specified base. Additionally the [`Log10`](../ScriptReference/Mathf.Log10.md) function returns the base–10 logarithm of the specified number.

## Other functions

For the full list of functions in the Mathf class, see the [Mathf script reference](../ScriptReference/Mathf.md).

Mathf