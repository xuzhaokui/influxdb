// automatically generated, do not modify

package internal

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NodeInfo struct {
	_tab flatbuffers.Table
}

func (rcv *NodeInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NodeInfo) ID() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NodeInfo) Host() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func NodeInfoStart(builder *flatbuffers.Builder)            { builder.StartObject(2) }
func NodeInfoAddID(builder *flatbuffers.Builder, ID uint64) { builder.PrependUint64Slot(0, ID, 0) }
func NodeInfoAddHost(builder *flatbuffers.Builder, Host flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(Host), 0)
}
func NodeInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
