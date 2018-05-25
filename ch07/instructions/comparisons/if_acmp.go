package comparisons

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type IF_ACMPEQ struct { base.BranchInstruction}

func (ifAcmpeq *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2{
		base.Branch(frame,ifAcmpeq.Offset)
	}
}

type IF_ACMPNE struct { base.BranchInstruction}

func (ifAcmpne *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 != ref2{
		base.Branch(frame,ifAcmpne.Offset)
	}
}


