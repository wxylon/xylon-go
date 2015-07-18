//http://www.cnblogs.com/vimsk/archive/2012/11/13/2768784.html
package main

import (
	. "github.com/lxn/go-winapi"
	"os"
	"syscall"
	"unsafe"
)

var (
	originWndProc HWND
)

const (
	winWidth  int32 = 500
	winHeight int32 = 300
)

func _TEXT(str string) *uint16 {
	a := syscall.StringToUTF16Ptr(str)
	//fmt.Println(a)
	return a
}

func WndProc(hwnd HWND, msg uint32, wparam uintptr, lparam uintptr) uintptr {
	switch msg {
	case WM_DESTROY:
		os.Exit(0)
	}
	return CallWindowProc(uintptr(originWndProc), hwnd, msg, wparam, lparam)
}

func main() {
	var message MSG
	var hwnd HWND
	var wproc uintptr
	hwnd = CreateWindowEx(
		WS_EX_CLIENTEDGE,
		_TEXT("EDIT"),
		_TEXT("HELLO GUI"),
		WS_OVERLAPPEDWINDOW,
		(GetSystemMetrics(SM_CXSCREEN)-winWidth)>>1,
		(GetSystemMetrics(SM_CYSCREEN)-winHeight)>>1,
		winWidth,
		winHeight,
		0,
		0,
		GetModuleHandle(nil),
		unsafe.Pointer(nil))
	wproc = syscall.NewCallback(WndProc)
	originWndProc = HWND(SetWindowLong(hwnd, GWL_WNDPROC, int32(wproc)))
	ShowWindow(hwnd, SW_SHOW)
	for {
		if GetMessage(&message, 0, 0, 0) == 0 {
			break
		}
		TranslateMessage(&message)
		DispatchMessage(&message)
	}
	os.Exit(int(message.WParam))
}
