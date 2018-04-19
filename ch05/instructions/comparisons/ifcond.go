package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//x == 0
type IFEQ struct{ base.BranchInstruction}

func (ifeq *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame,ifeq.Offset)
	}

}

//x!=0
type IFNE struct{ base.BranchInstruction}
//x<0
type IFLT struct{ base.BranchInstruction}
// x<=0
type IFLE struct{ base.BranchInstruction}
//x>0
type IFGT struct{ base.BranchInstruction}
//x >= 0
type IFGE struct{ base.BranchInstruction}

