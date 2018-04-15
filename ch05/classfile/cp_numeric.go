package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

type ConstantFloatInfo struct {
	val float32
}


type ConstantLongInfo struct {
	val int64
}


type ConstantDoubleInfo struct {
	val float64
}

func (cti *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	cti.val = int32( reader.readUint32())
}

func (cti *ConstantFloatInfo) readInfo(reader *ClassReader) {
	cti.val = math.Float32frombits( reader.readUint32())
}

func (cti *ConstantLongInfo) readInfo(reader *ClassReader) {
	cti.val = int64( reader.readUint64())
}

func (cti *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	cti.val = math.Float64frombits( reader.readUint64())
}