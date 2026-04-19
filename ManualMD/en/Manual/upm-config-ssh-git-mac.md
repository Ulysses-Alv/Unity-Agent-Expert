# Load SSH keys automatically on macOS

If you use macOS, follow these steps to configure the OpenSSH client so you can use your passphrase-protected SSH key without prompts.

## Prerequisites

* macOS 10.13 or later.

## Before you begin

Check if you have any existing SSH keys. Refer to the GitHub Docs article, [Checking for existing SSH keys](https://docs.github.com/en/enterprise-server@3.8/authentication/connecting-to-github-with-ssh/checking-for-existing-ssh-keys).

If you don’t have an SSH key:

1. Create one by following the GitHub Docs article, [Generating a new SSH key](https://docs.github.com/en/enterprise-server@3.8/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent).
2. Add the new SSH key to your GitHub account, by following the GitHub Docs article, [Adding a new SSH key to your GitHub account](https://docs.github.com/en/enterprise-server@3.8/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account).

## Procedure

1. Check your home folder for an `.ssh` subdirectory and check if it contains a `config` file. **Note**: `.ssh` is a hidden subdirectory.
2. Create the `~/.ssh/config` file if it doesn’t exist.
3. Add the following content to `~/.ssh/config` to set the key to load in the authentication agent and specify its use by the target server. Make sure you add this entry before any global settings marked as `Host *`.

   ```
   Host SERVER_NAME
       UseKeychain yes
       IdentitiesOnly yes
       IdentityFile FILE_PATH
   ```

   * `SERVER_NAME` is the server that uses the file specified by `IdentityFile`. A sample value is `github.com`.
   * `FILE_PATH` is the path to the SSH file you created. A sample value is `~/.ssh/<FILE>`, where `<FILE>` might be `id_rsa`, `id_ecdsa`, `id_ed25519`, or a custom name.

   Example:

   ```
   Host github.com
       UseKeychain yes
       IdentitiesOnly yes
       IdentityFile ~/.ssh/id_ed25519
   ```
4. Open the Terminal application.
5. Load your key file into the `ssh-agent`, replacing `<FILE>` with the actual file name of your key, then type your passphrase, if prompted.

   ```
   ssh-add ~/.ssh/<FILE>
   ```

   Example:

   ```
   ssh-add ~/.ssh/id_ed25519
   ```

Configuration is complete. macOS starts `ssh-agent` by default, so you can now use the Unity Package Manager to fetch packages from that Git repository over SSH using your passphrase-protected SSH key.

## Additional resources

* [Use passphrase-protected SSH keys with SSH Git URLs](upm-config-ssh-git.md)