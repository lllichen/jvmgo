package heap

import "jvmgo/ch06/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef (pool *ConstantPool,info *classfile.ConstantFieldRefInfo) *FieldRef{
	ref := &FieldRef{}
	ref.cp = pool
	ref.copyMemberRefInfo(&info.ConstantMemberRefInfo)
	return ref
}
