// automatically generated, do not modify

package internal

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ShardInfo struct {
	_tab flatbuffers.Table
}

func (rcv *ShardInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ShardInfo) ID() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ShardInfo) OwnerIDs() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func ShardInfoStart(builder *flatbuffers.Builder)            { builder.StartObject(2) }
func ShardInfoAddID(builder *flatbuffers.Builder, ID uint64) { builder.PrependUint64Slot(0, ID, 0) }
func ShardInfoAddOwnerIDs(builder *flatbuffers.Builder, OwnerIDs uint64) {
	builder.PrependUint64Slot(1, OwnerIDs, 0)
}
func ShardInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
