//http://www.cnblogs.com/vimsk/archive/2012/11/07/2759375.html
package main

import (
	"github.com/lxn/go-winapi"
	"strconv"
	"syscall"
)

func _TEXT(_str string) *uint16 {
	return syscall.StringToUTF16Ptr(_str)
}

func _toString(_n int32) string {
	return strconv.Itoa(int(_n))
}

func main() {
	var hwnd winapi.HWND
	cxScreen := winapi.GetSystemMetrics(winapi.SM_CXSCREEN)
	cyScreen := winapi.GetSystemMetrics(winapi.SM_CYSCREEN)
	winapi.MessageBox(hwnd, _TEXT("屏幕:(长:"+_toString(cxScreen)+";宽:"+_toString(cyScreen)+")"), _TEXT("主窗体"), winapi.MB_OK)
}
