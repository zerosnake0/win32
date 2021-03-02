package win32

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

func UTF16PtrFromString(s *string) (uintptr, error) {
	if s == nil {
		return 0, nil
	}
	v, err := windows.UTF16PtrFromString(*s)
	if err != nil {
		return 0, err
	}
	return uintptr(unsafe.Pointer(v)), nil
}
