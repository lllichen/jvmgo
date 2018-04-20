package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type BIPUSH struct {
	val int8
}

func (bipush *BIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	bipush.val = reader.ReadInt8()
}

func (bipush *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(bipush.val)
	frame.OperandStack().PushInt(i)
}

type SIPUSH struct {
	val int16
}

func (siPush *SIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	siPush.val = reader.ReadInt16()
}

func (siPush *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(siPush.val)
	frame.OperandStack().PushInt(i)
}

