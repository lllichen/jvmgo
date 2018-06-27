package stores

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type FSTORE struct {
	base.Index8Instruction
}


func (fstore *FSTORE) Execute (frame *rtda.Frame){
	_fstore(frame,uint(fstore.Index))
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}


func (fstore_0 *FSTORE_0) Execute (frame *rtda.Frame){
	_fstore(frame,0)
}


type FSTORE_1 struct {
	base.NoOperandsInstruction
}


func (fstore_1 *FSTORE_1) Execute (frame *rtda.Frame){
	_fstore(frame,1)
}


type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (fstore_2 *FSTORE_2) Execute (frame *rtda.Frame){
	_fstore(frame,2)
}


type FSTORE_3 struct {
	base.NoOperandsInstruction
}


func (fstore_3 *FSTORE_3) Execute (frame *rtda.Frame){
	_fstore(frame,3)
}


func _fstore(frame *rtda.Frame,index uint){
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index,val)
}
