package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//int 左位移
type ISHL struct {base.NoOperandsInstruction}

func (ishl *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

//int 算数右位移
type ISHR struct {base.NoOperandsInstruction}

func (ishr *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

//int 逻辑右位移
type IUSHR struct {base.NoOperandsInstruction}

func (iushr *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2)&0x1f
	result := int32(uint32(v1)>>s)
	stack.PushInt(result)
}

// long 左位移
type LSHL struct {base.NoOperandsInstruction}

func (*LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()

	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

// long 算数右位移
type LSHR struct {base.NoOperandsInstruction}

func (lshr *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result  := v1 >> s
	stack.PushLong(result)
}


// long  逻辑右位移
type LUSHR struct {base.NoOperandsInstruction}

func (lushr *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}




