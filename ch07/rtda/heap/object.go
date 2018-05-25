package heap

type Object struct {
	class *Class
	fields Slots
}

func (object *Object) Fields() Slots {
	return object.class.StaticVars()
}

func (object *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(object.class)
}