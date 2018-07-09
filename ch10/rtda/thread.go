package rtda

import "jvmgo/ch10/rtda/heap"

type Thread struct {
	pc int
	stack *Stack
}

func NewThread() *Thread{
	return &Thread{stack:newStack(1024)}
}

func (thread *Thread) PushFrame(frame *Frame){
	thread.stack.push(frame)
}


func (thread *Thread) PopFrame( ) *Frame{
	return thread.stack.pop()
}



func (thread *Thread) CurrentFrame( ) *Frame{
	return thread.stack.top()
}
func (thread *Thread) PC() int {
	return thread.pc
}

func (thread *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(thread,method)
}
func (thread *Thread) SetPc(pc int) {
	thread.pc = pc
}

func (thread *Thread) TopFrame() *Frame {
	return thread.stack.top()
}

func (thread *Thread) IsStackEmpty() bool {
	return thread.stack.isEmpty()
}

func (thread *Thread) ClearStack()  {
	thread.stack.clear()
}

