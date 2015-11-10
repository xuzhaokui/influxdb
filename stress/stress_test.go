package stress

import (
	//"encoding/json"
	//"net/http"
	//"net/http/httptest"
	"testing"
	"time"

	//"github.com/influxdb/influxdb/client"
)

func TestTimer_StartTimer(t *testing.T) {
	var epoch time.Time

	tmr := &Timer{}

	tmr.StartTimer()

	s := tmr.Start()

	if s == epoch {
		t.Errorf("expected tmr.start to not be %v", s)
	}
}

func TestNewTimer(t *testing.T) {
	var epoch time.Time

	tmr := NewTimer()

	s := tmr.Start()

	if s == epoch {
		t.Errorf("expected tmr.start to not be %v", s)
	}

	e := tmr.End()

	if e != epoch {
		t.Errorf("expected tmr.stop to be %v, got %v", epoch, e)
	}
}

func TestTimer_StopTimer(t *testing.T) {
	var epoch time.Time

	tmr := NewTimer()

	tmr.StopTimer()

	e := tmr.End()

	if e == epoch {
		t.Errorf("expected tmr.stop to not be %v", e)
	}
}

func TestTimer_Elapsed(t *testing.T) {

	tmr := NewTimer()
	time.Sleep(2 * time.Second)
	tmr.StopTimer()

	e := tmr.Elapsed()

	if time.Duration(2*time.Second) > e || e > time.Duration(3*time.Second) {
		t.Errorf("expected around %s got %s", time.Duration(2*time.Second), e)
	}

}

//func TestNewResponseTime(t *testing.T) {
//	r := NewResponseTime(100)
//
//	if r.Value != 100 {
//		t.Errorf("expected Value to be %v, got %v", 100, r.Value)
//	}
//
//	var epoch time.Time
//
//	if r.Time == epoch {
//		t.Errorf("expected r.Time not to be %v", epoch)
//	}
//
//}

//func TestRun(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		var data client.Response
//		w.WriteHeader(http.StatusOK)
//		_ = json.NewEncoder(w).Encode(data)
//	}))
//	defer ts.Close()
//
//	ms := make(Measurements, 0)
//	ms.Set("this,is,a,test")
//
//	url := ts.URL[7:]
//
//	cfg, _ := DecodeFile("test.toml")
//
//	cfg.Write.Address = url
//
//	d := make(chan struct{})
//	timestamp := make(chan time.Time)
//
//	tp, _, rts, tmr := Run(cfg, d, timestamp)
//
//	ps := cfg.Series[0].SeriesCount * cfg.Series[0].PointCount
//
//	if tp != ps {
//		t.Fatalf("unexpected error. expected %v, actual %v", ps, tp)
//	}
//
//	if len(rts) != ps/cfg.Write.BatchSize {
//		t.Fatalf("unexpected error. expected %v, actual %v", ps/cfg.Write.BatchSize, len(rts))
//	}
//
//	var epoch time.Time
//
//	if tmr.Start() == epoch {
//		t.Errorf("expected trm.start not to be %s", epoch)
//	}
//
//	if tmr.End() == epoch {
//		t.Errorf("expected trm.end not to be %s", epoch)
//	}
//}
