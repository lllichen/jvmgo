package references

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
	"jvmgo/ch06/rtda/heap"
)

type PUT_STATIC struct {
	base.Index16Instruction
}

func (putStatic *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(putStatic.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedFieldRef()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeRef")
	}
	if !field.IsFianl() {
		if currentClass != class || currentMethod.Name() != "<clint>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	decriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch decriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	}
}
