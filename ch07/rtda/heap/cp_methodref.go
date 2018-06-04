package heap

import "jvmgo/ch06/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(pool *ConstantPool,info *classfile.ConstantMethodRefInfo) *MethodRef{
	ref := &MethodRef{}
	ref.cp = pool
	ref.copyMemberRefInfo(&info.ConstantMemberRefInfo)
	return ref
}
