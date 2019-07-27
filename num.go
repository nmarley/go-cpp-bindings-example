package num

// #cgo LDFLAGS: -L./src -lnummer -lstdc++
// #cgo CXXFLAGS: -std=c++14 -I./src
// #include "num.h"
// #include <stdlib.h>
import "C"
import "unsafe"

type GoNum struct {
	num C.Num
}

func Something() {
	// Example seed, used to generate private key. Always use
	// a secure RNG with sufficient entropy to generate a seed.
	seed := []byte{
		0, 50, 6, 244, 24, 199, 1, 25, 52, 88, 192,
		19, 18, 12, 89, 6, 220, 18, 102, 58, 209,
		82, 12, 62, 89, 110, 182, 9, 44, 20, 254, 22,
	}
	C.DoSomething(C.CBytes(seed), C.int(len(seed)))
	// C.DoSomething(seed, C.int(len(seed)))
}

func New() GoNum {
	var ret GoNum
	ret.num = C.NumInit()
	return ret
}
func (n GoNum) Free() {
	C.NumFree((C.Num)(unsafe.Pointer(n.num)))
	// C.free(pointer)
}
func (n GoNum) Inc() {
	C.NumIncrement((C.Num)(unsafe.Pointer(n.num)))
}
func (n GoNum) GetValue() int {
	return int(C.NumGetValue((C.Num)(unsafe.Pointer(n.num))))
}
