// automatically generated, do not modify

package internal

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DatabaseInfo struct {
	_tab flatbuffers.Table
}

func (rcv *DatabaseInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DatabaseInfo) Name() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *DatabaseInfo) DefaultRetentionPolicy() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *DatabaseInfo) Policies(obj *RetentionPolicyInfo, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(RetentionPolicyInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DatabaseInfo) PoliciesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DatabaseInfo) ContinuousQueries(obj *ContinuousQueryInfo, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		if obj == nil {
			obj = new(ContinuousQueryInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DatabaseInfo) ContinuousQueriesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DatabaseInfoStart(builder *flatbuffers.Builder) { builder.StartObject(4) }
func DatabaseInfoAddName(builder *flatbuffers.Builder, Name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Name), 0)
}
func DatabaseInfoAddDefaultRetentionPolicy(builder *flatbuffers.Builder, DefaultRetentionPolicy flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(DefaultRetentionPolicy), 0)
}
func DatabaseInfoAddPolicies(builder *flatbuffers.Builder, Policies flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(Policies), 0)
}
func DatabaseInfoStartPoliciesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DatabaseInfoAddContinuousQueries(builder *flatbuffers.Builder, ContinuousQueries flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(ContinuousQueries), 0)
}
func DatabaseInfoStartContinuousQueriesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DatabaseInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
