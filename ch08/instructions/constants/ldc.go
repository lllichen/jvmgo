package constants

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
	"jvmgo/ch08/rtda/heap"
)

type LDC struct {
	base.Index8Instruction
}

func (ldc *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame,ldc.Index)
}

type LDC_W struct {
	base.Index16Instruction
}

func (ldcW *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame,ldcW.Index)
}

type LDC2_W struct {
	base.Index16Instruction
}

func (ldc2W *LDC2_W) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(ldc2W.Index)
	switch c.(type) {
	case int64 : stack.PushLong(c.(int64))
	case float64 : stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}


func _ldc(frame *rtda.Frame,index uint) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)

	switch c.(type) {
	case int32:stack.PushInt(c.(int32))
	case float32:stack.PushFloat(c.(float32))
	case string:
		internedStr := heap.JString(class.Loader(),c.(string))
		stack.PushRef(internedStr)
	//case *heap.ClassRef
	default:
		panic("todo: ldc!")
	}
}
