package math

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
)

type IMUL struct {
	base.NoOperandsInstruction
}

func (imul *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}


type LMUL struct {
	base.NoOperandsInstruction
}

func (lmul *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}


type FMUL struct {
	base.NoOperandsInstruction
}

func (fmul *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}


type DMUL struct {
	base.NoOperandsInstruction
}

func (dmul *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}
