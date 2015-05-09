// automatically generated, do not modify

package internal

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ContinuousQueryInfo struct {
	_tab flatbuffers.Table
}

func (rcv *ContinuousQueryInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ContinuousQueryInfo) Query() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func ContinuousQueryInfoStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func ContinuousQueryInfoAddQuery(builder *flatbuffers.Builder, Query flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Query), 0)
}
func ContinuousQueryInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
