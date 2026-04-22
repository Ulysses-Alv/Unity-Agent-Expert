# Search for files

The Files search provider runs a file system search that supports regex and globs to match a specific pattern, such as file type, file name, or folder name.

Unlike the Asset (project) search provider, the Files provider doesn’t require an index, so you can use it even if your project has no indices.

**Tip:** You can use the Files provider in the general **Search** window, or in a specialised window. To open a specialized search window for the Files provider, from the main menu, select **Window** > **Search** > **Find Files**.

## Query syntax

**Provider token:** `find:`

## Use the Files provider with other providers

To implicitly include `find` in searches by other providers, use the **Show more results** toggle in the **View** menu (![The View menu](../uploads/Main/view_menu.png)). This option adds results from an implicit `find:` search to the following tabs:

* The **Files** tab, which is only available when **Show more results** is active.
* The **All**, **Project**, and **Resources** tabs.

If **Show more results** is inactive, you have to use the `find:` token to search for files.

### Query examples

| Filter type | Query example | Description |
| --- | --- | --- |
| **Word search** | `find: Paint Mat` | Searches for all assets paths containing the words `paint` AND the word `mat` (for example, `PaintBrush_Mat.mat`, `DryWallPainted_Mat.mat`). |
| **Glob** | `find: AnimationScripts/*.cs` | Finds all C# files under `Assets/AnimationScripts`. |
| **Regex** | `find: (.*\/)+.+\s.+` | Find all paths with a space in their name. |

### Glob wildcards

You can use a glob queries and [C# Regex](https://docs.microsoft.com/en-us/dotnet/standard/base-types/regular-expression-language-quick-reference) to match patterns.

| **Global wildcards** | **Description** | **Example** | **Matches** | **Doesn’t match** | **Equivalent regex** |
| --- | --- | --- | --- | --- | --- |
| `*` | Matches any number of any characters including none | `Law*` | `Law`, `Laws`, `Lawyer` | `Groklaw`, `La`, `aw` | `.*` |
| `.` | Matches any single character including none | `Law.` | `Law`, `Laws`, `Lawyer` | `Groklaw`, `La`, `aw` | `.` |

### Fuzzy logic query

To enable fuzzy logic, use the **Show more results** toggle in the **View** menu (![The View menu](../uploads/Main/view_menu.png)).

### Provider filters

The Files provider has additional filters in the [visual query builder](search-visual-query.md). This table lists the filters and their textual query equivalents.

| Filter | Search token | Query example | Description |
| --- | --- | --- | --- |
| **Files with Spaces** | `find: \s+` | `find: \s+ material` | Find all files with spaces in their names or file paths, and the word `material` in the name or path. |
| **Numeric Files** | `find: \d+\.\w+$` | `find: \d+\.\w+$` | Find all files with numbers in their names. |
| **Search Packages** | `find: +packages` | `find: +packages png` | Find all `png` files and all files with the string `png` in their names or path, including those in packages. If you don’t use this token, package files are excluded from the search. |

## Results

**Search** window tab: **Project**.

## Actions

The context menu for the Files search provider includes the following actions:

| Action | Description |
| --- | --- |
| **Select** | Open the file in the ****Inspector**** and **Project** windows. This is the default double-click action. To change the default action, refer to [Preferences](Preferences.md#search). |
| **Open** | Open the file in the Unity Editor or an external editor. |
| **Reimport** | For imported files, reimports the file from the source file. |
| **Reveal** | Opens the file in your operating system’s file browser. |
| **Delete** | Deletes the file. The Unity Editor requests confirmation before deleting. |
| **Copy Path** | Copies the file’s path relative to your project’s root folder to the clipboard. |
| **Copy GUID** | Copies the file’s GUID to the clipboard. |
| **Properties** | Opens a **Focused inspector** window for the file. |

## Additional resources

* [Visual query builder](search-visual-query.md)
* [Textual query references](search-textual-query.md)
* [Activate and deactivate search providers](search-manage-providers.md)
* [Search query operators](search-query-operators.md)