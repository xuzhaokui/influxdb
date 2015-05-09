// automatically generated, do not modify

package internal

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Command struct {
	_tab flatbuffers.Table
}

func GetRootAsCommand(buf []byte, offset flatbuffers.UOffsetT) *Command {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Command{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Command) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Command) ImplType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Command) Impl(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func CommandStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func CommandAddImplType(builder *flatbuffers.Builder, ImplType byte) {
	builder.PrependByteSlot(0, ImplType, 0)
}
func CommandAddImpl(builder *flatbuffers.Builder, Impl flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(Impl), 0)
}
func CommandEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
