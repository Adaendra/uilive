//go:build windows
// +build windows

package terminal

import (
	"os"
	"unsafe"

	"golang.org/x/sys/windows"

	"github.com/Adaendra/uilive/pkg/kernel"
)

func GetTermSize() (int, int) {
	out, err := os.Open("CONOUT$")
	if err != nil {
		return 0, 0
	}
	defer out.Close()

	var csbi windows.ConsoleScreenBufferInfo
	ret, _, _ := kernel.ProcGetConsoleScreenBufferInfo.Call(out.Fd(), uintptr(unsafe.Pointer(&csbi)))
	if ret == 0 {
		return 0, 0
	}

	return int(csbi.Window.Right - csbi.Window.Left + 1), int(csbi.Window.Bottom - csbi.Window.Top + 1)
}
