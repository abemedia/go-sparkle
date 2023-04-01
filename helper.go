//go:build darwin
// +build darwin

package sparkle

import "C"

func bool2int(b bool) C.int {
	if b {
		return 1
	}
	return 0
}
