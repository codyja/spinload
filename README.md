# SpinLoad
##NOTE: Under Active Development
SpinLoad is a simple web server for generating desired load types on demand. Can be used for testing infrastructure, monitoring applications and 
alerts, and autoscaling configurations. You can send a `POST` to spinload and request 50% CPU usage for a period of
time by using query parameters. 

### Use Cases

1. Leverage hey to send a series of requests for 5 minutes, each requesting 100% CPU usage to test autoscaling at the
   application or auto-scale group level.
1.

### Examples

1. Request 25% cpu usage for 20 seconds at a time:
```shell
curl http://server:1986/load?cpu=25&time=20s
```

### Features
* CPU - Target CPU percent usage to consume. Leverages a simple infinite loop to consume CPU time and GO channels to
  stay around desired CPU percent. 
* Memory - A target memory usage to consume (WIP)
* DiskIO - Desired Disk IOPs to use (WIP)
* Queue Size - Prometheus formatted metric showing an artificial queue size (WIP)

