# Verify the Unity Accelerator version

You can verify the Unity **Accelerator** version from a client machine accessing a remote Unity Accelerator instance and also from the server machine.

## Verify on a client machine

From the client, access the API agent health end point by visiting the following page from a browser: `http://[AcceleratorIP]:[AcceleratorPort]/api/agent-health`

Alternatively, examine the unity-accelerator.log file in the log folder of the Accelerators configuration directory. The sysinfo message displays the currently loaded Unity Accelerator version:

```
{"level":"info","ts":"2025-03-05T11:10:52.144Z","msg":"sysinfo","agent_id":"john.smith_id","agent_name":"john.smith","component":"agent","pid":91985,"agentversion":"v1.0.941+g6b39b61","goversion":"go1.15.4","os":"darwin","arch":"amd64","compiler":"gc","maxprocs":10,"numcpu":10}
```

The data is arranged in key-value pairs, and the `agentversion` key displays the current version `"agentversion":"v1.0.941+g6b39b61"`

## Verify on a server

From the Accelerator server you can use the Unity Accelerator application to report the version with `unity-accelerator --version`, or `unity-accelerator.exe --version` on Windows

On Windows, the Unity Accelerator application is usually installed under `C:\Program Files\Unity\accelerator`, and on macOS it is usually installed under `/Applications/Unity/accelerator` unless a different location was specified during install.

## Verify on Linux systems

For Unity Accelerator installations on Linux systems, all Linux executables have signature files that you can verify to ensure that no malicious entities have tampered with your downloaded Unity Accelerator version.

You can verify the signature files with a trusted version of GnuPG.

### Download the Unity Accelerator public key

Download and import the Unity Accelerator public key that signs the Linux installer files:

```
$ wget -qO- https://download-accelerator.unity3d.com/unity-accelerator-pub.asc | gpg --import -
```

You only need to import the key once.

### Verify the public key fingerprint

Check the fingerprint of the imported key:

```
gpg --fingerprint sst-ops@unity3d.com
```

You should see:

```
    pub   rsa4096 2019-12-02 [SC]
          9BC4 B04D F2E8 74E5 64AA  47E7 6D44 3B8B 002C F61F
    uid           [ unknown] Unity Technologies ApS (Accelerator signing key) <sst-ops@unity3d.com>
```

If the fingerprint matches, the key is authentic and trustable.
For more information about marking a key trustable, refer to <https://www.gnupg.org/gph/en/manual/x334.html>.

### Download and verify the Unity Accelerator installer signature

To check the version of Unity Accelerator that you installed is original and unmodified, verify the file’s signature. It’s available in the [downloads page](accelerator-download.md).

Then run the following command:

```
gpg --verify unity-accelerator-v1.5.63+g9672c1c8-linux-x64-installer.run.sig unity-accelerator-v1.5.63+g9672c1c8-linux-x64-installer.run
```

If everything is correct, you will see:

```
    gpg: Good signature from "Unity Technologies ApS (Accelerator signing key) <sst-ops@unity3d.com>"
```

This confirms the installer is genuine and unmodified.

### Forged or untrusted keys

You may also see a message like:

```
    WARNING: This key is not certified with a trusted signature!
    There is no indication that the signature belongs to the owner.
```

This **does not indicate a problem** with the signature.

It means that you have imported the Unity Accelerator public key, but you have not personally marked it as “trusted” in your keyring.

This is normal for public distribution. As long as you see:

```
    Good signature from "Unity Technologies ApS (Accelerator signing key) <sst-ops@unity3d.com>"
```

the file is verified correctly.

## Additional resources

* [Command line arguments reference](EditorCommandLineArguments.md#accelerator)
* [Configure Unity Accelerator in the Editor](accelerator-configure.md)
* [Stop and restart Unity Accelerator](accelerator-stop-restart.md)