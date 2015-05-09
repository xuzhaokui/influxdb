package meta

import (
	"fmt"
	"os"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/influxdb/influxdb/influxql"
	"github.com/influxdb/influxdb/meta/internal"
)

//go:generate flatc -g meta.fbs
//go:generate flatc -g command.fbs

// Data represents the top level collection of all metadata.
type Data struct {
	Version   uint64 // autoincrementing version
	Nodes     []NodeInfo
	Databases []DatabaseInfo
	Users     []UserInfo
}

// Authenticate returns an authenticated user by username.
func (d *Data) Authenticate(username, password string) (*UserInfo, error) {
	panic("not yet implemented")
}

// Authorize user to execute query against database.
// Database can be "" for queries that do not require a database.
// If u is nil, this means authorization is disabled.
func (d *Data) Authorize(u *UserInfo, q *influxql.Query, database string) error {
	panic("not yet implemented")
}

// Node returns a node by id.
func (d *Data) Node(id uint64) *NodeInfo {
	panic("not yet implemented")
}

// NodeByHost returns a node by host.
func (d *Data) NodeByHost(host string) *NodeInfo {
	panic("not yet implemented")
}

// NodesByID returns a list of nodes by a set of ids.
func (d *Data) NodesByID(ids []uint64) []*NodeInfo {
	panic("not yet implemented")
}

// Database returns a database by name.
func (d *Data) Database(name string) *DatabaseInfo {
	panic("not yet implemented")
}

// User returns a User by name.
func (d *Data) User(name string) *UserInfo {
	panic("not yet implemented")
}

// MarshalBinary encodes data to a binary format.
func (data *Data) MarshalBinary() ([]byte, error) {
	b := flatbuffers.NewBuilder(0)
	nodes := data.marshalNodes(b)

	internal.DataStart(b)
	internal.DataAddNodes(b, nodes)
	internal.DataAddNodes(b, nodes)
	b.Finish(internal.DataEnd(b))

	return b.Bytes[b.Head():], nil
}

func (data *Data) marshalNodes(b *flatbuffers.Builder) flatbuffers.UOffsetT {
	offsets := make([]flatbuffers.UOffsetT, len(data.Nodes))
	for i, n := range data.Nodes {
		internal.NodeInfoStart(b)
		internal.NodeInfoAddID(b, n.ID)
		internal.NodeInfoAddHost(b, b.CreateString(n.Host))
		offsets[i] = internal.NodeInfoEnd(b)
	}

	internal.DataStartNodesVector(b, len(offsets))
	for i := len(offsets) - 1; i >= 0; i-- {
		b.PrependUOffsetT(offsets[i])
	}
	return b.EndVector(len(offsets))
}

// MarshalBinary decodes the object from a binary format.
func (data *Data) UnmarshalBinary(buf []byte) error {
	fbs := internal.GetRootAsData(buf, 0)

	// Construct nodes.
	data.Nodes = make([]NodeInfo, fbs.NodesLength())
	for i := range data.Nodes {
		var other internal.NodeInfo
		fbs.Nodes(&other, i)

		n := &data.Nodes[i]
		n.ID = other.ID()
		n.Host = other.Host()
	}

	return nil
}

// NodeInfo represents information about a single node in the cluster.
type NodeInfo struct {
	ID   uint64
	Host string
}

// DatabaseInfo represents information about a database in the system.
type DatabaseInfo struct {
	Name                   string
	DefaultRetentionPolicy string
	Polices                []RetentionPolicyInfo
	ContinuousQueries      []ContinuousQueryInfo
}

// RetentionPolicy returns a policy on the database by name.
func (db *DatabaseInfo) RetentionPolicy(name string) *RetentionPolicyInfo {
	panic("not yet implemented")
}

// RetentionPolicyInfo represents metadata about a retention policy.
type RetentionPolicyInfo struct {
	Name               string
	ReplicaN           int
	Duration           time.Duration
	ShardGroupDuration time.Duration
	ShardGroups        []ShardGroupInfo
}

// ShardGroupInfo represents metadata about a shard group.
type ShardGroupInfo struct {
	ID        uint64
	StartTime time.Time
	EndTime   time.Time
	Shards    []ShardInfo
}

// ShardInfo represents metadata about a shard.
type ShardInfo struct {
	ID       uint64
	OwnerIDs []uint64
}

// ContinuousQueryInfo represents metadata about a continuous query.
type ContinuousQueryInfo struct {
	Query string
}

// UserInfo represents metadata about a user in the system.
type UserInfo struct {
	Name       string
	Hash       string
	Admin      bool
	Privileges map[string]influxql.Privilege
}

func warn(v ...interface{})              { fmt.Fprintln(os.Stderr, v...) }
func warnf(msg string, v ...interface{}) { fmt.Fprintf(os.Stderr, msg+"\n", v...) }
