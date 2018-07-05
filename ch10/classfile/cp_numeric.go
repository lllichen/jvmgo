package classfile

import (
	"math"
)

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
func (cti *ConstantIntegerInfo) Value() int32 {
	return cti.val
}

func (ctf *ConstantFloatInfo) readInfo(reader *ClassReader) {
	ctf.val = math.Float32frombits( reader.readUint32())
}
func (ctf *ConstantFloatInfo) Value() float32 {
	return ctf.val
}

func (ctl *ConstantLongInfo) readInfo(reader *ClassReader) {
	ctl.val = int64( reader.readUint64())
}
func (ctl *ConstantLongInfo) Value() int64 {
	return ctl.val
}

func (ctd *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	ctd.val = math.Float64frombits( reader.readUint64())
}

func (ctd *ConstantDoubleInfo) Value() float64 {
	return ctd.val
}