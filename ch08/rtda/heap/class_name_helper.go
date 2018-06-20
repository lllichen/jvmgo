package heap



func (class *Class) getComponentClassName(className string) string{
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("No array: "+ className)
}

func toClassName(desriptor string) string {
	if desriptor[0] == '[' {
		return desriptor
	}

	if desriptor[0] == 'L' {
		return desriptor[1 :len(desriptor)-1]
	}

	for className, d := range primitiveTypes {
		if d == desriptor {
			return className
		}
	}

	panic("Invalid descriptor: " + desriptor)
}