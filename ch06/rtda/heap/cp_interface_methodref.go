package heap

import "jvmgo/ch06/classfile"

type InterfaceMethodRef struct {
	MemberRef
	Method *Method
}

func newInterfaceMethodRef (pool *ConstantPool, info *classfile.ConstantInterfaceRefInfo) *InterfaceMethodRef{
	ref := &InterfaceMethodRef{}
	ref.cp = pool
	ref.copyMemberRefInfo(&info.ConstantMemberRefInfo)
	return ref
}
