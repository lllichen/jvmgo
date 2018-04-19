package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
	"math"
)

type DREM struct {
	base.NoOperandsInstruction
}

func (drem *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1,v2)
	stack.PushDouble(result)
}

type FREM struct {
	base.NoOperandsInstruction
}


func (frem *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(result)
}

type IREM struct {
	base.NoOperandsInstruction
}

func (irem *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException:/ by zero")
	}
	result := v1%v2
	stack.PushInt(result)
}

type LREM struct {
	base.NoOperandsInstruction
}

func (lrem *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException:/ by zero")
	}
	result := v1%v2
	stack.PushLong(result)
}
