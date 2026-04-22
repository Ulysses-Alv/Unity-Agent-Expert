# Use Unity Accelerator on the command line

Unity **Accelerator** has a large set of command line tools you can use to troubleshoot, configure, and make use of one or more cache servers. For ease of use, it’s best practice to add the location of the `unity-accelerator` executable to your `PATH`. The executable was installed in the path you chose when running the installer, for example `C:\Program Files\Unity\accelerator`.

The CLI tool has an extensive help system. To get started, you can run `unity-accelerator` with no arguments:

```
$ unity-accelerator
Unity Accelerator v1.0.524+g96c5e18

Run on a local network to speed up transactions with Unity Services.

.....
```

You can also run `unity-accelerator --all-help` to output all the help text for all commands in one go.

When working on a specific cache server, to reconfigure it for example, it’s best to change your working directory to the same as where the `unity-accelerator.cfg` file is. You can also set the `UNITY_ACCELERATOR_PERSIST` environment variable to that directory, or you can try to remember to always use the `--persist <dir>` option.

One rather common set of cli tools used are the dashboard tools:

```
$ unity-accelerator dashboard password newaccount
Password: ****
   Again: ****
$ unity-accelerator dashboard list
admin
newaccount
$ unity-accelerator dashboard url
http://172.18.37.249:8080/dashboard/
```

That final tool, `dashboard url`, is useful to find out what URL to put in your browser to visit the built-in dashboard.

## Additional resources

* [Command line arguments reference](EditorCommandLineArguments.md#accelerator)