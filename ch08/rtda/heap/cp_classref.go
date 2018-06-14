package heap

import "jvmgo/ch08/classfile"

type ClassRef struct {
	SymRef
}

func newClassRef(cp *ConstantPool,info *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = info.Name()
	return ref
}


