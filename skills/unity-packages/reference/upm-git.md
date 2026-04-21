# Git dependencies

When the Package Manager fetches a package from a Git repository, it adds the package locally to your project. This allows you to test unpublished changes, but you can’t use it to contribute to that Git repository. To set up an existing local Git repository as a dependency in your project, use a [path to your local Git repository](upm-localpath.md) instead.

**Note**: You can’t specify **Git dependencies** in a [package.json](upm-manifestPkg.md) file because the Package Manager doesn’t support Git dependencies between packages. It supports Git dependencies only for projects, so you can declare Git dependencies only in the project’s [manifest.json](upm-manifestPrj.md) file.

**Tip**: If you want to update your Git dependency to a specific version ([revision](#revision)) from the repository, refer to [Locked Git dependencies](#git-locks).

This section includes the following topics:

* [Requirements](#req)
* [Git URLs](#syntax)
  + [Using the HTTP/HTTPS protocol](#Git-HTTPS)
  + [Using the SSH protocol](#Git-SSH)
  + [Using the FILE protocol](#Git-FILE)
  + [Using the GIT protocol](#Git-GIT)
* [Extended syntax](#extended)
  + [Specifying a Git revision](#revision)
  + [Specifying a package in a subfolder of the repository](#subfolder)
  + [Specifying revisions and paths at the same time](#paths-and-revisions)
* [Locked Git dependencies](#git-locks)
* [Git LFS Support](#git-lfs)

## Requirements

To use Git dependencies in a project, make sure you installed the [Git client](https://git-scm.com/) (minimum version 2.14.0) on your computer and that you have added the Git executable path to the PATH system environment variable.

**Warning:** Unity tested the Package Manager to work with Git 2.14.0 and above. Unity can’t guarantee the results if you use Git versions below 2.14.0.

If the repository tracks files with [Git LFS](https://git-lfs.github.com/), make sure the Git LFS client is also installed on your machine. If it’s not installed, the Package Manager [can’t retrieve the files stored on the LFS server](upm-errors.md#git-lfs) and instead checks out the LFS pointer files without any error or warning messages.

You can use the Package Manager window to install a package directly from a Git repository. For more information, refer to [Installing from a Git URL](upm-ui-giturl.md).

## Git URLs

The Package Manager supports all [Git protocols](https://git-scm.com/book/en/v2/Git-on-the-Server-The-Protocols), with the exception of local file paths. To specify a Git URL as a dependency, add the name of the package to add with a Git URL instead of the version number or local file path to the **project manifest**. For example, this demonstrates how to specify a remote Git using different protocols:

```
{
  "dependencies": {
    "com.mycompany.mypackage1": "https://github.example.com/myuser/myrepository1.git",
    "com.mycompany.mypackage2": "ssh://git@github.example.com/myuser/myrepository2.git",
    "com.mycompany.mypackage3": "file://localhost/github.example.com/myuser/myrepository3.git",
    "com.mycompany.mypackage4": "git://github.example.com/myuser/myrepository4.git",
    etc.
  }
}
```

The Package Manager recognizes that a dependency formatted as a URL is a Git URL by looking for the `.git` file extension at the end of the repository path. Some Git repository hosting services don’t support URLs with this extension while others enforce it. For this reason, the Git dependency syntax allows you to omit the extension if you use the [GIT protocol](#Git-GIT), or if you add a special `git+` prefix to the [HTTP/HTTPS](#Git-HTTPS), [SSH](#Git-SSH), or [FILE](#Git-FILE) URL.

**Note:** The `git+` prefix is a special marker in the `manifest.json` file that indicates that the dependency is Git based. The Package Manager doesn’t pass it to Git when cloning the repository.

For more information about the format of Git-supported URLs, refer to the documentation for the git [clone](https://git-scm.com/docs/git-clone#_git_urls_a_id_urls_a) command. For an overview of the difference between the protocols that Git uses, refer to [the Git documentation for using protocols](https://git-scm.com/book/en/v2/Git-on-the-Server-The-Protocols).

You can also use extended syntax for Git dependencies:

* If the package you want isn’t at the root of the repository, you can specify a path to a package’s subfolder in the repository. This is necessary only if the package you want isn’t at the root of the repository. For example, the string `?path=/folder1/folder2` in:

  `"https://github.example.com/myuser/myrepository.git?path=/folder1/folder2"`.

  For more information, refer to [Specifying a package in a subfolder](#subfolder).
* You can specify a Git revision, which can be a tag, branch name, or a specific commit hash to lock onto. This ensures that the Package Manager always loads that exact revision. If you don’t specify a revision, the Package Manager clones the repository at the default branch and latest commit and locks onto that revision. For example, the string `#v2.0.0` in:

  `"https://github.example.com/myuser/myrepository.git#v2.0.0"`

  For more information, refer to [Specifying a Git revision](#revision).

### Using the HTTP/HTTPS protocol

You can use the HTTPS protocol with the full URL:

```
{
  "dependencies": {
    "com.mycompany.mypackage": "https://github.example.com/myuser/myrepository.git"
  }
}
```

If your Git server doesn’t support the `.git` extension, you can add the special `git+` prefix, with or without the extension:

```
{
  "dependencies": {
    "com.mycompany.mypackage1": "git+https://github.example.com/myuser/myrepository1.git",
    "com.mycompany.mypackage2": "git+https://github.example.com/myuser/myrepository2"
  }
}
```

**Note**: Alternatively, you can use the GIT protocol instead of the `git+` prefix. For more information, refer to [Using the GIT protocol](#Git-GIT).

If the repository is publicly accessible, HTTPS is the recommended scheme for sharing Git URLs with users because you can copy and paste the URL directly from the Git repository hosting service web page.

![Copy the URL from the package repository](../uploads/Main/upm-git-cloning.png)

Copy the URL from the package repository

If the repository isn’t publicly accessible and you are using HTTPS, the repository server fails to authenticate you because you can’t interact with the server to provide your credentials. In this case, the Editor notifies you that authentication failed.

To work around these authentication issues, you can either:

* Use a [Git credentials helper](https://git-scm.com/docs/gitcredentials) to authenticate beforehand. For more information, refer to [Use private repositories with HTTPS Git URLs](upm-config-https-git.md).
* Use the [SSH protocol](#Git-SSH) instead. If you set up and configure an SSH key pair with the Git repository hosting service, the Package Manager can authenticate the request seamlessly on your behalf.

### Using the SSH protocol

You can use the SSH protocol with the full URL:

```
{
  "dependencies": {
    "com.mycompany.mypackage": "ssh://git@mycompany.github.com/gitproject/com.mycompany.mypackage.git"
  }
}
```

If your Git server doesn’t support the `.git` extension, you can add the special `git+` prefix, with or without the extension:

```
{
  "dependencies": {
    "com.mycompany.mypackage1": "git+ssh://git@github.example.com/myuser/myrepository1.git",
    "com.mycompany.mypackage2": "git+ssh://git@github.example.com/myuser/myrepository2"
  }
}
```

**Note**: Alternatively, you can use the GIT protocol instead of the `git+` prefix. For more information, refer to [Using the GIT protocol](#Git-GIT).

You can also use the SCP-like shorthand, which the Package Manager always recognizes as a Git dependency:

```
{
  "dependencies": {
    "com.mycompany.mypackage": "git@mycompany.github.com:gitproject/com.mycompany.mypackage.git"
  }
}
```

#### Using PuTTY on Windows

Git uses the keys at the default location when you use SSH to authenticate. However, if you are using [PuTTY](https://www.putty.org/) as the SSH client on Windows, you need to configure the `GIT_SSH` environment variable to make it point to `plink.exe`.

#### Authenticating with SSH

If you want to use the SSH protocol, you need to set up SSH keys outside of Unity. For more information on setting up authentication for a specific host, refer to the help pages for [Bitbucket](https://support.atlassian.com/bitbucket-cloud/docs/set-up-an-ssh-key/), [GitLab](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html), and [GitHub](https://github.com/settings/tokens).

**Note**: If you encrypted your SSH key with a passphrase, the Package Manager can’t retrieve the package, because it doesn’t provide a way to enter the passphrase in a terminal or command line. In this case, the Editor notifies you that authentication failed. For information about using an authentication agent, refer to [Using passphrase-protected SSH keys with SSH Git URLs](upm-config-ssh-git.md). For additional information about using the **ssh-agent** for authentication, refer to [Solutions for SSH](upm-errors.md#ssh-add).

### Using the FILE protocol

The Package Manager doesn’t recognize Git URLs with the `file:` prefix as Git dependencies unless they’re properly formatted. This means you must use either the `git+file:` protocol, or the `.git` suffix with the `file:` protocol:

```
{
  "dependencies": {
    "com.mycompany.mypackage1": "git+file://github.example.com/myuser/myrepository1",
    "com.mycompany.mypackage2": "git+file:///github.example.com/myuser/myrepository2",
    "com.mycompany.mypackage3": "file:///github.example.com/myuser/myrepository3.git"
  }
}
```

**Note**: Alternatively, you can use the GIT protocol instead of the `git+` prefix. For more information, refer to [Using the GIT protocol](#Git-GIT).

The Package Manager interprets any other syntax as a [local path](upm-localpath.md) instead.

### Using the GIT protocol

The Package Manager recognizes the `git:` protocol, with or without the `.git` path suffix:

```
{
  "dependencies": {
    "com.mycompany.mypackage1": "git://github.example.com/myuser/myrepository1.git",
    "com.mycompany.mypackage2": "git://github.example.com/myuser/myrepository2"
  }
}
```

The GIT protocol doesn’t need or support the `git+` prefix.

## Extended syntax

You can use extended syntax to identify a [specific Git revision](#revision), a [package in a subfolder](#subfolder), or [both](#paths-and-revisions).

You can use extended syntax with any Git protocol that Unity supports.

### Specifying a Git revision

To declare a specific revision that you want the Package Manager to clone, add the revision prefixed by the number sign (`#`) at the end of the URL:

```
{
  "dependencies": {
    "com.mycompany.mypackage1": "https://github.example.com/myuser/myrepository1.git#revision",
    "com.mycompany.mypackage2": "git+https://github.example.com/myuser/myrepository2#revision"
  }
}
```

The revision can be any tag, branch or commit hash. You must provide a full commit hash. Unity doesn’t support shortened SHA–1 hashes. This table shows examples for specifying revisions:

| **Syntax** | **URL example** |
| --- | --- |
| Latest default branch | `"https://github.example.com/myuser/myrepository.git"` |
| Specific branch | `"https://github.example.com/myuser/myrepository.git#my-branch"` |
| Specific tag | `"https://github.example.com/myuser/myrepository.git#tag-pointing-to-package-version"` |
| Commit hash | `"https://github.example.com/myuser/myrepository.git#9e72f9d5a6a3dadc38d813d8399e1b0e86781a49"` |

### Specifying a package in a subfolder of the repository

If you specify a repository using the Git URL syntax, the Package Manager assumes that the package must be at the root of the repository. However, some packages aren’t located at the root level of their repository, and some repositories contain more than one package.

You can use the `path` query parameter in the Git URL to notify the Package Manager where to find the package. The path you specify must be relative to the root of the repository, and the subfolder you specify must contain the **package manifest** (`package.json` file).

To specify a repository subfolder for a Git dependency, use the `path` query parameter:

```
{
  "dependencies": {
    "com.mycompany.mypackage": "https://github.example.com/myuser/myrepository.git?path=/subfolder"
  }
}
```

In this case, the Package Manager registers the package located in the specified repository subfolder and disregards the rest of the repository.

Sometimes a repository contains several related packages. If you want to add more than one package from the same repository, you must add two separate entries to your project manifest:

```
{
  "dependencies": {
    "com.mycompany.mypackage1": "https://github.example.com/myuser/myrepository.git?path=/subfolder1",
    "com.mycompany.mypackage3": "https://github.example.com/myuser/myrepository.git?path=/subfolder2/subfolder3"
  }
}
```

**Note**: If you specify the same repository multiple times, the Package Manager clones the same repository multiple times, which leads to reduced performance and additional network usage.

### Specifying revisions and paths at the same time

You can specify paths and revisions with any Git protocol that Unity supports. However, the `path` query parameter always precedes the revision anchor. The reverse order fails. This is an example of the correct order to use:

```
{
  "dependencies": {
    "com.mycompany.mypackage": "https://github.example.com/myuser/myrepository.git?path=/example/folder#v1.2.3"
  }
}
```

## Locked Git dependencies

One of the core principles of the Package Manager is determinism. If you share your project with other users, the Package Manager should install the same set of package dependencies and versions, and that includes packages that it fetches from Git. To achieve this, the Package Manager tracks commit hashes of Git dependencies by using a [lock file](upm-conflicts-auto.md).

When you add a Git dependency with a revision set to a branch or a tag, the Package Manager fetches the corresponding commit hash to store in the lock file. Over time, branches and tags can point to different commits on the Git repository. For example, a branch can have newer commits added to it.

To update the package to a different commit that a branch or tag points to, use the [Install package from git URL](upm-ui-giturl.md) button and enter a Git URL. You can use the same Git URL, because the Package Manager ignores the locked commit hash when you submit a new request. However, you can also specify a new revision number, tag, or branch as a [revision](#revision) instead.

Alternatively, you could create a script with the [Client.Add](../ScriptReference/PackageManager.Client.Add.md) C# API method with that Git URL.

## Git LFS Support

The Package Manager supports Git dependencies from repositories that use [Git Large File Storage (LFS)](https://git-lfs.com/). However, because Unity uses shallow clones, LFS-tracked files aren’t always fetched correctly. Sometimes, the package might import placeholder pointer files instead of the actual file contents. For more information, refer to [Git LFS errors](upm-errors.md#lfs-errors).

Since Git LFS is designed to work with minimal configuration overhead, it supports both HTTPS and SSH authentication:

* For information about HTTPS authentication, refer to [Using the HTTP/HTTPS protocol](#Git-HTTPS).
* For information about SSH authentication, refer to [Using the SSH protocol](#Git-SSH).

Retrieval of files stored on the LFS server fails if users need authentication and don’t have valid credentials with permission to access the remote repository.

Package authors can help the Git LFS client locate the LFS Server by providing it with the URL in the `.lfsconfig` configuration file in the repository. There are two ways to do this:

```
# Option 1: global setting
[lfs]
  url = ssh://git@HOSTNAME/path/to/repo.git

# Option 2: per-remote setting
[remote "origin"]
  lfsurl = ssh://git@HOSTNAME/path/to/repo.git
```

If the repository contains a `.lfsconfig` file, make sure you include it in the `.npmignore` file to avoid including it in published releases of the package.

### Git LFS cache

As of Unity 2021.2, you can optionally enable a Git LFS cache for the Package Manager to use when checking out Git-based dependencies. This avoids having to download the same file every time you check out a different revision of the repository.

The Git LFS cache for the Package Manager is different from the Git LFS cache in the `.git/lfs` folder of your Git repository. The Package Manager can’t use the default Git cache because it doesn’t keep cloned repositories after it copies packages to the project cache.

To enable the Git LFS cache for the Package Manager, choose one of the following options:

* To enable the Git LFS cache and use the `git-lfs` subfolder under the [default global cache root](upm-cache.md) as its location, set the `UPM_ENABLE_GIT_LFS_CACHE` environment variable to any (non-empty) value.
* To enable the Git LFS cache and use a custom location for it, set the `UPM_GIT_LFS_CACHE_PATH` environment variable to the custom path. When you set the location, the Git LFS cache option is automatically enabled.

For more information about setting environment variables for the global cache, refer to [Customize the global cache](upm-config-cache.md).

**Note:** This optimization requires extra disk space when using Git LFS-enabled packages. You need to decide which is the greater benefit: Git LFS file caching costs disk space but saves you from downloading the same files again. However, some situations can’t make use of the cache and use up disk space without reusing the files. For example, your Git dependencies might resolve to revisions that reference different LFS-tracked file content, such as these scenarios:

* Using different Git revisions in dependencies in multiple projects
* Frequently updating a package to revisions that contain different changed LFS files

## Additional resources

* [Use private repositories with HTTPS Git URLs](upm-config-https-git.md)
* [Use passphrase-protected SSH keys with SSH Git URLs](upm-config-ssh-git.md)
* [Git Large File Storage (LFS)](https://git-lfs.github.com/)
* [Git on the Server - The Protocols](https://git-scm.com/book/en/v2/Git-on-the-Server-The-Protocols)
* [Installing from a Git URL](upm-ui-giturl.md)