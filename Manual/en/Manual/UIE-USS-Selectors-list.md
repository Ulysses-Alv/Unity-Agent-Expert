# Selector lists

A selector list is a comma-separated list of selectors that share the same style rule.

## Syntax

A selector list consists of multiple selectors separated by commas:

```
selector1, selector2 {...}
```

## Example

The following two USS snippets have the same effect.

```
#container2 {
  background-color: pink;
  border-radius: 10px;
}

Button {
  background-color: pink;
  border-radius: 10px;
}
```

```
#container2, Button {
   background-color: pink;
     border-radius: 10px;
 }
```

## Additional resources

* [Best practices for USS](UIE-USS-WritingStyleSheets.md)