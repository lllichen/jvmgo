package stores

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type ASTORE struct {
	base.Index8Instruction
}

func (astore *ASTORE) Execute (frame *rtda.Frame){
	_astore(frame,uint(astore.Index))
}

type ASTORE_0 struct {
	base.NoOperandsInstruction
}


func (astore_0 *ASTORE_0) Execute (frame *rtda.Frame){
	_astore(frame,0)
}


type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (astore_1 *ASTORE_1) Execute (frame *rtda.Frame){
	_astore(frame,1)
}

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (astore_2 *ASTORE_2) Execute (frame *rtda.Frame){
	_astore(frame,2)
}

type ASTORE_3 struct {
	base.NoOperandsInstruction
}


func (astore_3 *ASTORE_3) Execute (frame *rtda.Frame){
	_astore(frame,3)
}

func _astore(frame *rtda.Frame,index uint){
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index,val)
}
