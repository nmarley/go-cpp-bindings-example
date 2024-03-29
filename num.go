package num

// #cgo LDFLAGS: -L./src -lnummer -lstdc++
// #cgo CXXFLAGS: -std=c++14 -I./src
// #include "num.h"
import "C"
import "unsafe"

type GoNum struct {
	num C.Num
}

func New() GoNum {
	var ret GoNum
	ret.num = C.NumInit()
	return ret
}
func (n GoNum) Free() {
	C.NumFree((C.Num)(unsafe.Pointer(n.num)))
}
func (n GoNum) Inc() {
	C.NumIncrement((C.Num)(unsafe.Pointer(n.num)))
}
func (n GoNum) GetValue() int {
	return int(C.NumGetValue((C.Num)(unsafe.Pointer(n.num))))
}
