// Code generated by linux/mkall.go generatePtracePair("mips", "mips64"). DO NOT EDIT.

// +build linux
// +build mips mips64

package unix

import "unsafe"

// PtraceRegsMips is the registers used by mips binaries.
type PtraceRegsMips struct {
	Regs     [32]uint64
	Lo       uint64
	Hi       uint64
	Epc      uint64
	Badvaddr uint64
	Status   uint64
	Cause    uint64
}

// PtraceGetRegsMips fetches the registers used by mips binaries.
func PtraceGetRegsMips(pid int, regsout *PtraceRegsMips) error {
	return ptrace(PTRACE_GETREGS, pid, 0, uintptr(unsafe.Pointer(regsout)))
}

// PtraceSetRegsMips sets the registers used by mips binaries.
func PtraceSetRegsMips(pid int, regs *PtraceRegsMips) error {
	return ptrace(PTRACE_SETREGS, pid, 0, uintptr(unsafe.Pointer(regs)))
}

// PtraceRegsMips64 is the registers used by mips64 binaries.
type PtraceRegsMips64 struct {
	Regs     [32]uint64
	Lo       uint64
	Hi       uint64
	Epc      uint64
	Badvaddr uint64
	Status   uint64
	Cause    uint64
}

// PtraceGetRegsMips64 fetches the registers used by mips64 binaries.
func PtraceGetRegsMips64(pid int, regsout *PtraceRegsMips64) error {
	return ptrace(PTRACE_GETREGS, pid, 0, uintptr(unsafe.Pointer(regsout)))
}

// PtraceSetRegsMips64 sets the registers used by mips64 binaries.
func PtraceSetRegsMips64(pid int, regs *PtraceRegsMips64) error {
	return ptrace(PTRACE_SETREGS, pid, 0, uintptr(unsafe.Pointer(regs)))
}
