package control

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (goTo *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, goTo.Offset)
}



