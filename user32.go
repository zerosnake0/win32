package win32

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32DLL = windows.NewLazyDLL("user32.dll")

	procFindWindowW        = user32DLL.NewProc("FindWindowW")
	procFindWindowExW      = user32DLL.NewProc("FindWindowExW")
	procGetClientRect      = user32DLL.NewProc("GetClientRect")
	procGetDC              = user32DLL.NewProc("GetDC")
	procReleaseDC          = user32DLL.NewProc("ReleaseDC")
	procSendMessage        = user32DLL.NewProc("SendMessageW")
	procSendMessageTimeout = user32DLL.NewProc("SendMessageTimeout")
	procPostMessage        = user32DLL.NewProc("PostMessageW")
)

// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-findwindoww
func FindWindowW(className, windowName *string) (HWND, error) {
	lpClassName, err := UTF16PtrFromString(className)
	if err != nil {
		return 0, err
	}
	lpWindowName, err := UTF16PtrFromString(windowName)
	if err != nil {
		return 0, err
	}
	r, _, err := procFindWindowW.Call(lpClassName, lpWindowName)
	if r == 0 {
		return 0, err
	}
	return HWND(r), nil
}

// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-findwindowexw
func FindWindowExW(parent, childAfter HWND, className, windowName *string) (HWND, error) {
	lpClassName, err := UTF16PtrFromString(className)
	if err != nil {
		return 0, err
	}
	lpWindowName, err := UTF16PtrFromString(windowName)
	if err != nil {
		return 0, err
	}
	r, _, err := procFindWindowExW.Call(uintptr(parent), uintptr(childAfter), lpClassName, lpWindowName)
	if r == 0 {
		if err == windows.SEVERITY_SUCCESS {
			return 0, nil
		}
		return 0, err
	}
	return HWND(r), nil
}

func IterateWindowW(parent HWND, className, windowName *string, cb func(h HWND) bool) error {
	lpClassName, err := UTF16PtrFromString(className)
	if err != nil {
		return err
	}
	lpWindowName, err := UTF16PtrFromString(windowName)
	if err != nil {
		return err
	}
	child := uintptr(0)
	for {
		child, _, err = procFindWindowExW.Call(uintptr(parent), child, lpClassName, lpWindowName)
		if child == 0 {
			if err == windows.SEVERITY_SUCCESS {
				return nil
			}
			return err
		}
		if !cb(HWND(child)) {
			return nil
		}
	}
}

// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-getclientrect
func GetClientRect(hwnd HWND) *RECT {
	var rect RECT
	ret, _, _ := procGetClientRect.Call(uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect)))
	if ret == 0 {
		return nil
	}
	return &rect
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
func GetDC(hwnd HWND) HDC {
	ret, _, _ := procGetDC.Call(uintptr(hwnd))
	return HDC(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-releasedc
func ReleaseDC(hwnd HWND, hDC HDC) bool {
	ret, _, _ := procReleaseDC.Call(
		uintptr(hwnd),
		uintptr(hDC))
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessage
func SendMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procSendMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)
	return ret
}

func SendMessageTimeout(hwnd HWND, msg uint32, wParam, lParam uintptr, fuFlags, uTimeout uint32) uintptr {
	ret, _, _ := procSendMessageTimeout.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam,
		uintptr(fuFlags),
		uintptr(uTimeout))
	return ret
}

func PostMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) bool {
	ret, _, _ := procPostMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)
	return ret != 0
}
