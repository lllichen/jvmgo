package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type IF_ICMPEQ struct {base.BranchInstruction}
type IF_ICMPNE struct {base.BranchInstruction}

func (if_icmpne *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 != v2 {
		base.Branch(frame,if_icmpne.Offset)
	}
}

type IF_ICMPLT struct {base.BranchInstruction}
type IF_ICMPLE struct {base.BranchInstruction}
type IF_ICMPGT struct {base.BranchInstruction}
type IF_ICMPGE struct {base.BranchInstruction}


