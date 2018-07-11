package heap

func (class *Class) IsAssignableFrom(other *Class) bool {
	s, t := other, class
	if s == t {
		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			if !t.IsInterface() {
				return s.IsSubClassOf(t)
			} else {
				return s.IsImplements(t)
			}
		} else {
			if !t.IsInterface(){
				return t.isJlObject()
			}else {
				t.IsSuperClassOf(s)
			}
		}
	} else {
		if !t.IsArray() {
			if !t.IsInterface() {
				return t.isJlObject()
			} else {
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.IsAssignableFrom(sc)
		}

	}

	return false
}

func (class *Class) IsSubClassOf(other *Class) bool {
	for c := class.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (class *Class) IsImplements(iface *Class) bool {
	for c := class; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (class *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range class.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// c extends self
func (class *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(class)
}
