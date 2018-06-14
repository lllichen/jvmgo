package references

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
	"jvmgo/ch08/rtda/heap"
)

type GET_STATIC struct {
	base.Index16Instruction
}

func (getStatic *GET_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fielRef := cp.GetConstant(getStatic.Index).(*heap.FieldRef)
	field := fielRef.ResolvedField()
	class := field.Class()
	if !class.InitStarted(){
		frame.RevertNextPC()
		base.InitClass(frame.Thread(),class)
		return
	}

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z','B','C','S','I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L':
		stack.PushRef(slots.GetRef(slotId))
	}
}

