package reuseport

import (
	"syscall"
)

// available returns whether or not SO_REUSEADDR and SO_REUSEPORT are available in the OS.
func available() (reuseaddr, reuseport bool) {
	return false, false
}

func Control(network, address string, c syscall.RawConn) error {
	return nil
}
