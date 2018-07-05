package heap

import "jvmgo/ch10/classfile"

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

func (fieldRef *FieldRef) ResolvedField() *Field{
	if fieldRef.field == nil {
		fieldRef.resolveFieldRef()
	}
	return fieldRef.field
}


func (fieldRef *FieldRef) resolveFieldRef() {
	d := fieldRef.cp.class
	c := fieldRef.ResolvedClass()
	field := lookupField(c,fieldRef.name,fieldRef.descriptor)

	if fieldRef == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	fieldRef.field = field
}

func lookupField(class *Class, name,descriptor string) *Field {
	for _,field := range class.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _,iface := range class.interfaces {
		if field := lookupField(iface,name,descriptor);field != nil {
			return field
		}
	}

	if class.superClass != nil {
		return lookupField(class.superClass,name,descriptor)
	}
	return nil
}


