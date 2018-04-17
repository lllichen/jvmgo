package rtda

type Frame struct {
	lower *Frame
	localVars LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int // the next instruction after the call
}

func (frame *Frame) LocalVars() LocalVars {
	return frame.localVars
}
func (frame *Frame) OperandStack() OperandStack {
	return *frame.operandStack
}
func (frame *Frame) Thread() *Thread {
	return frame.thread
}
func (frame *Frame) SetNextPC(nextPc int) {
	frame.nextPC = nextPc
}

func NewFrame(maxLocals, maxStack uint) *Frame{
	return &Frame{localVars: newLocalVars(maxLocals),
	operandStack:newOperandStack(maxStack)}
}