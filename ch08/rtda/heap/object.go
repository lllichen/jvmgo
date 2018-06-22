package heap

type Object struct {
	class *Class
	data interface{}
}


func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		data: newSlots(class.instanceSlotCount),
	}
}

func (object *Object) Fields() Slots {
	return object.data.(Slots)
}

func (object *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(object.class)
}

func (object *Object) Class() *Class{
	return object.class
}

func (object *Object) SetRefVar(name, descriptor string, ref *Object)  {
	field := object.class.getField(name, descriptor, false)
	slots := object.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

