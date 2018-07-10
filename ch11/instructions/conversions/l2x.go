package conversions

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
)

type L2I struct {
	base.NoOperandsInstruction
}

func (l2i *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}



type L2F struct {
	base.NoOperandsInstruction
}

func (l2f *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}


type L2D struct {
	base.NoOperandsInstruction
}

func (l2d *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}