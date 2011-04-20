// errchk $G $D/$F.go

// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "unsafe"

func main() {
	var x unsafe.Pointer
	println(*x) // ERROR "invalid indirect.*unsafe.Pointer"
}
