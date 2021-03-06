package reserved

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
    "jvmgo/native"
    _ "jvmgo/native/java/io"
    _ "jvmgo/native/java/lang"
    _ "jvmgo/native/java/security"
    _ "jvmgo/native/java/util/concurrent/atomic"
    _ "jvmgo/native/sun/io"
    _ "jvmgo/native/sun/misc"
    _ "jvmgo/native/sun/reflect"
)


// Invoke native method
type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
