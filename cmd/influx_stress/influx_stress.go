package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/influxdb/influxdb/stress"
)

var (
	//database  = flag.String("database", "", "name of database")
	//address   = flag.String("addr", "", "IP address and port of database (e.g., localhost:8086)")

	config     = flag.String("config", "", "The stress test file")
	cpuprofile = flag.String("cpuprofile", "", "File where cpu profile will be written")
)

func main() {
	var c *stress.Config
	var err error

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Println(err)
			return
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *config == "" {
		c, err = stress.BasicStress()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		c, err = stress.DecodeFile(*config)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	w := stress.NewWriter(&c.Write.PointGenerators.Basic, &c.Write.InfluxClients.Basic)
	r := stress.NewReader(&c.Read.QueryGenerators.Basic, &c.Read.QueryClients.Basic)
	s := stress.NewStressTest(&c.Provision.Basic, w, r)

	s.Start(stress.BasicWriteHandler, stress.BasicReadHandler)

	return

}
