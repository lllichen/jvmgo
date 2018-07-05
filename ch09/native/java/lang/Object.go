package lang

import (
	"jvmgo/ch09/native"
	"jvmgo/ch09/rtda"
	"unsafe"
)

func init(){
	native.Register("java/lang/Object","getClass","()Ljava/lang/Class;",getClass)
	native.Register("java/lang/Object", "hashCode", "()I", hashCode)
}

func getClass(frame *rtda.Frame){
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}

func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}
