# spinload
Simple web server to generate load on demand. Can be used for testing infrastructure, monitoring workflows, and autoscaling configurations.

### Features
* CPU - Target CPU percent usage to consume. Leverages a simple infinite loop to consume CPU time and GO channels to
  stay around desired CPU percent. 
* Memory - A target memory usage to consume (WIP)
* DiskIO - Desired Disk IOPs to use (WIP)
* Queue Size - Prometheus formatted metric showing an artificial queue size (WIP)

