package conversions

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
)

type I2L struct {
	base.NoOperandsInstruction
}

func (i2l *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}


type I2F struct {
	base.NoOperandsInstruction
}

func (i2f *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}


type I2D struct {
	base.NoOperandsInstruction
}

func (i2d *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}


type I2B struct {
	base.NoOperandsInstruction
}

func (i2b *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(int8(i))
	stack.PushInt(b)
}


type I2C struct {
	base.NoOperandsInstruction
}


func (i2c *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(int16(i))
	stack.PushInt(b)
}

type I2S struct {
	base.NoOperandsInstruction
}


func (i2s *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(uint16(i))
	stack.PushInt(b)
}