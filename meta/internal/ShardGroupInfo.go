// automatically generated, do not modify

package internal

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ShardGroupInfo struct {
	_tab flatbuffers.Table
}

func (rcv *ShardGroupInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ShardGroupInfo) ID() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ShardGroupInfo) StartTime() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ShardGroupInfo) EndTime() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ShardGroupInfo) Shards(obj *ShardInfo, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(ShardInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ShardGroupInfo) ShardsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ShardGroupInfoStart(builder *flatbuffers.Builder)            { builder.StartObject(4) }
func ShardGroupInfoAddID(builder *flatbuffers.Builder, ID uint64) { builder.PrependUint64Slot(0, ID, 0) }
func ShardGroupInfoAddStartTime(builder *flatbuffers.Builder, StartTime int64) {
	builder.PrependInt64Slot(1, StartTime, 0)
}
func ShardGroupInfoAddEndTime(builder *flatbuffers.Builder, EndTime int64) {
	builder.PrependInt64Slot(2, EndTime, 0)
}
func ShardGroupInfoAddShards(builder *flatbuffers.Builder, Shards flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(Shards), 0)
}
func ShardGroupInfoStartShardsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ShardGroupInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
