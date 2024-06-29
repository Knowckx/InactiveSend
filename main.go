package main

import (
	"time"
	"unicode/utf16"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	DllFullPath      = `your_fullpath_of_dll` // Use your local path like  `C:\InactiveSend\dll\AutoItX3_x64.dll`
	aotdll           = windows.NewLazySystemDLL(DllFullPath)
	procControlSend  = aotdll.NewProc("AU3_ControlSend")
	procControlClick = aotdll.NewProc("AU3_ControlClick")
)

var (
	SuccessErr = "The operation completed successfully."
)

func DllStr(s string) uintptr {
	utf16 := utf16.Encode([]rune(s + "\x00"))
	ret := &utf16[0]
	return uintptr(unsafe.Pointer(ret))
}

func Uint(input int) uintptr {
	out := uintptr(input)
	return out
}

func CheckDllCallError(err error) {
	if err == nil || err.Error() == SuccessErr {
		return
	}
	panic(err)
}

func ControlSend(title string, text string, controlID string, content string, flag int) uintptr {
	ret, _, err := procControlSend.Call(
		DllStr(title), DllStr(text), DllStr(controlID), DllStr(content), Uint(flag),
	)
	CheckDllCallError(err)
	return ret
}

// https://www.autoitscript.com/autoit3/docs/functions/ControlClick.htm
func ControlClick(title string, text string, controlID string, button string, clicks int, x int, y int) uintptr {
	ret, _, err := procControlClick.Call(
		DllStr(title), DllStr(text), DllStr(controlID), DllStr(button), Uint(clicks), Uint(x), Uint(y),
	)
	CheckDllCallError(err)
	return ret
}

func main() {
	appTitle := "test - Notepad"
	ControlClick(appTitle, "", "Edit1", "left", 1, 100, 50)
	time.Sleep(1 * time.Second)
	content := "Hello InactiveSend"
	ControlSend(appTitle, "", "Edit1", content, 0)
}
