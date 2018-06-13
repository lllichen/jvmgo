package references

import (
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda/heap"
)

type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8
}

func (invokeInterface *INVOKE_INTERFACE) FetchOperands(reader *base.ByteCodeReader) {
	invokeInterface.index = uint(reader.ReadUint16())
	reader.ReadUint8()// count
	reader.ReadUint8()// must be 0
}

func (invokeInterface *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()

	methodRef := cp.GetConstant(invokeInterface.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate(){
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount()-1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}
	base.InvokeMethod(frame,methodToBeInvoked)
}

