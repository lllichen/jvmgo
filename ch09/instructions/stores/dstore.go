package stores

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type DSTORE struct {
	base.Index8Instruction
}


func (dstore *DSTORE) Execute (frame *rtda.Frame){
	_lstore(frame,uint(dstore.Index))
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (dstore_0 *DSTORE_0) Execute (frame *rtda.Frame){
	_lstore(frame,0)
}


type DSTORE_1 struct {
	base.NoOperandsInstruction
}


func (dstore_1 *DSTORE_1) Execute (frame *rtda.Frame){
	_lstore(frame,1)
}


type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (dstore_2 *DSTORE_2) Execute (frame *rtda.Frame){
	_lstore(frame,2)
}


type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (dstore_3 *DSTORE_3) Execute (frame *rtda.Frame){
	_lstore(frame,3)
}


func _dstore(frame *rtda.Frame,index uint){
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index,val)
}
