package references

import (
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/instructions/base"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

//hack
func (invokeSpecial *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}


