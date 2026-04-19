# Check how Unity batches lights

To open and use the debugger, follow these steps:

1. Open the debugger window by going to **Window** > **2D** > **Light Batching Debugger**.  
   ![Opening the Light Batching Debugger window](../../uploads/urp/2D/light-batching-debugger-0.png)
2. View the Light Batching Debugger updates in real time by keeping the Game view and the debugger window open at the same time.  
   ![Light Batching Debugger window without a selected batch.](../../uploads/urp/2D/light-batching-debugger-1.png)
3. Select a batch from the left side of the debugger window to view Lights and Shadow Casters in the current batch.  
   ![Light Batching Debugger window with selected batch.](../../uploads/urp/2D/light-batching-debugger-2.png)  
   Sorting Layers that are color coded differently means that they’re in different batches with different **Batch IDs** and aren’t batched together.  
   ![Differently color coded](../../uploads/urp/2D/light-batching-debugger-color-1.png)  
   Sorting Layers that share the same color code are batched together and will share the same **Batch ID**.  
   ![Similarly color coded](../../uploads/urp/2D/light-batching-debugger-color-2.png)
4. Select adjacent batches to compare the differences between the selected batches. The debugger window displays the Light(s) and Shadow Caster(s) included in each batch in separate panels.  
   ![Light Batching Debugger window](../../uploads/urp/2D/light-batching-debugger-3.png)  
   In this example, **Light A** exists in **Batch 0** and not in **Batch 1**. The debugger provides instructions at the bottom of the window on what you need to do to have Unity batch the two selected batches together; that is, **Batch 0** contains **Light A** which currently only targets the **BG** Sorting Layer. By having **Light A** also target the **Default** Sorting Layer, Unity may be able to batch both **Batch 0** and **Batch 1** together.