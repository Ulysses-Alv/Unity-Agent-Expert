# Monitor Unity Accelerator

Unity **Accelerator** has a built-in dashboard to monitor and configure changes.

The URL to access your dashboard is `http://hostname[:port]/dashboard` where `hostname:port` is the hostname/IP and port number of the Unity Accelerator you have installed. Note that the default port is `80` for http and `443` for https but you can change this after the installation.

Each Unity Accelerator hosts [Prometheus metric reports](https://prometheus.io/) as `http://hostname[:port]/metrics`, which you can query from the local network. For a full list of metrics monitoring, refer to [Unity Accelerator Prometheus metrics reference](accelerator-metrics-reference.md). The full configuration of Unity Accelerator is available through its unity-accelerator.cfg file.

## Additional resources

* [Unity Accelerator Prometheus metrics reference](accelerator-metrics-reference.md)
* [Configure Unity Accelerator in the Editor](accelerator-configure.md)