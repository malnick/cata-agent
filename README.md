# Cata-Agent
Container Data. Agent used by Cata service.

## Overview
Cata is a framework consiting of "agents" and a "console". The cata-agent resides on your host machines for your docker infrastructure runs. The cata-agent then exposes the system on each host (memory, storage, CPU usage as well as Docker inspect data) and exposes a REST API to the Cata-console to ingest. 

Cata-agents can be configured with host-level alarms using environment variables. These alarms allow quick configuration of basic monitoring requirements. 

## ENV Config

```CATA_CONSOLES=```: []string 
  
  An array of consoles to post data. Example: ```CATA_CONSOLE=my.console1,my.console2```

  Defualt is ```localhost```.

```CATA_POST_SPLAY=```: int

  An integer representing POST splay in minutes. Default is 5.

## Alarms - NOT FULLY IMPLEMENTED
Alarms can be set with env variables, executed with the docker daemon as such:

```
docker run -d -e 'CATA_ALARM_MEMORY=90,70,50' -e 'CATA_ALARM_CPU=80,60,40' yourorg/cata_agent:latest
```

The cata-agent reads in the list of values from the alarm, setting basic 'critical', 'warning' and 'ok' requirements for each alarm. 

#### Available Alarms
MEMORY
CPU
STORAGE

## REST Interface
The cata-agent exposes a basic REST interface for your host. 

```/```: The basic configuration for the cata-agent

#### Host Index
```/host```: Host data

```/host/cpu```: CPU data

```/host/load```: Load data

```/host/memory```: Host memory

```/host/storage```: Host storage

#### Container Index
```/containers```: Container data
