//go:build !windows
// +build !windows

package terminal

import (
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

type windowSize struct {
	rows uint16
	cols uint16
}

func GetTermSize() (int, int) {
	var (
		out *os.File
		err error
		sz  windowSize
	)
	if runtime.GOOS == "openbsd" {
		out, err = os.OpenFile("/dev/tty", os.O_RDWR, 0)
		if err != nil {
			return 0, 0
		}

	} else {
		out, err = os.OpenFile("/dev/tty", os.O_WRONLY, 0)
		if err != nil {
			return 0, 0
		}
	}
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, out.Fd(), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&sz)))

	return int(sz.cols), int(sz.rows)
}
