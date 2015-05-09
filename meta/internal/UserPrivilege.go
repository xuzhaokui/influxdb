// automatically generated, do not modify

package internal

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type UserPrivilege struct {
	_tab flatbuffers.Table
}

func (rcv *UserPrivilege) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UserPrivilege) Database() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *UserPrivilege) Privilege() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func UserPrivilegeStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func UserPrivilegeAddDatabase(builder *flatbuffers.Builder, Database flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Database), 0)
}
func UserPrivilegeAddPrivilege(builder *flatbuffers.Builder, Privilege int32) {
	builder.PrependInt32Slot(1, Privilege, 0)
}
func UserPrivilegeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
