# System Metrics

System Metrics microservice implemented on Go.

## Instructions

To start this app:

- Installing Dependencies

```
make deps
```

- Running Application:

```
make run
```

- Run Test Application:

```
make test
```

## Basic Features

- On endpoint "/api/system" expose information on json format about hostname, os, uptime, cpu, disk, memory and network interfaces.
