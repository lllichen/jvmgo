package extended

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type IFNULL struct {base.BranchInstruction}

func (ifnull *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, ifnull.Offset)
	}
}

type IFNONNULL struct {base.BranchInstruction}

func (ifnonnull *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, ifnonnull.Offset)
	}
}


