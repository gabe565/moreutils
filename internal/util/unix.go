//go:build unix

package util

import "golang.org/x/sys/unix"

func Umask(mask int) int {
	return unix.Umask(mask)
}
