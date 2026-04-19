# Manage USS variables in UI Builder

[USS variables](UIE-USS-variables.md) define values that you can reuse in other USS rules. USS variables are primarily used for themes, with the default Unity themes exposing a long list of standard variables to make your UI more consistent with standard controls.

## Create a USS variable in UI Builder

You can create USS variables in UI Builder. You can also use a text editor to create variables [directly in a USS file](UIE-USS-CustomProperties.md).

To create a USS variable in UI Builder:

1. In the **StyleSheets** panel, do the following:

   * If there is no selector, select **Add new selector** to create a new selector.
   * If you want to create a variable that’s available to all selectors, select the **.root** selector.
   * If you want to create a variable that’s available only to a specific selector, select that selector.
2. In the **Variables** section of the ****Inspector**** panel, select **+**. A dropdown list displays the available variable types, such as **Color**, **Number**, **Dimension**, **String**, **Enum**, and **Resource**.
3. Select the variable type. The type determines the type of value you can assign to the variable. A default name and value are set based on the variable type.
4. To rename the variable, in the first column field, enter the name. The name can contain only letters, numbers, and hyphens.
5. To set the value of the variable, in the second column field, enter the value. You can use any valid USS value, including other variables.

   ![USSVariablesCreate](../uploads/Main/UIBuilder/USSVariablesCreate.png)

   USSVariablesCreate
6. To reorder the variables, drag the variable to a new position. The order of the variables determines the order in which they’re applied. If you have two variables with the same name, the variable that’s defined last is used.
7. To delete a variable, select the variable and then select **-**. You can delete multiple variables at once.
8. To save the variable, in the **Viewport** panel, select **File** > **Save**. This saves the variables to the USS file.

**Tip**: When editing a USS variable, you can click outside the field to finish the edit, or press **Enter**. To cancel the edit, press **Esc**.

## Assign a USS variable to a style property

You can assign a USS variable that’s in the current theme or any assigned StyleSheets to the current UI Document (UXML) to a style property in UI Builder. When you work on an Editor UI, make sure to [enable **Editor Extension Authoring**](UIB-interface-overview.md#enable-editor-extension-authoring-for-ui-documents-uxml) to see all available Editor variables.

To assign a variable:

1. In the **StyleSheet** panel, select the selector.
2. In the **Inspector** panel, right-click a style field.
3. Select **Set Variable**. This converts the style field into a text field.
4. Enter the name of the USS variable. As you type the name of the variable, a dropdown list displays the available variables.

   ![USSVariablesSearch](../uploads/Main/UIBuilder/USSVariablesSearch.png)

   USSVariablesSearch

   You can also select a variable from the dropdown list, and check its current value and the StyleSheet asset it’s coming from. Pressing **Enter** sets the variable and reverts the style field back to its original type.

   **Tip**: For style fields that are already text fields, you can also assign a variable by starting to type `--` instead of a number.

You can identify which style fields have a variable assigned by checking if the field’s label has a chain icon on the left.

![USSVariablesSet](../uploads/Main/UIBuilder/USSVariablesSet.png)

USSVariablesSet

To remove a variable assignment, right-click a style field and select **Remove Variable**.

**Note**: Any inline styles set on a style property will override the USS variables.

## Additional resources

* [UI Builder](UIBuilder.md)
* [USS variables](UIE-USS-variables.md)