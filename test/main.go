package main

/*
#cgo LDFLAGS: ./libgopy.so
#include <stdlib.h>
#include <stdarg.h>
#include "../lib/includes/libgopy.h"
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func call_f64(fun string, args ...any) {
	str := fun
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var arg C.t_pyargs
	var a C.double
	a = args[0]
	arg.value = unsafe.Pointer(&a)
	arg.t = 'f'
	arg.next = nil

	res := C.call_f64(cstr, 1, unsafe.Pointer(&arg))

	fmt.Printf("%v: %v\n", fun, res)
}

func call_i64(fun string) {
	str := fun
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	res := C.call_i64(cstr)

	fmt.Printf("%v: %v\n", fun, res)
}

func call_byte(fun string) {
	str := fun
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	r := C.call_byte(cstr)
	res := C.GoBytes(unsafe.Pointer(r.bytes), C.int(r.size))

	fmt.Printf("%v: %v -> %v\n", fun, res, string(res))
}

func getTypes(args ...any) string {
	res := ""
	for _, a := range args {
		res += reflect.TypeOf(a).String()
	}
	return res
}

func main() {
	C.init()

	str := "test_script"
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	res1 := C.load(cstr)
	fmt.Printf("res1: %v\n", res1)
	if res1 != 0 {
		return
	}

	call_f64("func2", 6.5)
	call_i64("func1")
	call_byte("func5")

	res := getTypes(1, 2, 3, "test", 4.5, []byte("aaa"))
	fmt.Printf("\ntypes: %v\n", res)

	C.finalize()
}
