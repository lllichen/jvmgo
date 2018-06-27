package loads

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
)

type FLOAD struct { base.Index8Instruction}

func (fload *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame,uint(fload.Index))
}

type FLOAD_0 struct { base.NoOperandsInstruction}

func (fload_0 *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct { base.NoOperandsInstruction}

func (fload_1 *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct { base.NoOperandsInstruction}

func (fload_2 *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct { base.NoOperandsInstruction}

func (fload_3 *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}

func _fload(frame *rtda.Frame, index uint){
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

