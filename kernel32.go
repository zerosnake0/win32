package win32

import (
	"syscall"
)

var (
	kernel32DLL = syscall.NewLazyDLL("kernel32.dll")
)
