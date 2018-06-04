package math

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type INEG struct {
	base.NoOperandsInstruction
}

func (ineg *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := -val
	stack.PushInt(result)
}


type LNEG struct {
	base.NoOperandsInstruction
}

func (lneg *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	result := -val
	stack.PushLong(result)
}

type FNEG struct {
	base.NoOperandsInstruction
}

func (fneg *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	result := -val
	stack.PushFloat(result)
}


type DNEG struct {
	base.NoOperandsInstruction
}

func (dneg *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	result := -val
	stack.PushDouble(result)
}