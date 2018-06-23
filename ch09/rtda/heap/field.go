package heap

import (
	"jvmgo/ch09/classfile"
)

type Field struct {
	ClassMember
	constValueIndex uint
	slotId uint
}


func newFields (class *Class, cfFields []*classfile.MemberInfo ) []*Field{
	fields := make ([]*Field ,len(cfFields))
	for i, cfFields := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfFields)
		fields[i].copyAttributes(cfFields)
	}
	return fields
}



func (field *Field) isLongOrDouble () bool{
	return field.descriptor == "J" || field.descriptor == "D"
}
func (field *Field) SlotId() uint {
	return field.slotId
}
func (field *Field) Descriptor() string{
	return field.ClassMember.descriptor
}
func (field *Field) copyAttributes(info *classfile.MemberInfo) {
	if valAttr := info.ConstantValueAttribute();valAttr != nil {
		field.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}
func (field *Field) IsStatic() bool {
	return 0 != field.accessFlags&ACC_STATIC
}
func (field *Field) IsFinal() bool{
	return 0 != field.accessFlags&ACC_FINAL
}
func (field *Field) ConstValueIndex() uint {
	return field.constValueIndex
}

