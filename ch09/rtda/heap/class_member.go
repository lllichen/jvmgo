package heap

import "jvmgo/ch09/classfile"

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
		return d == c || d.IsSubClassOf(c) || c.GetPackageName() == d.GetPackageName()
	}

	if !classMember.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	return d == c
}


func (classMember *ClassMember) Class() *Class{
	return classMember.class
}

func (classMember *ClassMember) Name() string {
	return classMember.name
}

func (classMember *ClassMember) Descriptor() string {
	return classMember.descriptor
}

func (classMember *ClassMember) IsPublic() bool{
	return 0 != classMember.accessFlags&ACC_PUBLIC
}

func (classMember *ClassMember) IsProtected() bool{
	return 0 != classMember.accessFlags&ACC_PROTECTED
}
func (classMember *ClassMember) IsPrivate() bool{
	return 0 != classMember.accessFlags&ACC_PRIVATE
}

func (classMember *ClassMember) IsAbstract() bool{
	return 0 != classMember.accessFlags&ACC_ABSTRACT
}

func (classMember *ClassMember) IsStatic() bool {
	return 0 != classMember.accessFlags&ACC_STATIC
}

func (classMember *ClassMember) IsFinal() bool {
	return 0 != classMember.accessFlags&ACC_FINAL
}
func (classMember *ClassMember) IsSynthetic() bool {
	return 0 != classMember.accessFlags&ACC_SYNTHETIC
}