# Mirror multiple Unity Accelerator instances

You can configure multiple Unity **Accelerator** instances to mirror the same cached data. You might want to have two mirrored cache servers running on the same network for high availability. For example, if one crashes, has a hardware failure, or needs maintenance for whatever reason, the other server can compensate. You might also mirror cache servers that are remote from each other, for example a main office might mirror its data to a branch office, providing copies of its cached data for use at the branch office as well.

The mirroring features only provide outgoing data capabilities, also known as push, or forwarding. You can push or forward to multiple cache servers, or even chain them. For example, A sending to B, with B sending to C, and even C sending back to A. You don’t need to worry about circular chains because items are time stamped so a receiving cache server ignores any incoming items that aren’t newer than what it already has.

## Mirror example

In the following example, server A has an `ip:port` of `1.1.1.1:1111` and server B is `2.2.2.2:2222`. Both have no data to start with, so the following puts an item into A and notes that it’s not in B:

```
$ echo test1 | unity-accelerator cache put 1.1.1.1:1111 namespace1 key1
$ unity-accelerator cache get 1.1.1.1:1111 namespace1 key1
test1
$ unity-accelerator cache get 2.2.2.2:2222 namespace1 key1
key not found
```

The following configures A to forward to B. It alters A’s configuration, so log into the machine running accelerator A, shut down the accelerator and perform the following:

```
# Shutdown Accelerator A
$ cd <directory where unity-accelerator.cfg is>
$ unity-accelerator config set ProtobufForwardPuts protobuf:2.2.2.2:2222
# Start Accelerator A back up
```

You can now put another item into A’s cache and see it forwarded to B:

```
$ echo test2 | unity-accelerator cache put 1.1.1.1:1111 namespace2 key2
$ unity-accelerator cache get 1.1.1.1:1111 namespace2 key2
test2
$ unity-accelerator cache get 2.2.2.2:2222 namespace2 key2
key not found
# It will take up to a minute before the forwarding occurs, due to how the algorithm works with batching, etc. But within a minute or two you should see...
$ unity-accelerator cache get 2.2.2.2:2222 namespace2 key2
test2
```

However, this only covers new data: any new items put into accelerator A will be forwarded to accelerator B. It doesn’t cover any existing data. In this example, the first item still is not in accelerator B.

```
$ unity-accelerator cache get 2.2.2.2:2222 namespace1 key1
key not found
```

To cover this case, and in cases where you just want to forward all data one time (which is useful when switching from one accelerator to a new one) you can use the “push” command line tool:

```
$ unity-accelerator cache push 1.1.1.1:1111 2.2.2.2:2222
{"msg":"push completed", "destination":"2.2.2.2:2222", "namespaces":2, "keys":2, "bytes":12, "supercededs":1, "errors":0, "elapsed":"44.054078ms"}
$ unity-accelerator cache get 2.2.2.2:2222 namespace1 key1
test1
```

The push tool sends a request to the first accelerator and tells it to send all its data to the second accelerator. When it is done, the first accelerator sends back the “push completed” response with detail as to what was done.

There are additional, advanced configuration options you can research by running `unity-accelerator config help | less` and searching for Forward. You can specify the forwarding concurrency, queue size, error handling, etc.

Also note that you can alter the configuration values for the MaintenanceStart and MaintenanceDuration of each mirrored accelerator so that they don’t all shutdown for maintenance at the same time.

## Additional resources

* [Monitor Unity Accelerator](accelerator-monitor.md)
* [Use Unity Accelerator on the command line](accelerator-command-line.md)
* [Command line arguments reference](EditorCommandLineArguments.md#accelerator)