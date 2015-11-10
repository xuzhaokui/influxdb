# Stress Test

## Provisioner
Provisions the InfluxDB instance where the stress test is going to be ran against.

Think things like, creating the database, setting up retention policies, continuous queries, etc.

## Writer
The Writer is responsible for Writing data into an InfluxDB instance. It has two components: PointGenerator and InfluxClient.

### PointGenerator
The PointGenerator is responsible for generating points that will be written into InfluxDB. Additionally, it is reponsible for keeping track of the latest timestamp of the points it is writing (Just incase the its needed by the Reader).

Any type that implements the methods `Generate()` and `Time()` is a PointGenerator.

### InfluxClient


## Reader
### QueryGenerator
### QueryClient

