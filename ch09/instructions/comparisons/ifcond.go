package comparisons

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
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

func (ifne *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame,ifne.Offset)
	}
}


//x<0
type IFLT struct{ base.BranchInstruction}


func (iflt *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame,iflt.Offset)
	}
}
// x<=0
type IFLE struct{ base.BranchInstruction}

func (ifle *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame,ifle.Offset)
	}
}
//x>0
type IFGT struct{ base.BranchInstruction}

func (ifgt *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame,ifgt.Offset)
	}
}

//x >= 0
type IFGE struct{ base.BranchInstruction}


func (ifge *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame,ifge.Offset)
	}
}