package main

import ()

type Config struct {
	Provision Provision `toml:"provision"`
	Write     Write     `toml:"write"`
	Read      Read      `toml:"read"`
}

type Provision struct {
	Basic BasicProvisioner `toml:"basic"`
}

type Write struct {
	PointGenerators PointGenerators `toml:"point_generator"`
	InfluxClients   InfluxClients   `toml:"influx_client"`
}

type PointGenerators struct {
	Basic BasicPointGenerator `toml:"basic"`
}

type InfluxClients struct {
	Basic BasicClient `toml:"basic"`
}

type Read struct {
	QueryGenerators QueryGenerators `toml:"query_generator"`
	QueryClients    QueryClients    `toml:"query_client"`
}

type QueryGenerators struct {
	Basic BasicQuery `toml:"basic"`
}

type QueryClients struct {
	Basic BasicQueryClient `toml:"basic"`
}
