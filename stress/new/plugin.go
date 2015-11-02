package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
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
	//TagCount    int
	//Tags        []tag
	//Fields      []field
	time time.Time
}

func (b *Basic) Generate() <-chan Point {
	c := make(chan Point, 0)

	go func(c chan Point) {
		defer close(c)

		for i := 0; i < b.PointCount; i++ {
			b.time = time.Now()
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
	Address   string
	Database  string
	Precision string
	//SSL       bool
}

func (c *BasicClient) Batch(ps <-chan Point) {
	var buf bytes.Buffer
	var wg sync.WaitGroup

	results := make(chan Timer, 0)

	counter := NewConcurrencyLimiter(10)

	ctr := 0

	for p := range ps {
		b := p.Line()
		ctr++

		buf.Write(b)
		buf.Write([]byte("\n"))

		if ctr%10000 == 0 && ctr != 0 {
			b := buf.Bytes()

			b = b[0 : len(b)-2]

			wg.Add(1)
			counter.Increment()
			go func(byt []byte) {
				t := NewTimer()

				t.StartTimer()
				c.send(byt)
				t.StopTimer()

				counter.Decrement()

				results <- t
				wg.Done()
			}(b)

			var temp bytes.Buffer
			buf = temp
		}

	}

	wg.Wait()
}
func post(url string, datatype string, data io.Reader) error {

	resp, err := http.Post(url, datatype, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf(string(body))
	}

	return nil
}

func (c *BasicClient) send(b []byte) response {
	instanceURL := fmt.Sprintf("http://%v/write?db=%v&precision=%v", c.Address, c.Database, c.Precision)

	_ = post(instanceURL, "application/x-www-form-urlencoded", bytes.NewBuffer(b))

	r := response{
		Status:   "200",
		Time:     time.Now(),
		Duration: time.Duration(1),
	}

	return r
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
			c <- q.Template
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

	r := response{
		Status:   "Executed",
		Time:     time.Now(),
		Duration: time.Duration(1),
	}
	return r

}

func main() {
	b := &Basic{
		PointCount:  1000,
		SeriesCount: 1000000,
		Measurement: "cpu",
	}

	c := &BasicClient{
		Address:   "localhost:8086",
		Database:  "stress",
		Precision: "n",
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

	s := NewStressTest(w, r)

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		s.Start()
		wg.Done()
	}()
	go func() {
		s.Start()
		wg.Done()
	}()

	wg.Wait()

}
