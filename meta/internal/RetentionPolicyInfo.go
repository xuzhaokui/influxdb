// automatically generated, do not modify

package internal

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RetentionPolicyInfo struct {
	_tab flatbuffers.Table
}

func (rcv *RetentionPolicyInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RetentionPolicyInfo) Name() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *RetentionPolicyInfo) Duration() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RetentionPolicyInfo) ShardGroupDuration() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RetentionPolicyInfo) ReplicaN() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RetentionPolicyInfo) ShardGroups(obj *ShardGroupInfo, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(ShardGroupInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *RetentionPolicyInfo) ShardGroupsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func RetentionPolicyInfoStart(builder *flatbuffers.Builder) { builder.StartObject(5) }
func RetentionPolicyInfoAddName(builder *flatbuffers.Builder, Name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Name), 0)
}
func RetentionPolicyInfoAddDuration(builder *flatbuffers.Builder, Duration int64) {
	builder.PrependInt64Slot(1, Duration, 0)
}
func RetentionPolicyInfoAddShardGroupDuration(builder *flatbuffers.Builder, ShardGroupDuration int64) {
	builder.PrependInt64Slot(2, ShardGroupDuration, 0)
}
func RetentionPolicyInfoAddReplicaN(builder *flatbuffers.Builder, ReplicaN uint32) {
	builder.PrependUint32Slot(3, ReplicaN, 0)
}
func RetentionPolicyInfoAddShardGroups(builder *flatbuffers.Builder, ShardGroups flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(ShardGroups), 0)
}
func RetentionPolicyInfoStartShardGroupsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func RetentionPolicyInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
