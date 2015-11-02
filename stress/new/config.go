package main

import (
	"bytes"
	"fmt"
	"time"
)

type tag struct {
	Key   string `toml:"key"`
	Value string `toml:"value"`
}

// tag is a struct that contains data
// about a field for in a series
type field struct {
	Key  string `toml:"key"`
	Type string `toml:"type"`
}

type series struct {
	PointCount  int     `toml:"point_count"`
	Tick        string  `toml:"tick"`
	Jitter      bool    `toml:"jitter"`
	Measurement string  `toml:"measurement"`
	SeriesCount int     `toml:"series_count"`
	TagCount    int     `toml:"tag_count"`
	Tags        []tag   `toml:"tag"`
	Fields      []field `toml:"field"`
}

////////////////////////
/// WRITE CONFIG //////
//////////////////////

type WriteConfig struct {
	Enabled       bool   `toml:"enabled"`
	BatchInterval string `toml:"batch_interval"`
	BatchSize     int    `toml:"batch_size"`
	Concurrency   int    `toml:"concurrency"`
	Precision     string `toml:"precision"`
	StartDate     string `toml:"start_date"`
	Output        Output `toml:"output"`
	Series        Series `toml:"series"`
}

type Output struct {
	Type string `toml:"type"`
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type Series struct {
	Standard PointGenerator `toml:"point_generator"`
}
