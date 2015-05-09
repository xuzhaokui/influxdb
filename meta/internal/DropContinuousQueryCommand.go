// automatically generated, do not modify

package internal

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DropContinuousQueryCommand struct {
	_tab flatbuffers.Table
}

func (rcv *DropContinuousQueryCommand) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DropContinuousQueryCommand) Query() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func DropContinuousQueryCommandStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func DropContinuousQueryCommandAddQuery(builder *flatbuffers.Builder, Query flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Query), 0)
}
func DropContinuousQueryCommandEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
