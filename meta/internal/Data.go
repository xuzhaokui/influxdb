// automatically generated, do not modify

package internal

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Data struct {
	_tab flatbuffers.Table
}

func GetRootAsData(buf []byte, offset flatbuffers.UOffsetT) *Data {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Data{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Data) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Data) Nodes(obj *NodeInfo, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(NodeInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Data) NodesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Data) Databases(obj *DatabaseInfo, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(DatabaseInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Data) DatabasesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Data) Users(obj *UserInfo, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(UserInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Data) UsersLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DataStart(builder *flatbuffers.Builder) { builder.StartObject(3) }
func DataAddNodes(builder *flatbuffers.Builder, Nodes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Nodes), 0)
}
func DataStartNodesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DataAddDatabases(builder *flatbuffers.Builder, Databases flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(Databases), 0)
}
func DataStartDatabasesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DataAddUsers(builder *flatbuffers.Builder, Users flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(Users), 0)
}
func DataStartUsersVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
