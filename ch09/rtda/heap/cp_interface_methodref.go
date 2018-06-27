package heap

import "jvmgo/ch09/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodRefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (interfaceMethodRef *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if interfaceMethodRef.method == nil {
		interfaceMethodRef.resolveInterfaceMethodRef()
	}
	return interfaceMethodRef.method
}

// jvms8 5.4.3.4
func (interfaceMethodRef *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := interfaceMethodRef.cp.class
	c := interfaceMethodRef.ResolvedClass()

	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupInterfaceMethod(c, interfaceMethodRef.name, interfaceMethodRef.descriptor)

	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	interfaceMethodRef.method = method
}

func lookupInterfaceMethod(iface *Class, name ,description string ) *Method {
	for _,method := range iface.methods {
		if method.name == name && method.descriptor == description {
			return method
		}
	}
	return lookupMethodInInterfaces(iface.interfaces, name, description)
}
