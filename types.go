package win32

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types
type (
	PVOID   uintptr
	HANDLE  PVOID
	HWND    HANDLE
	HDC     HANDLE
	HBITMAP HANDLE
	HGDIOBJ HANDLE
	// long   int32
	// LONG   long
	// DWORD unsignedLong
)

// https://docs.microsoft.com/en-us/windows/win32/api/windef/ns-windef-rect
type RECT struct {
	// LONG
	Left, Top, Right, Bottom int32
}

type RGBQUAD struct {
	RgbBlue     byte
	RgbGreen    byte
	RgbRed      byte
	RgbReserved byte
}

// https://docs.microsoft.com/en-us/previous-versions/dd183376(v=vs.85)
type BITMAPINFOHEADER struct {
	BiSize          uint32 // The number of bytes required by the structure.
	BiWidth         int32  // The width of the bitmap, in pixels.
	BiHeight        int32  // The height of the bitmap, in pixels.
	BiPlanes        uint16 // The number of planes for the target device. This value must be set to 1.
	BiBitCount      uint16 // The number of bits-per-pixel.
	BiCompression   uint32 // The type of compression
	BiSizeImage     uint32 // The size, in bytes, of the image. This may be set to zero for BI_RGB bitmaps.
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfo
type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors *RGBQUAD
}
