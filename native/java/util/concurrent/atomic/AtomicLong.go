package atomic

import (
	"jvmgo/native"
    "jvmgo/rtda"
)

func init() {
	native.Register("java/util/concurrent/atomic/AtomicLong", "VMSupportsCS8", "()Z", vmSupportsCS8)
}

func vmSupportsCS8(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(false)
}
