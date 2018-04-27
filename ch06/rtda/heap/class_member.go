package heap

import "jvmgo/ch06/classfile"

type ClassMember struct {
	accessFlags uint16
	name string
	descriptor string
	class *Class
}

func (classMember *ClassMember) copyMemberInfo(info *classfile.MemberInfo) {
	classMember.accessFlags = info.AccessFlags()
	classMember.name = info.Name()
	classMember.descriptor = info.Descriptor()
}