package reuseport

import (
	"syscall"

	"golang.org/x/sys/windows"
)

// available returns whether or not SO_REUSEADDR and SO_REUSEPORT are available in the OS.
func available() (reuseaddr, reuseport bool) {
	return true, false
}

func Control(network, address string, c syscall.RawConn) (err error) {
	controlErr := c.Control(func(fd uintptr) {
		err = windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_REUSEADDR, 1)
	})
	if controlErr != nil {
		err = controlErr
	}
	return
}
