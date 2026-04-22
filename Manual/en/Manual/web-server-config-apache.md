# Set up your Apache server configuration for Web builds

Set up your Apache server configuration file to use the Apache server with Unity Web builds.

Apache is an open source, cross-platform web server that you can use to serve applications you create with Unity Web.

For more information about Apache and how to install and use it, refer to the [Apache documentation](https://httpd.apache.org/docs/).

## Choose your Apache Web build method

You can configure Web builds in Apache using one of the two methods:

* Virtual host in `httpd.conf`.
* Virtual host in `.htaccess`.

The recommended best practice is to configure a virtual host in Apache’s `httpd.conf`, if you have access to that configuration.

### Use the virtual host in `httpd.conf` method

To configure Web builds using the virtual host in `httpd.conf` method, use the following code in your Apache server configuration:

```
<Directory /var/www/html/root/path/to/your/unity/content/>
<IfModule mod_mime.c>
RemoveType .gz
AddEncoding gzip .gz
# The correct MIME type for .data.gz would be application/octet-stream, but due to Safari bug https://bugs.webkit.org/show_bug.cgi?id=247421, it is preferable to use MIME Type application/gzip instead.
#AddType application/octet-stream .data.gz
AddType application/gzip .data.gz
AddType application/wasm .wasm.gz
AddType application/javascript .js.gz
AddType application/octet-stream .symbols.json.gz

RemoveType .br
AddEncoding br .br
AddType application/octet-stream .data.br
AddType application/wasm .wasm.br
AddType application/javascript .js.br
AddType application/octet-stream .symbols.json.br

# Insert additional code here

</IfModule>
</Directory>
```

### Use the virtual host in `.htaccess` method

If you’re unable to modify `httpd.conf`, for example, due to unavailable admin rights to your web hosting, but your Apache supports `.htaccess`, then you can configure `.htaccess` in the following way:

```
# NOTE: "mod_mime" Apache module must be enabled for this configuration to work.
<IfModule mod_mime.c>

# The following lines are required for builds without decompression fallback, compressed with gzip
RemoveType .gz
AddEncoding gzip .gz
AddType application/gzip .data.gz # The correct MIME type here would be application/octet-stream, but due to Safari bug https://bugs.webkit.org/show_bug.cgi?id=247421, it's preferable to use MIME Type application/gzip instead.
AddType application/wasm .wasm.gz
AddType application/javascript .js.gz
AddType application/octet-stream .symbols.json.gz

# The following lines are required for builds without decompression fallback, compressed with Brotli
RemoveType .br
RemoveLanguage .br
AddEncoding br .br
AddType application/octet-stream .data.br
AddType application/wasm .wasm.br
AddType application/javascript .js.br
AddType application/octet-stream .symbols.json.br

# The following line improves loading performance for uncompressed builds
AddType application/wasm .wasm

# Insert additional code here

</IfModule>
```

## Improve performance for gzip-compressed builds

To improve loading performance of a gzip-compressed Web build with decompression fallback, add the following line to your code:

```
AddEncoding gzip .unityweb
```

## Improve performance for Brotli-compressed builds with decompression fallback

To improve loading performance for Brotli-compressed builds with decompression fallback, add the following line of code to your configuration:

```
 AddEncoding br .unityweb
```

## Enable native C/C++ multithreading

If you built your project with **Enable Native C/C++ Multithreading** activated in the [**Player** settings](class-PlayerSettingsWebGL.md), you need to set specific server headers. These headers include Cross-Origin Opener Policy (COOP), Cross-Origin Embedded Policy (COEP), and Cross-Origin Resource Policy (CORP).

Add this code to your configuration file:

```
 <FilesMatch "\.(htm|html|js|js.gz|js.br)$">
     Header add Cross-Origin-Opener-Policy "same-origin"
     Header add Cross-Origin-Embedder-Policy "require-corp"
     Header add Cross-Origin-Resource-Policy "cross-origin"
 </FilesMatch>
```

## Allow Cross-Origin Resource Sharing (CORS) requests

To allow CORS requests, add the following code to your configuration:
`lang-apacheconf
Header add Access-Control-Allow-Origin "*"`

## Additional resources

* [Compressed builds and server configuration](webgl-deploying.md)
* [Server configuration code samples](webgl-server-configuration-code-samples.md)
* [Set up your IIS server configuration for Web builds](web-server-config-iis.md)
* [Set up your Nginx server configuration for Web builds](web-server-config-nginx.md)
* [Set up your Node.js server configuration for Web builds](web-server-config-nodejs.md)