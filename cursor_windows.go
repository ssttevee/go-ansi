package ansi

import (
	"os"
	"syscall"
	"unsafe"
)

func CursorHorizontalAbsolute(x int) {
	handle := syscall.Handle(os.Stdout.Fd())

	var csbi consoleScreenBufferInfo
	procGetConsoleScreenBufferInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&csbi)))

	var cursor coord
	cursor.x = short(x)
	cursor.y = csbi.cursorPosition.y

	if csbi.size.x < cursor.x {
		cursor.x = csbi.size.x
	}

	procSetConsoleCursorPosition.Call(uintptr(handle), uintptr(*(*int32)(unsafe.Pointer(&cursor))))
}
