# Unity Accelerator Prometheus metrics reference

Each Unity **Accelerator** hosts [Prometheus metric reports](https://prometheus.io/) as `/metrics`, which you can query from the local network. The following is a full list of metrics that you can access:

| **Metric** | **Description** |
| --- | --- |
| `process_resident_memory_bytes` | The amount of memory used by the cache server. |
| `uta_agent_sys_cpu_percent` | The amount of CPU used by the system. |
| `uta_agent_sys_mem_bytes_used` | The amount of memory used by the system. |
| `uta_agent_protobuf_connects` | The number of times Unity Editors have connected to the service for the Asset Import Pipeline. |
| `uta_agent_protobuf_bytes_in` | The number of bytes received from the Asset Import Pipeline. |
| `uta_agent_protobuf_bytes_out` | The number of bytes sent to the Asset Import Pipeline. |
| `uta_agent_protobuf_requests{code="0",method="get"}` | The number of items sent to the Asset Import Pipeline. |
| `uta_agent_protobuf_requests{code="1",method="get"}` | The number of errors occurring from items sent to the Asset Import Pipeline. |
| `uta_agent_protobuf_requests{code="2",method="get"}` | The number of namespace errors occurring from items sent to the Asset Import Pipeline. |
| `uta_agent_protobuf_requests{code="3",method="get"}` | The number of requests made from the Asset Import Pipeline for items that have not been cached. |
| `uta_agent_protobuf_requests{code="0",method="put"}` | The number of items received from the Asset Import Pipeline. |
| `uta_agent_protobuf_requests{code="1",method="put"}` | The number of errors occurring from items received from the Asset Import Pipeline. |
| `uta_agent_protobuf_requests{code="2",method="put"}` | The number of namespace errors occurring from items received from the Asset Import Pipeline. |
| `uta_agent_protobuf_requests{code="0",method="delete"}` | The number of requests made from Asset Import Pipeline to remove items. |
| `uta_agent_protobuf_requests{code="1",method="delete"}` | The number of errors occurring from the Asset Import Pipeline removing items. |
| `uta_agent_protobuf_requests{code="2",method="delete"}` | The number of namespace errors occurring from requests to remove items by the Asset Import Pipeline. |
| `uta_agent_protobuf_requests{code="3",method="delete"}` | The number of requests to remove items made from the Asset Import Pipeline for items that haven’t been cached. |
| `uta_agent_cache_bytes_evicted_unscheduled` | Number of bytes evicted from cache during an unscheduled pass. Unscheduled evictions occur when running out of disk space while processing a request. Normally, scheduled background evictions should keep up with demand, so an increase in this metric’s value indicates an overwhelmed server. Increasing the configuration’s `CacheMinFreeBytes` or `CacheMinFreePercent` might help give more free space, or you might need to upgrade the underlying hardware or install an additional Unity Accelerator instance. |
| `uta_agent_cache_files_evicted_unscheduled` | Similar to the previous `uta_agent_cache_bytes_evicted_unscheduled` metrics, but counts files evicted instead of byte count. This might be an easier metric to alert on. For example: Alert when files metric increases by more than 1 per second. |
| `uta_agent_protobuf_cache_hits` | Number of Asset Import Pipeline requests that were served from cache. |
| `uta_agent_protobuf_cache_misses` | Number of Asset Import Pipeline requests that couldn’t be served from cache. |
| `uta_agent_protobuf_cache_bytes_out` | Number of bytes served from cache to the Asset Import Pipeline. |
| `uta_agent_protobuf_cache_bytes_in` | Number of bytes stored into the cache received from the Asset Import Pipeline. |

## Additional resources

* [Monitor Unity Accelerator](accelerator-monitor.md)
* [Use Unity Accelerator on the command line](accelerator-command-line.md)