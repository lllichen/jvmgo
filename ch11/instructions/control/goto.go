package control

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (goTo *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, goTo.Offset)
}



