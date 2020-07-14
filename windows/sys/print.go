package sys

import (
	"syscall"
	"fmt"
)

var Color = struct {
	Black       int
	Blue        int
	Green       int
	Cyan        int
	Red         int
	Purple      int
	Yello       int
	White       int
	Gray        int
	LightBlue   int
	LightGreen  int
	LightCyan   int
	LightRed    int
	LightPurple int
	LightYello  int
	LightWhite  int
}{
	Black:       0,
	LightBlue:   1,
	LightGreen:  2,
	LightCyan:   3,
	LightRed:    4,
	LightPurple: 5,
	LightYello:  6,
	LightWhite:  7,
	Gray:        8,
	Blue:        9,
	Green:       10,
	Cyan:        11,
	Red:         12,
	Purple:      13,
	Yello:       14,
	White:       15,
}

func ColorPrint(color int, s ... interface{}) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(color))
	fmt.Print(s...)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}

func ColorPrintln(color int, s ... interface{}) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(color))
	fmt.Println(s...)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}

func ColorPrintf(color int, pattern string, s ... interface{}) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(color))
	fmt.Printf(pattern, s...)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}
