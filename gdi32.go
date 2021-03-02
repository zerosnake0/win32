package win32

import (
	"syscall"
	"unsafe"
)

var (
	gdi32DLL = syscall.NewLazyDLL("gdi32.dll")

	procBitBlt                 = gdi32DLL.NewProc("BitBlt")
	procCreateCompatibleDC     = gdi32DLL.NewProc("CreateCompatibleDC")
	procCreateCompatibleBitmap = gdi32DLL.NewProc("CreateCompatibleBitmap")
	procCreateDIBSection       = gdi32DLL.NewProc("CreateDIBSection")
	procDeleteDC               = gdi32DLL.NewProc("DeleteDC")
	procDeleteObject           = gdi32DLL.NewProc("DeleteObject")
	procSelectObject           = gdi32DLL.NewProc("SelectObject")
)

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-bitblt
func BitBlt(hdcDest HDC, nXDest, nYDest, nWidth, nHeight int,
	hdcSrc HDC, nXSrc, nYSrc int, dwRop uint) bool {
	ret, _, _ := procBitBlt.Call(
		uintptr(hdcDest),
		uintptr(nXDest),
		uintptr(nYDest),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hdcSrc),
		uintptr(nXSrc),
		uintptr(nYSrc),
		uintptr(dwRop))
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatibledc
func CreateCompatibleDC(hdc HDC) HDC {
	ret, _, _ := procCreateCompatibleDC.Call(
		uintptr(hdc))
	return HDC(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createcompatiblebitmap
func CreateCompatibleBitmap(hdc HDC, width, height int) HBITMAP {
	ret, _, _ := procCreateCompatibleBitmap.Call(
		uintptr(hdc),
		uintptr(width),
		uintptr(height))
	return HBITMAP(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createdibsection
func CreateDIBSection(hdc HDC, pbmi *BITMAPINFO, iUsage uint, ppvBits *unsafe.Pointer, hSection HANDLE, dwOffset uint) HBITMAP {
	ret, _, _ := procCreateDIBSection.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(pbmi)),
		uintptr(iUsage),
		uintptr(unsafe.Pointer(ppvBits)),
		uintptr(hSection),
		uintptr(dwOffset))
	return HBITMAP(ret)
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deletedc
func DeleteDC(hdc HDC) bool {
	ret, _, _ := procDeleteDC.Call(
		uintptr(hdc))
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-deleteobject
func DeleteObject(hObject HGDIOBJ) bool {
	ret, _, _ := procDeleteObject.Call(
		uintptr(hObject))
	return ret != 0
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
func SelectObject(hdc HDC, hgdiobj HGDIOBJ) HGDIOBJ {
	ret, _, _ := procSelectObject.Call(
		uintptr(hdc),
		uintptr(hgdiobj))
	return HGDIOBJ(ret)
}
