package heap

import "jvmgo/ch07/classfile"

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

func (methodRef *MethodRef) resolveMethod() *Method {
	if methodRef.method == nil {
		methodRef.resolveMethodRef()
	}
	return methodRef.method
}
func (methodRef *MethodRef) resolveMethodRef() {
	d := methodRef.cp.class
	c := methodRef.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookeupMethod(c,methodRef.name,methodRef.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodErrors")
	}
	if !method.isAccessibleTo(d){
		panic("java.lang.IllegalAccessError")
	}
	methodRef.method = method
}
func lookeupMethod(class *Class, name string, descriptor string) *Method {
	method := LookupMethodInClass(class,name.descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}


