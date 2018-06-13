package loads

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type LLOAD struct { base.Index8Instruction}

func (lload *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame,uint(lload.Index))
}

type LLOAD_0 struct { base.NoOperandsInstruction}

func (lload_0 *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct { base.NoOperandsInstruction}

func (lload_1 *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct { base.NoOperandsInstruction}

func (lload_2 *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct { base.NoOperandsInstruction}

func (lload_3 *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}

func _lload(frame *rtda.Frame, index uint){
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

