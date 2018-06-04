package loads

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type ALOAD struct { base.Index8Instruction}

func (aload *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame,uint(aload.Index))
}

type ALOAD_0 struct { base.NoOperandsInstruction}

func (aload_0 *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct { base.NoOperandsInstruction}

func (aload_1 *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct { base.NoOperandsInstruction}

func (aload_2 *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct { base.NoOperandsInstruction}

func (aload_3 *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rtda.Frame, index uint){
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

