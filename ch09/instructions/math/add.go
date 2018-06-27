package math

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type IADD struct {base.NoOperandsInstruction}

func (iadd *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

type LADD struct {base.NoOperandsInstruction}


func (ladd *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}
type FADD struct {base.NoOperandsInstruction}

func (fadd *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}
type DADD struct {base.NoOperandsInstruction}


func (dadd *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}