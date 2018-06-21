package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}


func toDescriptor(className string) string {
	if className[0] == '[' {
		return className
	}

	if d, ok := primitiveTypes[className]; ok {
		return d
	}
	return "L" + className + ";"
}


func getComponentClassName(className string) string{
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: "+ className)
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