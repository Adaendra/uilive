package kernel

import "syscall"

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	ProcGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
	ProcSetConsoleCursorPosition   = kernel32.NewProc("SetConsoleCursorPosition")
	ProcFillConsoleOutputCharacter = kernel32.NewProc("FillConsoleOutputCharacterW")
)
