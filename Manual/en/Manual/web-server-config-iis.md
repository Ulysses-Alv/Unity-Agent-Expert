# Set up your IIS server configuration for Web builds

Set up your Internet Information Services (IIS) server configuration file to use the IIS server with Unity Web builds.

The IIS server is ideal for Windows-based environments and for use with Microsoft products.

For more information on IIS and how to install and use the server, refer to the Microsoft [IIS documentation](https://www.iis.net/overview).

## Configure your IIS server for compressed Web builds without decompression fallback

To update your IIS server configuration file so that it works with a compressed Web build without decompression fallback, follow these steps:

1. Use the following server configuration:

   ```
   <?xml version="1.0" encoding="UTF-8"?>

   <configuration>
   <system.webServer>
     <!--
       Compressed Unity builds without decompression fallback can't be properly hosted on a server which
       has static compression enabled because this might result in the build files being compressed twice.
       The following line disables static server compression.
     -->
     <urlCompression doStaticCompression="false" />
     <!-- To host compressed Unity builds, the correct mimeType should be set for the compressed build files. -->
     <staticContent>
       <!--
         NOTE: IIS will throw an exception if a mimeType is specified multiple times for the same extension.
         To avoid possible conflicts with configurations that are already on the server, you should remove the mimeType for the corresponding extension using the <remove> element,
         before adding mimeType using the <mimeMap> element.
       -->
       <!-- The following lines are required for builds compressed with gzip, which don't include decompression fallback. -->
       <remove fileExtension=".data.gz" />
       <mimeMap fileExtension=".data.gz" mimeType="application/gzip" /><!-- The correct MIME type here would be application/octet-stream, but due to Safari bug https://bugs.webkit.org/show_bug.cgi?id=247421, it's preferable to use MIME Type application/gzip instead. -->
       <remove fileExtension=".wasm.gz" />
       <mimeMap fileExtension=".wasm.gz" mimeType="application/wasm" />
       <remove fileExtension=".js.gz" />
       <mimeMap fileExtension=".js.gz" mimeType="application/javascript" />
       <remove fileExtension=".symbols.json.gz" />
       <mimeMap fileExtension=".symbols.json.gz" mimeType="application/octet-stream" />
       <!-- The following lines are required for builds compressed with Brotli, which don't include decompression fallback. -->
       <remove fileExtension=".data.br" />
       <mimeMap fileExtension=".data.br" mimeType="application/octet-stream" />
       <remove fileExtension=".wasm.br" />
       <mimeMap fileExtension=".wasm.br" mimeType="application/wasm" />
       <remove fileExtension=".js.br" />
       <mimeMap fileExtension=".js.br" mimeType="application/javascript" />
       <remove fileExtension=".symbols.json.br" />
       <mimeMap fileExtension=".symbols.json.br" mimeType="application/octet-stream" />
     </staticContent>

     <!--
       Hosting compressed Unity builds without decompression fallback relies on native browser decompression,
       therefore a proper "Content-Encoding" response header should be added for the compressed build files.
       NOTE: IIS will throw an exception if the following section is used without the "URL Rewrite" module installed.
       Download the "URL Rewrite" module from https://www.iis.net/downloads/microsoft/url-rewrite
     -->
     <rewrite>
       <outboundRules>
         <!--
           NOTE: IIS will throw an exception if the same rule name is used multiple times.
           To avoid possible conflicts with configurations that are already on the server, you should remove the mimeType for the corresponding extension using the <remove> element,
         before adding mimeType using the <mimeMap> element.
         -->
         <!-- The following section is required for builds compressed with gzip, which don't include decompression fallback. -->
         <remove name="Append gzip Content-Encoding header" />
         <rule name="Append gzip Content-Encoding header">
           <match serverVariable="RESPONSE_Content-Encoding" pattern=".*" />
           <conditions>
             <add input="{REQUEST_FILENAME}" pattern="\.gz$" />
           </conditions>
           <action type="Rewrite" value="gzip" />
         </rule>
         <!-- The following section is required for builds compressed with Brotli, which don't include decompression fallback. -->
         <remove name="Append brotli Content-Encoding header" />
         <rule name="Append brotli Content-Encoding header">
           <match serverVariable="RESPONSE_Content-Encoding" pattern=".*" />
           <conditions>
             <add input="{REQUEST_FILENAME}" pattern="\.br$" />
           </conditions>
           <action type="Rewrite" value="br" />
         </rule>
       </outboundRules>
     </rewrite>
   </system.webServer>
   </configuration>
   ```
2. Upload this configuration file to the server as `<Application Folder>/Build/web.config`.

Your server configuration is ready for interactions with your Web build.

**Note**: To host compressed Web builds without decompression fallback, you need to install the `URL Rewrite` IIS module on the server.
Otherwise, IIS will throw an exception when using this configuration file.
This module is available in the Microsoft documentation on [URL Rewrite](https://www.iis.net/downloads/microsoft/url-rewrite).

## Configure your IIS server configuration for uncompressed Web builds

To update your server configuration file so that it works with an uncompressed Web build, follow these instructions:

1. Use the following server configuration:

   ```
   <?xml version="1.0" encoding="UTF-8"?>
   <configuration>
   <system.webServer>
     <!--
       IIS does not provide default handlers for .data and .wasm files (and in some cases .json files),
       therefore these files won’t be served unless their mimeType is explicitly specified.
     -->
     <staticContent>
       <!--
         NOTE: IIS will throw an exception if a mimeType is specified multiple times for the same extension.
         To avoid possible conflicts with configurations that are already on the server, you should remove the mimeType for the corresponding extension using the <remove> element,
         before adding mimeType using the <mimeMap> element.
       -->
       <remove fileExtension=".data" />
       <mimeMap fileExtension=".data" mimeType="application/octet-stream" />
       <remove fileExtension=".wasm" />
       <mimeMap fileExtension=".wasm" mimeType="application/wasm" />
       <remove fileExtension=".symbols.json" />
       <mimeMap fileExtension=".symbols.json" mimeType="application/octet-stream" />
     </staticContent>
   </system.webServer>
   </configuration>
   ```
2. Upload this configuration file to the server as `<Application Folder>/Build/web.config`.

Your server configuration is ready for interactions with your Web build.

## Additional resources

* [Compressed builds and server configuration](webgl-deploying.md)
* [Server configuration code samples](webgl-server-configuration-code-samples.md)
* [Set up your Nginx server configuration for Web builds](web-server-config-nginx.md)
* [Set up your Apache server configuration for Web builds](web-server-config-apache.md)
* [Set up your Node.js server configuration for Web builds](web-server-config-nodejs.md)