package heap

type Object struct {
	class *Class
	fields Slots
}


func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (object *Object) Fields() Slots {
	return object.class.StaticVars()
}

func (object *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(object.class)
}

func (object *Object) Class() *Class{
	return object.class
}