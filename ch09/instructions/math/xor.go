package math

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type IXOR struct {
	base.NoOperandsInstruction
}

func (ixor *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

type LXOR struct {
	base.NoOperandsInstruction
}

func (lxor *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)

}
