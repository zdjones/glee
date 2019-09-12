package main

import (
	"github.com/benbjohnson/glee"
)

func lssEqualLen() {
	a := glee.String(3)
	b := glee.String(3)
	glee.Assert(a[0] == b[0])
	glee.Assert(a[1] == b[1])

	if a < b {
		return
	}
	return
}
