package stores

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type ISTORE struct {
	base.Index8Instruction
}


func (istore *ISTORE) Execute (frame *rtda.Frame){
	_istore(frame,uint(istore.Index))
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}


func (istore_0 *ISTORE_0) Execute (frame *rtda.Frame){
	_istore(frame,0)
}

type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (istore_1 *ISTORE_1) Execute (frame *rtda.Frame){
	_istore(frame,1)
}

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (istore_2 *ISTORE_2) Execute (frame *rtda.Frame){
	_istore(frame,2)
}


type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (istore_3 *ISTORE_3) Execute (frame *rtda.Frame){
	_istore(frame,3)
}


func _istore(frame *rtda.Frame,index uint){
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index,val)
}
