# Use private repositories with HTTPS Git URLs

When you use Git in a terminal to access a private repository over HTTPS, Git prompts you to enter a username and password. Then, Git submits these credentials to the server and proceeds with the command if the server accepts those credentials and allows access to the repository.

When the Unity Package Manager fetches packages using Git URLs, there’s no terminal for users to enter credentials. As a result, when the server requests credentials from Git, Git doesn’t issue a prompt. Instead, it reports an error to the Unity Package Manager. To solve this problem, you must configure Git with a Git credential helper, and that helper must already have the required credentials loaded for that repository. If the credentials are valid, Git can successfully run the commands issued by the Unity Package Manager.

**Note**: A Git credential helper has no effect when using Git URLs with the SSH protocol, including the SCP-like syntax. For information about accessing a private Git repository over SSH, refer to [Using passphrase-protected SSH keys with SSH Git URLs](upm-config-ssh-git.md).

## Git Credential Manager

Although Git supports several [credential helpers to store credentials](https://git-scm.com/book/en/v2/Git-Tools-Credential-Storage), the [Git Credential Manager](https://github.com/git-ecosystem/git-credential-manager) (GCM) is the recommended credential helper. GCM is flexible, easy to install, and is actively supported. It’s built on .NET, which means it can run on Windows, macOS, and Linux distributions that support .NET.

By default, GCM uses Windows Credential Manager (on Windows) and macOS Keychain (on macOS) as the configured credential store. GCM doesn’t have a default store configured for Linux. Refer to the [GCM credential stores documentation](https://github.com/git-ecosystem/git-credential-manager/blob/release/docs/credstores.md) for more information on the different credential store configurations supported by GCM.

## Prerequisites

Before you can fetch packages from private Git repositories using HTTPS URLs, make sure you install GCM.

The [Git for Windows](https://gitforwindows.org/) installer includes a step to install and configure GCM automatically. You can also install GCM separately if you:

* Used a different method to install Git on Windows.
* Use macOS or Linux.

Refer to the [GCM install instructions](https://github.com/git-ecosystem/git-credential-manager/blob/release/docs/install.md) for more information about installing GCM.

## Procedures

The procedure you follow depends on whether your repository uses two-factor authentication (2FA).

### Procedure for repositories that don’t use 2FA

Follow these steps to access packages in private repositories that use HTTPS Git URLs without 2FA:

1. Configure Git to use GCM by running the following command in a terminal:

   ```
   git config --global credential.helper manager
   ```
2. Access the repository one time by using a terminal. For example, run the following command:

   ```
   git ls-remote --heads https://<url-to-repository> HEAD
   ```
3. When Git prompts you, enter your credentials. If your user account has access to the remote Git server and the server accepts your credentials, then the Git credential helper will securely store your credentials.
4. Use the Unity Package Manager. It will use your stored credentials, when necessary, to fetch packages from HTTPS-based repositories that you have permissions to access.

### Procedure for repositories that use 2FA

If your repository has 2FA enabled, you can’t use your credentials to access packages via HTTPS. Instead, you must use a Personal Access Token (PAT) as your password.

Follow these steps to access packages in private repositories that use HTTPS Git URLs with 2FA:

1. Create a Personal Access Token:

   * Go to your Git hosting provider’s website, such as GitHub or GitLab.
   * Open your account’s Developer Settings, Access Tokens, or equivalent section.
   * Create a new token with the following information:

     + **Name** or **Description**: Something meaningful (for example, Unity UPM Access).
     + **Expiration date**: Optional, but recommended for security.
     + **Scopes** or **Permissions**: Typically `read_repository` or similar if you just need clone or pull access.
   * Save the token somewhere secure. Most providers will show the token only once.
2. Configure Git to Use Your PAT:

   * When Git prompts for credentials when cloning, pulling, or pushing:

     + **Username**: Your account username (or sometimes your email, depending on the host).
     + **Password**: Your PAT (instead of your regular password).
3. Store Your Credentials (Optional):  
     
   To avoid entering your credentials each time, you can use Git’s credential manager:

   * Windows

     ```
       git config --global credential.helper manager-core
     ```
   * macOS

     ```
       git config --global credential.helper osxkeychain
     ```

## Additional resources

* [Troubleshooting Git issues](upm-errors.md#git-not-found)
* [Use passphrase-protected SSH keys with SSH Git URLs](upm-config-ssh-git.md)