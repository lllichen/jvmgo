package math

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
)

type ISUB struct {base.NoOperandsInstruction}

func (isub *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1-v2
	stack.PushInt(result)
}


type LSUB struct {base.NoOperandsInstruction}

func (lsub *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1-v2
	stack.PushLong(result)
}


type FSUB struct {base.NoOperandsInstruction}

func (fsub *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1-v2
	stack.PushFloat(result)
}


type DSUB struct {base.NoOperandsInstruction}

func (dsub *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1-v2
	stack.PushDouble(result)
}

