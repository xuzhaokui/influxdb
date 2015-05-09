package meta_test

import (
	"reflect"
	"testing"

	"github.com/influxdb/influxdb/meta"
)

// Ensure the data can be encoded to and from a binary format.
func TestData_Marshal(t *testing.T) {
	// Encode object.
	data := meta.Data{
		Nodes: []meta.NodeInfo{
			{ID: 100, Host: "server0"},
			{ID: 101, Host: "server1"},
		},
	}
	buf, err := data.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	// Decode object.
	var other meta.Data
	if err := other.UnmarshalBinary(buf); err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(data, other) {
		t.Fatalf("mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", data, other)
	}
}
