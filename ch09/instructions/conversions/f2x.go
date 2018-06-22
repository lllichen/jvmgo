package conversions

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type F2D struct {
	base.NoOperandsInstruction
}

func (f2d *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

type F2I struct {
	base.NoOperandsInstruction
}


func (*F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

type F2L struct {
	base.NoOperandsInstruction
}


func (*F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := int64(f)
	stack.PushLong(d)
}