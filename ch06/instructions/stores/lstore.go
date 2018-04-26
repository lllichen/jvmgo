package stores

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type LSTORE struct {
	base.Index8Instruction
}


func (lstore *LSTORE) Execute (frame *rtda.Frame){
	_lstore(frame,uint(lstore.Index))
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (*LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame,0)
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (*LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame,1)
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (*LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame,2)
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (*LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame,3)
}

func _lstore(frame *rtda.Frame,index uint){
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index,val)
}
