package math

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
)

type IDIV struct {
	base.NoOperandsInstruction
}

func (idiv *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0{
		panic("java.lang.ArithmeticException:/ by zero")
	}
	result := v1 /v2
	stack.PushInt(result)
}


type LDIV struct {
	base.NoOperandsInstruction
}

func (ldiv *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0{
		panic("java.lang.ArithmeticException:/ by zero")
	}
	result := v1 /v2
	stack.PushLong(result)
}



type FDIV struct {
	base.NoOperandsInstruction
}

func (fdiv *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()

	result := v1/v2
	stack.PushFloat(result)
}


type DDIV struct {
	base.NoOperandsInstruction
}

func (ddiv *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1/v2
	stack.PushDouble(result)
}

