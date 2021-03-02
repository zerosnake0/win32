package win32

import (
	"errors"
	"image"
	"reflect"
	"unsafe"
)

func CaptureScreen(hwnd HWND) (*image.RGBA, error) {
	r := GetClientRect(hwnd)
	if r == nil {
		return nil, errors.New("unable to get client rect")
	}
	hdc := GetDC(hwnd)
	if hdc == 0 {
		return nil, errors.New("unable to get DC")
	}
	defer ReleaseDC(hwnd, hdc)

	cdc := CreateCompatibleDC(hdc)
	if cdc == 0 {
		return nil, errors.New("unable to create compatible dc")
	}
	defer DeleteDC(cdc)

	w := int(r.Right - r.Left)
	h := int(r.Bottom - r.Top)

	bt := BITMAPINFO{}
	bt.BmiHeader.BiSize = uint32(reflect.TypeOf(bt.BmiHeader).Size())
	bt.BmiHeader.BiWidth = int32(w)
	bt.BmiHeader.BiHeight = -int32(h)
	bt.BmiHeader.BiPlanes = 1
	bt.BmiHeader.BiBitCount = 32
	bt.BmiHeader.BiCompression = BI_RGB

	var bitsPtr unsafe.Pointer
	bitmap := CreateDIBSection(cdc, &bt, DIB_RGB_COLORS, &bitsPtr, 0, 0)
	if bitmap == 0 || bitmap == ERROR_INVALID_PARAMETER {
		return nil, errors.New("unable to create dib section")
	}
	defer DeleteObject(HGDIOBJ(bitmap))

	sr := SelectObject(cdc, HGDIOBJ(bitmap))
	if sr == 0 || sr == HGDI_ERROR {
		return nil, errors.New("unable to select object")
	}

	if !BitBlt(cdc, 0, 0, w, h, hdc, 0, 0, SRCCOPY) {
		return nil, errors.New("unable to bit blt")
	}

	var slice []byte
	hdrp := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	hdrp.Data = uintptr(bitsPtr)
	hdrp.Len = w * h * 4
	hdrp.Cap = w * h * 4

	imageBytes := make([]byte, len(slice))
	for i := 0; i < len(imageBytes); i += 4 {
		imageBytes[i], imageBytes[i+1], imageBytes[i+2], imageBytes[i+3] =
			slice[i+2], slice[i+1], slice[i], slice[i+3]
	}
	return &image.RGBA{
		Pix:    imageBytes,
		Stride: 4 * w,
		Rect:   image.Rect(0, 0, w, h),
	}, nil
}
