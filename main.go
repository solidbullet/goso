package main

// #cgo LDFLAGS: -L . -lc_test -lstdc++
// #cgo CFLAGS: -I ./
// #include "test.h"
import "C"

func main() {

	C.TestA()

}

//http://www.cnblogs.com/sohoer2003/p/4329085.html
