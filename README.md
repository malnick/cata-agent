# Cata-Agent
Container Data. Agent used by Cata service.

## Overview
Cata is a web service that listens for POSTs from distributed cata-agents running as containers on each docker host in your infrastructure. Cata agents return data about each host: 

#### Memory Usage
1. Memory available
1. Memory being used by containers by container-name/hash

#### CPU Usage
1. CPU available on each host
1. CPU being used by containers 

#### Storage
1. Storage available on host
1. Storage being used by container (aufs, diff, etc); by directory so you can quickly get insights to where the data is filling up.

#### Container Data
1. Volumes mounts
1. Docker Inspect Output (/container/$id/inspect)


