package references

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
	"fmt"
)

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (arrayLength *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	fmt.Printf("frame : %+v \n",frame)
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
