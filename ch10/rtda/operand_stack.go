package rtda

import (
	"math"
	"jvmgo/ch10/rtda/heap"
)

type OperandStack struct {
	size uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots:make([]Slot,maxStack ),
		}
	}
	return nil
}

func (operandStack *OperandStack) PushInt(val int32){
	 operandStack.slots[operandStack.size].num = val
	 operandStack.size++
}

func (operandStack *OperandStack) PopInt() int32{
	operandStack.size--
	return operandStack.slots[operandStack.size].num
}


func (operandStack *OperandStack) PushFloat(val float32){
	bits := math.Float32bits(val)
	operandStack.slots[operandStack.size].num = int32(bits)
	operandStack.size++
}

func (operandStack *OperandStack) PopFloat() float32{
	operandStack.size--
	val := uint32(operandStack.slots[operandStack.size].num)
	return  math.Float32frombits(val)
}



func (operandStack *OperandStack) PushLong(val int64){
	operandStack.slots[operandStack.size].num = int32(val)
	operandStack.slots[operandStack.size+1].num = int32(val>>32)
	operandStack.size+=2
}

func (operandStack *OperandStack) PopLong() int64{
	operandStack.size-=2
	low := uint32(operandStack.slots[operandStack.size].num)
	high :=  uint32(operandStack.slots[operandStack.size+1].num)
	return  int64(high)<<32 | int64(low)
}



func (operandStack *OperandStack) PushDouble(val float64){
	bits := math.Float64bits(val)
	operandStack.PushLong(int64(bits))
}

func (operandStack *OperandStack) PopDouble() float64{
	val := operandStack.PopLong()
	return math.Float64frombits(uint64(val))
}


func (operandStack *OperandStack) PushRef(val *heap.Object){
	operandStack.slots[operandStack.size].ref = val
	operandStack.size++
}

func (operandStack *OperandStack) PopRef() *heap.Object{
	operandStack.size--
	//fmt.Printf("operandStack : %+v \n",operandStack)
	//fmt.Printf("operandStack : %v \n",operandStack.size)
	ref := operandStack.slots[operandStack.size].ref
	operandStack.slots[operandStack.size].ref = nil
	return ref
}

func (operandStack *OperandStack) PushSlot(slot Slot)  {
	operandStack.slots[operandStack.size] = slot
	operandStack.size++
}


func (operandStack *OperandStack) PopSlot() Slot  {
	operandStack.size--
	return operandStack.slots[operandStack.size]

}


func (operandStack *OperandStack) PushBoolean(val bool) {
	if val {
		operandStack.PushInt(1)
	} else {
		operandStack.PushInt(0)
	}
}
func (operandStack *OperandStack) PopBoolean() bool {
	return operandStack.PopInt() == 1
}


func (operandStack *OperandStack) GetRefFromTop(n uint) *heap.Object {
	return operandStack.slots[operandStack.size-1-n].ref
}



