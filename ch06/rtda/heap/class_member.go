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

func (classMember *ClassMember) isAccessibleTo(d *Class) bool {
	if classMember.IsPublic() {
		return true
	}
	c := classMember.class

	if classMember.IsProtected() {
		return d == c || d.isSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}

	if classMember.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}

	return d == c

}


func (classMember *ClassMember) Class() *Class{
	return classMember.class
}