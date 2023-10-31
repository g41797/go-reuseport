//go:build freebsd

package reuseport

import (
	"syscall"

	"golang.org/x/sys/unix"
)

// available returns whether or not SO_REUSEADDR and SO_REUSEPORT are available in the OS.
func available() (reuseaddr, reuseport bool) {
	return true, true
}

func Control(network, address string, c syscall.RawConn) (err error) {
	controlErr := c.Control(func(fd uintptr) {
		err = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
		if err != nil {
			return
		}
		err = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		if err != nil {
			return
		}
		err = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT_LB, 1)
	})
	if controlErr != nil {
		err = controlErr
	}
	return
}
