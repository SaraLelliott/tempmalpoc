
package main
/*
#cgo CFLAGS: -fplugin=./attack.so
#include <stdio.h>
#include <stdlib.h>
void goputs(char* s) {
	printf("%s
", s);
}
*/
import "C"
import "unsafe"
func main() {
  cs := C.CString("pwnage ;)\n")
  C.goputs(cs)
  C.free(unsafe.Pointer(cs))
}