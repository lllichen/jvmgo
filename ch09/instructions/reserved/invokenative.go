package reserved

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
	"jvmgo/ch09/native"
)
import _ "jvmgo/ch09/native/lang"

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (invokeNative *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)

	if nativeMethod == nil {
		methodinfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodinfo)
	}

	nativeMethod(frame)
}



