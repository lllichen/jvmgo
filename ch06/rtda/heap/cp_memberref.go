package heap

import "jvmgo/ch06/classfile"

type MemberRef struct {
	SymRef
	name string
	descriptor string
}

func (memberRef *MemberRef) copyMemberRefInfo (refInfo *classfile.ConstantMemberRefInfo ){
	memberRef.className = refInfo.ClassName()
	memberRef.name,memberRef.descriptor = refInfo.NameAndDescriptor()
}


