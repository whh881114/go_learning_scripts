package main

import "C"
import "unsafe"

/*
#include <stdio.h>
*/

func main() {
	cstr := C.CString("Hello, world!")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
