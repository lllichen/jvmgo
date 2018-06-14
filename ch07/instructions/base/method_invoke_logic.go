package base

import (
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
	"fmt"
)

func InvokeMethod(frame *rtda.Frame, method *heap.Method) {
	thread := frame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlotSlot := int(method.ArgSlotCount())
	if argSlotSlot > 0 {
		for i := argSlotSlot-1; i >= 0; i-- {
			slot := frame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i),slot)
		}
	}

	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		}else {
			panic(fmt.Sprintf("native method: %v.%v%v\n", method.Class().Name(), method.Name(),method.Descriptor()))
		}
	}
}
