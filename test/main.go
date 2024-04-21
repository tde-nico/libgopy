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

type Args struct {
	list  *C.t_pyargs
	count int
}

func (a *Args) init_args(args []any) {
	var tmp *C.t_pyargs
loop:
	for _, arg := range args {
		tmp = (*C.t_pyargs)(C.malloc(C.sizeof_t_pyargs))

		t := reflect.TypeOf(arg).String()
	typeof:
		switch t {
		case "float64":
			double := C.double(arg.(float64))
			tmp.value = unsafe.Pointer(&double)
			tmp.t = 'f'
			break typeof
		case "int":
			integer := C.int(arg.(int))
			tmp.value = unsafe.Pointer(&integer)
			tmp.t = 'd'
			break typeof
		case "[]uint8":
			cstr := C.CString(string(arg.([]byte)))
			tmp.value = unsafe.Pointer(cstr)
			tmp.t = 'b'
			break typeof
		case "string":
			cstr := C.CString(arg.(string))
			tmp.value = unsafe.Pointer(cstr)
			tmp.t = 'b'
			break typeof
		default:
			fmt.Printf("Unknown type: %v\n", t)
			continue loop
		}
		tmp.next = nil

		a.count += 1
		if a.list == nil {
			a.list = tmp
		} else {
			tmp.next = a.list
			a.list = tmp
		}
	}
}

func (a *Args) free() {
	var tmp *C.t_pyargs

	if a.list == nil {
		return
	}
	for arg := a.list; arg != nil; {
		tmp = arg.next
		switch arg.t {
		case 'b':
			C.free(arg.value)
		default:
		}
		C.free(unsafe.Pointer(arg))
		arg = tmp
	}
}

func call_f64(fun string, args ...any) {
	str := fun
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var a Args
	a.init_args(args)
	defer a.free()

	res := C.call_f64(cstr, C.int(a.count), a.list)

	fmt.Printf("%v: %v\n", fun, res)
}

func call_i64(fun string, args ...any) {
	str := fun
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var a Args
	a.init_args(args)
	defer a.free()

	res := C.call_i64(cstr, C.int(a.count), a.list)

	fmt.Printf("%v: %v\n", fun, res)
}

func call_byte(fun string, args ...any) {
	str := fun
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var a Args
	a.init_args(args)
	defer a.free()

	r := C.call_byte(cstr, C.int(a.count), a.list)
	res := C.GoBytes(unsafe.Pointer(r.bytes), C.int(r.size))

	fmt.Printf("%v: %v -> %v\n", fun, res, string(res))
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

	call_f64("func6")
	call_i64("func1")
	call_byte("func5")

	call_f64("func7", 6.5, 10.0, 9.7, 8.2)
	call_i64("func7", 6, 10, 9, 8)
	call_byte("func7", "Hello", "World", "Go", "Python")
	call_byte("func7", []byte("Hello"), []byte("World"), []byte("Go"), []byte("Python"))

	C.finalize()
}
