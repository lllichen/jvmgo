package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type IF_ICMPEQ struct {base.BranchInstruction}

func (ifIcmpeq *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 == v2 {
		base.Branch(frame,ifIcmpeq.Offset)
	}
}

type IF_ICMPNE struct {base.BranchInstruction}

func (ifIcmpne *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 != v2 {
		base.Branch(frame,ifIcmpne.Offset)
	}
}

type IF_ICMPLT struct {base.BranchInstruction}


func (ifIcmplT *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 < v2 {
		base.Branch(frame,ifIcmplT.Offset)
	}
}

type IF_ICMPLE struct {base.BranchInstruction}

func (ifIcmple *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 <= v2 {
		base.Branch(frame,ifIcmple.Offset)
	}
}


type IF_ICMPGT struct {base.BranchInstruction}

func (ifIcmpgt *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 > v2 {
		base.Branch(frame,ifIcmpgt.Offset)
	}
}
type IF_ICMPGE struct {base.BranchInstruction}

func (ifIcmpge *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 > v2 {
		base.Branch(frame,ifIcmpge.Offset)
	}
}

