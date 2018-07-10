package lang

import (
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
	"jvmgo/ch11/rtda/heap"
	"fmt"
)

type StackTraceElement struct {
	fileName string
	className string
	methodName string
	lineNumber int
}


func (stackTraceElement *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		stackTraceElement.className, stackTraceElement.methodName, stackTraceElement.fileName, stackTraceElement.lineNumber)
}

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace","(I)Ljava/lang/Throwable;",fillInStackTrace)
}

func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.Thread())

	this.SetExtra(stes)
}

func createStackTraceElements(object *heap.Object, thread *rtda.Thread)[]*StackTraceElement{
	skip := distanceToObject(object.Class())+2
	frames := thread.GetFrames()[skip:]

	stes := make([]*StackTraceElement,len(frames))

	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

func distanceToObject(class *heap.Class)int{
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass(){
		distance++
	}
	return distance
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:class.SourceFile(),
		className:class.JavaName(),
		methodName:method.Name(),
		lineNumber:method.GetLineNumber(frame.NextPC()-1),
	}
}
