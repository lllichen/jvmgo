package references

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (invokeStatic *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(invokeStatic.Index).(*heap.MethodRef)
	resolveMethod := methodRef.ResolvedMethod()
	if !resolveMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	class := resolveMethod.Class()
	if !class.InitStarted(){
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	base.InvokeMethod(frame,resolveMethod)
}

