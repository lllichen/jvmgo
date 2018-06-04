package references

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

type NEW struct {
	base.Index16Instruction
}

func (new *NEW) Execute(frame *rtda.Frame) {
	cp  := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(new.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}






