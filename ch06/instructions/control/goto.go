package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (goTo *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, goTo.Offset)
}



