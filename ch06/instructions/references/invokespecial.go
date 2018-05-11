package references

import (
	"jvmgo/ch06/rtda"
	"jvmgo/ch06/instructions/base"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

//hack
func (invokeSpecial *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}


