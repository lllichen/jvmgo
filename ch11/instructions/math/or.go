package math

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
)

type IOR struct {
	base.NoOperandsInstruction
}

func (ior *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)

}

type LOR struct {
	base.NoOperandsInstruction
}

func (lor *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}


