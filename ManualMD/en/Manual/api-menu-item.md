# Invoke from the main menu

To create a custom menu item that invokes a search, use the `MenuItem` attribute. The following example creates a menu item that searches the **Asset Store** for free assets:

```
[MenuItem("Search/Free Asset Gifts")]
static void SearchStore()
{
    var storeContext = SearchService.CreateContext("store", "price=0");
    var viewState = new SearchViewState(storeContext,
        UnityEngine.Search.SearchViewFlags.DisableNoResultTips |
        UnityEngine.Search.SearchViewFlags.GridView);
    viewState.windowTitle = new GUIContent("Free Stuff for Christmas");
    viewState.queryBuilderEnabled = true;
    SearchService.ShowWindow(viewState);
}
```