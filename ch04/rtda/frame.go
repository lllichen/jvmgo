package rtda

type Frame struct {
	lower *Frame
	localVars LocalVars
	operandStack *OperandStack
}

func (frame *Frame) LocalVars() LocalVars {
	return frame.localVars
}
func (frame *Frame) OperandStack() OperandStack {
	return *frame.operandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame{
	return &Frame{localVars: newLocalVars(maxLocals),
	operandStack:newOperandStack(maxStack)}
}