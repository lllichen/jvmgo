package loads

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
	"jvmgo/ch11/rtda/heap"
)

// load reference from array
type AALOAD struct {
	base.NoOperandsInstruction
}

func (aaLoad *AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)

	refs := arrRef.Refs()
	checkIndex( len(refs), index)
	stack.PushRef(refs[index])
}

// load byte from array
type BALOAD struct {
	base.NoOperandsInstruction
}


func (baLoad *BALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)

	bytes := arrRef.Bytes()
	checkIndex( len(bytes), index)
	stack.PushInt(int32(bytes[index]))
}

// load char from array
type CALOAD struct {
	base.NoOperandsInstruction
}


func (caLoad *CALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)

	chars := arrRef.Chars()
	checkIndex( len(chars), index)
	stack.PushInt(int32(chars[index]))
}

// load double from array
type DALOAD struct {
	base.NoOperandsInstruction
}


func (daLoad *DALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)

	doubles := arrRef.Doubles()
	checkIndex( len(doubles), index)
	stack.PushDouble(doubles[index])
}

//load float from array
type FALOAD struct {
	base.NoOperandsInstruction
}

func (faLoad *FALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex( len(floats), index)
	stack.PushFloat(floats[index])
}

//load int from array
type IALOAD struct {
	base.NoOperandsInstruction
}

func (iaLoad *IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex( len(ints), index)
	stack.PushInt(ints[index])
}

//load long from array
type LALOAD struct {
	base.NoOperandsInstruction
}


func (iaLoad *LALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex( len(longs), index)
	stack.PushLong(longs[index])
}

// load short from array
type SALOAD struct {
	base.NoOperandsInstruction
}

func (saLoad *SALOAD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	shorts := arrRef.Shorts()
	checkIndex( len(shorts), index)
	stack.PushInt(int32(shorts[index]))
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndex(arrLen int, index int32) {
	if index <0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundException")
	}
}