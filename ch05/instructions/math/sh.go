package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type ISHL struct {base.NoOperandsInstruction}//int 左位移

func (ishl *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}


type ISHR struct {base.NoOperandsInstruction}//int 算数右位移


// int  逻辑右位移
type LUSHR struct {base.NoOperandsInstruction}

func (lushr *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}



type LSHL struct {base.NoOperandsInstruction} // long 左位移
type LSHR struct {base.NoOperandsInstruction} // long 算数右位移

func (lshr *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result  := v1 >> s
	stack.PushLong(result)
}

type IUSHR struct {base.NoOperandsInstruction} //  long 逻辑右位移




