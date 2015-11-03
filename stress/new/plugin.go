package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"os"
	"runtime/pprof"
	//"github.com/influxdb/strees/provisioner"

	"github.com/influxdb/influxdb/client/v2"
)

///////////////////////////////////////////
type Basic struct {
	PointCount  int
	Tick        string
	Jitter      bool
	Measurement string
	SeriesCount int
	//Tags        []tag
	//Fields      []field
	StartDate string
	time      time.Time
}

func (b *Basic) Generate() <-chan Point {
	c := make(chan Point, 0)

	go func(c chan Point) {
		defer close(c)

		start, err := time.Parse("2006-Jan-02", b.StartDate)
		if err != nil {
			fmt.Println(err)
		}

		b.time = start

		tick, err := time.ParseDuration(b.Tick)
		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < b.PointCount; i++ {
			b.time = b.time.Add(tick)
			for j := 0; j < b.SeriesCount; j++ {
				p := Point{
					Measurement: b.Measurement,
					Tags:        append(make(Tags, 0), Tag{Key: "host", Value: fmt.Sprintf("server-%v", j)}),
					Fields:      append(make(Fields, 0), Field{Key: "value", Value: fmt.Sprintf("%v", rand.Intn(100))}),
					Timestamp:   b.time.UnixNano(),
				}

				c <- p
			}
		}
	}(c)

	return c
}

func (b *Basic) Time() time.Time {
	return b.time
}

type BasicClient struct {
	Address     string
	Database    string
	Precision   string
	BatchSize   int
	Concurrency int
	SSL         bool
}

func (c *BasicClient) Batch(ps <-chan Point, r chan<- response) {
	var buf bytes.Buffer
	var wg sync.WaitGroup

	counter := NewConcurrencyLimiter(c.Concurrency)

	ctr := 0

	for p := range ps {
		b := p.Line()
		ctr++

		buf.Write(b)
		buf.Write([]byte("\n"))

		if ctr%c.BatchSize == 0 && ctr != 0 {
			b := buf.Bytes()

			b = b[0 : len(b)-2]

			wg.Add(1)
			counter.Increment()
			go func(byt []byte) {

				rs := c.send(byt)

				counter.Decrement()
				r <- rs
				wg.Done()
			}(b)

			var temp bytes.Buffer
			buf = temp
		}

	}

	wg.Wait()
}

func post(url string, datatype string, data io.Reader) (*http.Response, error) {

	resp, err := http.Post(url, datatype, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	return resp, nil
}

func (c *BasicClient) send(b []byte) response {
	instanceURL := fmt.Sprintf("http://%v/write?db=%v&precision=%v", c.Address, c.Database, c.Precision)
	t := NewTimer()

	t.StartTimer()
	resp, err := post(instanceURL, "application/x-www-form-urlencoded", bytes.NewBuffer(b))
	t.StopTimer()
	if err != nil {
		fmt.Println(err)
	}

	r := response{
		Resp:  resp,
		Time:  time.Now(),
		Timer: t,
	}

	return r
}

func (c *BasicClient) Handle(resp <-chan response, fn func(r response)) {
	for rs := range resp {
		fn(rs)
	}
}

//////////////

type BasicQuery struct {
	Template Query
	time     time.Time
}

func (q *BasicQuery) QueryGenerate() <-chan Query {
	c := make(chan Query, 0)

	go func(chan Query) {
		defer close(c)

		for i := 0; i < 1000; i++ {
			time.Sleep(10 * time.Millisecond)
			//c <- q.Template
		}

	}(c)

	return c
}

func (q *BasicQuery) SetTime(t time.Time) {
	q.time = t

	return
}

type BasicQueryClient struct {
	Address  string
	Database string
	client   client.Client
}

func (b *BasicQueryClient) Init() {
	u, _ := url.Parse(fmt.Sprintf("http://%v", b.Address))
	cl := client.NewClient(client.Config{
		URL: u,
	})

	b.client = cl
}

func (b *BasicQueryClient) Query(cmd Query) response {
	q := client.Query{
		Command:  string(cmd),
		Database: b.Database,
	}
	_, _ = b.client.Query(q)

	r := response{}

	return r

}

///////////////////

func resetDB(c client.Client, database string) error {
	_, err := c.Query(client.Query{
		Command: fmt.Sprintf("DROP DATABASE %s", database),
	})

	if err != nil && !strings.Contains(err.Error(), "database not found") {
		return err
	}

	_, err = c.Query(client.Query{
		Command: fmt.Sprintf("CREATE DATABASE %s", database),
	})

	return nil
}

type BasicProvisioner struct {
	Address       string
	Database      string
	ResetDatabase bool
}

func (b *BasicProvisioner) Provision() {
	u, _ := url.Parse(fmt.Sprintf("http://%v", b.Address))
	cl := client.NewClient(client.Config{
		URL: u,
	})

	if b.ResetDatabase {
		resetDB(cl, b.Database)
	}
}

func main() {
	f, err := os.Create("cpuprofile")
	if err != nil {
		fmt.Println(err)
		return
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	b := &Basic{
		PointCount:  100,
		SeriesCount: 100000,
		Measurement: "cpu",
		StartDate:   "2006-Jan-02",
		Tick:        "1s",
	}

	c := &BasicClient{
		Address:     "localhost:8086",
		Database:    "stress",
		Precision:   "n",
		BatchSize:   10000,
		Concurrency: 10,
	}

	w := NewWriter(b, c)

	qg := &BasicQuery{
		Template: Query("SELECT * FROM cpu"),
	}

	qc := &BasicQueryClient{
		Address:  "localhost:8086",
		Database: "stress",
	}

	qc.Init()

	r := NewReader(qg, qc)

	bp := &BasicProvisioner{
		Address:       "localhost:8086",
		Database:      "stress",
		ResetDatabase: true,
	}

	s := NewStressTest(bp, w, r)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		s.Start()
		wg.Done()

		f, err := os.Create("memprofile")
		if err != nil {
			fmt.Println(err)
			return
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}()

	wg.Wait()

}
