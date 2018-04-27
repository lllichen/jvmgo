package heap

import "jvmgo/ch06/classfile"

type Constant interface {

}

type ConstantPool struct {
	class *Class
	consts []Constant
}

func newConstantPool (class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant,cpCount)
	rtCp := &ConstantPool{class,consts}
	for i := 1 ; i< cpCount ; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo :
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value() //int32
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value() //float32
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value() //int64
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value() //float64
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()  //string
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp,classInfo)
		case *classfile.ConstantFieldRefInfo:
			filedRefInfo := cpInfo.(*classfile.ConstantFieldRefInfo)
			consts[i] = newFieldRef(rtCp,filedRefInfo)
 		case *classfile.ConstantMethodRefInfo:
			methodRefInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			consts[i] = newMethodRef(rtCp,methodRefInfo)
		case *classfile.ConstantInterfaceRefInfo:
			interfaceRefInfo := cpInfo.(*classfile.ConstantInterfaceRefInfo)
			consts[i] = newInterfaceMethodRef(rtCp,interfaceRefInfo)
 		}
	}
	return rtCp
}


