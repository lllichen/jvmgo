package stack

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}


type POP2 struct {
	base.NoOperandsInstruction
}

func (pop *POP) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (pop *POP2) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}



