package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type LSTORE struct {
	base.Index8Instruction
}


type LSTORE_0 struct {
	base.NoOperandsInstruction
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func _lstore(frame *rtda.Frame,index uint){
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index,val)
}

func (lstore *LSTORE) Execute (frame *rtda.Frame){
	_lstore(frame,2)
}