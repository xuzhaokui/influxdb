// automatically generated, do not modify

package internal

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type User struct {
	_tab flatbuffers.Table
}

func (rcv *User) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *User) Name() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *User) Hash() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *User) Admin() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *User) Privileges(obj *UserPrivilege, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(UserPrivilege)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *User) PrivilegesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func UserStart(builder *flatbuffers.Builder) { builder.StartObject(4) }
func UserAddName(builder *flatbuffers.Builder, Name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Name), 0)
}
func UserAddHash(builder *flatbuffers.Builder, Hash flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(Hash), 0)
}
func UserAddAdmin(builder *flatbuffers.Builder, Admin byte) { builder.PrependByteSlot(2, Admin, 0) }
func UserAddPrivileges(builder *flatbuffers.Builder, Privileges flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(Privileges), 0)
}
func UserStartPrivilegesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func UserEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
