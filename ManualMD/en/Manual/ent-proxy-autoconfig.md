# Use Unity through web proxies

Some organizations use a web proxy, which is an intermediary server between users and the internet resource they’re requesting. Organizations use web proxies for several reasons, including security, privacy, and network efficiency.

If you are deploying Unity in an environment that requires internet access through a web proxy, you might need to configure the web proxy, client computers, or both. Depending on your environment, this configuration includes some or all of these tasks:

1. [Prepare your environment so Unity applications can use your web proxy](#prep-env)
2. [Configure client computers to trust the web proxy certificate](#interception)
3. [Allow the web proxy to download executable files](#allow-exe)

## 1) Prepare your environment so Unity applications can use your web proxy

The automatic proxy configuration feature lets you use Unity applications seamlessly when they need to communicate through a web proxy. Unity applications can communicate with the web proxy without interruption and without installing additional software. The amount of configuration required depends on the following factors:

* **Unity version and operating system**: Unity applications can access web proxy configuration from your operating system. Support for different operating systems varies based on your version of the Unity Editor.
  + On Windows, Editor versions 2022.3.21f1 or later support automatic proxy configuration.
  + On macOS, support for automatic proxy configuration within the 2022 stream of the Editor begins with 2022.3.21f1. In the 2023 (and Unity 6) stream, support for automatic proxy configuration begins with 2023.3.0a18, including Unity 6 Beta.
  + Other Editor versions and other operating systems require extra configuration.
* **Authentication type**: Check if your proxy supports basic authentication (username and password), or if it uses a SPNEGO mechanism and a protocol, such as Kerberos or NTLM.

Use the following tables (grouped by operating system) to decide which solution matches your environment:

### Windows

| Unity version | Proxy authentication | Solution | Other information |
| --- | --- | --- | --- |
| Editor 2022.3.21f1 or later and Hub version 3.4.2 or later | Accepts basic authentication | Do these tasks:  * [Enable automatic proxy configuration](ent-proxy-autoconfig-enable.md). * [Store your credentials in Credential Manager (Windows)](ent-proxy-autoconfig-store.md). | Refer to footnote 1 that follows this table. |
| Editor 2022.3.21f1 or later and Hub version 3.4.2 or later | Doesn’t accept basic authentication, or doesn’t require authentication | Do these tasks:  * [Enable automatic proxy configuration](ent-proxy-autoconfig-enable.md). * If your proxy requires authentication, [allow unauthenticated access for Unity applications in a proxy exception list](ent-proxy-exception-list.md). | Refer to footnote 1 that follows this table. |
| Editor 2022.3.20f1 or earlier or Hub 3.4.1 or earlier | Accepts basic authentication | Do these tasks:  * [Use environment variables to identify the web proxy](ent-proxy-env-vars.md). * [Allow unauthenticated access for Unity applications in a proxy exception list](ent-proxy-exception-list.md) (preferred), or [store username and password in the environment variables](ent-proxy-env-vars.md) for the proxy. | Refer to footnote 2 that follows this table. |
| Editor 2022.3.20f1 or earlier or Hub 3.4.1 or earlier | Requires authentication but doesn’t accept basic authentication | Do these tasks:  * [Use environment variables to identify the web proxy](ent-proxy-env-vars.md). * [Allow unauthenticated access for Unity applications in a proxy exception list](ent-proxy-exception-list.md) | Refer to footnote 2 that follows this table. |

1. You must be able to store proxy information in the operating system settings or in environment variables.
2. Automatic proxy configuration isn’t fully supported on Windows before Editor version 2022.3.

### macOS

| Unity version | Proxy authentication | Solution | Other information |
| --- | --- | --- | --- |
| Editor 2022.3.21f1, 2023.3.0a18 (or later) and Hub version 3.4.2 or later | Accepts basic authentication | Do these tasks:  * [Enable automatic proxy configuration](ent-proxy-autoconfig-enable.md). * [Store your credentials for automatic proxy configuration (macOS)](ent-proxy-autoconfig-store-mac.md). | Refer to footnote 1 that follows this table. |
| Editor 2022.3.21f1, 2023.3.0a18 (or later) and Hub version 3.4.2 or later | Doesn’t accept basic authentication, or doesn’t require authentication | Do these tasks:  * [Enable automatic proxy configuration](ent-proxy-autoconfig-enable.md). * If your proxy requires authentication, [allow unauthenticated access for Unity applications in a proxy exception list](ent-proxy-exception-list.md). | Refer to footnote 1 that follows this table. |
| Editor 2023.2, 2023.1, 2022.3.20f1 or earlier or Hub 3.4.1 or earlier | Accepts basic authentication | Do these tasks:  * [Use environment variables to identify the web proxy](ent-proxy-env-vars.md). * [Allow unauthenticated access for Unity applications in a proxy exception list](ent-proxy-exception-list.md) (preferred), or [store username and password in the environment variables](ent-proxy-env-vars.md) for the proxy. | Refer to footnote 2 that follows this table. |
| Editor 2023.2, 2023.1, 2022.3.20f1 or earlier or Hub 3.4.1 or earlier | Requires authentication but doesn’t accept basic authentication | Do these tasks:  * [Use environment variables to identify the web proxy](ent-proxy-env-vars.md). * [Allow unauthenticated access for Unity applications in a proxy exception list](ent-proxy-exception-list.md) | Refer to footnote 2 that follows this table. |

1. You must be able to store proxy information in the operating system settings or in environment variables.
2. Automatic proxy configuration isn’t fully supported on macOS for Editor versions 2023.2 and 2023.1 and any version before 2022.2.

### Linux

| Unity version | Proxy authentication | Solution | Other information |
| --- | --- | --- | --- |
| Any | Any | Do these tasks:  * [Use environment variables to identify the web proxy](ent-proxy-env-vars.md). * If your proxy requires authentication, [allow unauthenticated access for Unity applications in a proxy exception list](ent-proxy-exception-list.md). |  |

## 2) Configure client computers to trust the web proxy certificate

Check if your web proxy uses SSL inspection, which some enterprises use to inspect the contents of secure (TLS) requests. If your enterprise uses SSL inspection, configure client computers to [trust the web proxy certificate](ent-proxy-cert-trust.md).

## 3) Allow the web proxy to download executable files

Check if your web proxy prevents downloading of executable files, or requires user confirmation before downloading executable files. If so, you need to configure the web proxy to allow this traffic on specific endpoints. Otherwise, users won’t be able to run the installers that the Unity Hub uses to install components required by the Unity Editor. For details, refer to [Endpoints that download executable files](ent-proxy-exception-list.md#endpoints-exe).

## Known limitations

* For organizations that configure the web proxy with a proxy auto configuration (PAC) file, the [`HTTP` keyword](https://developer.mozilla.org/en-US/docs/Web/HTTP/Proxy_servers_and_tunneling/Proxy_Auto-Configuration_PAC_file#history_and_implementation:~:text=support%20as%20well%3A-,HTTP,-host%3Aport) isn’t supported. Instead, use the [`PROXY` keyword](https://chromium.googlesource.com/chromium/src/+/HEAD/net/docs/proxy.md#HTTP-proxy-scheme:~:text=Example%20identifier%20(PAC)%3A-,PROXY,-proxy%3A8080%2C) to specify a proxy that communicates over HTTP.
* For organizations that configure multiple web proxies with the same target URL, Unity applications will use only the first proxy discovered. If that web proxy isn’t reachable, Unity applications will not attempt to use any of the other web proxies.

## Additional resources

* [Enable automatic proxy configuration](ent-proxy-autoconfig-enable.md)
* [Store credentials for automatic proxy configuration (Windows)](ent-proxy-autoconfig-store.md)
* [Store credentials for automatic proxy configuration (macOS)](ent-proxy-autoconfig-store-mac.md)
* [Use environment variables to identify your web proxy](ent-proxy-env-vars.md)
* [Define exceptions on your web proxy](ent-proxy-exception-list.md)
* [Trusting the web proxy security certificate](ent-proxy-cert-trust.md)
* [Unity Licensing Client requirements](https://docs.unity3d.com/licensing/manual/Requirements.html#client-requirements) (Unity Floating Licensing)