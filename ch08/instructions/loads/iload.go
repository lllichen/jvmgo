package loads

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type ILOAD struct { base.Index8Instruction}

func (iload *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame,uint(iload.Index))
}

type ILOAD_0 struct { base.NoOperandsInstruction}

func (iload_0 *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct { base.NoOperandsInstruction}

func (iload_1 *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct { base.NoOperandsInstruction}

func (iload_2 *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct { base.NoOperandsInstruction}

func (iload_3 *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint){
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

