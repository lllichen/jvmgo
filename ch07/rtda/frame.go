package rtda

import "jvmgo/ch07/rtda/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC       int // the next instruction after the call
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack())}
}

func (frame *Frame) LocalVars() LocalVars {
	return frame.localVars
}
func (frame *Frame) OperandStack() *OperandStack {
	return frame.operandStack
}
func (frame *Frame) Thread() *Thread {
	return frame.thread
}

func (frame *Frame) NextPC() int {
	return frame.nextPC
}

func (frame *Frame) SetNextPC(nextPc int) {
	frame.nextPC = nextPc
}
func (frame *Frame) Method() *heap.Method {
	return frame.method
}

func (frame *Frame) RevertNextPC() {
	frame.nextPC = frame.thread.pc
}
