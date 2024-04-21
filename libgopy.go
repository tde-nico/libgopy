package libgopy

/*
#cgo LDFLAGS: ./lib/libgopy.so
#include <stdlib.h>
#include <stdarg.h>
#include "./lib/includes/libgopy.h"
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func init() {
	C.init()
}

func Finalize() {
	C.finalize()
}

func Load(module string) error {
	cstr := C.CString(module)
	defer C.free(unsafe.Pointer(cstr))
	res := C.load(cstr)
	if res != 0 {
		return fmt.Errorf("failed to load module: %v", module)
	}
	return nil
}

type Args struct {
	list  *C.t_pyargs
	count int
}

func (a *Args) init_args(args []any) {
	var tmp *C.t_pyargs
	for _, arg := range args {
		tmp = (*C.t_pyargs)(C.malloc(C.sizeof_t_pyargs))

		t := reflect.TypeOf(arg).String()
		switch t {
		case "float64":
			double := C.double(arg.(float64))
			tmp.value = unsafe.Pointer(&double)
			tmp.t = 'f'
		case "int64":
			integer := C.int(arg.(int64))
			tmp.value = unsafe.Pointer(&integer)
			tmp.t = 'd'
		case "int":
			integer := C.int(arg.(int))
			tmp.value = unsafe.Pointer(&integer)
			tmp.t = 'd'
		case "[]uint8":
			cstr := C.CString(string(arg.([]byte)))
			tmp.value = unsafe.Pointer(cstr)
			tmp.t = 'b'
		case "string":
			cstr := C.CString(arg.(string))
			tmp.value = unsafe.Pointer(cstr)
			tmp.t = 'b'
		default:
			fmt.Printf("Unknown type: %v\n", t)
			continue
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

func Call_f64(fun string, args ...any) float64 {
	str := fun
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var a Args
	a.init_args(args)
	defer a.free()

	res := C.call_f64(cstr, C.int(a.count), a.list)

	return float64(res)
}

func Call_i64(fun string, args ...any) int64 {
	str := fun
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var a Args
	a.init_args(args)
	defer a.free()

	res := C.call_i64(cstr, C.int(a.count), a.list)

	return int64(res)
}

func Call_byte(fun string, args ...any) []byte {
	str := fun
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var a Args
	a.init_args(args)
	defer a.free()

	r := C.call_byte(cstr, C.int(a.count), a.list)
	res := C.GoBytes(unsafe.Pointer(r.bytes), C.int(r.size))

	return res
}

func Call(fun string, args ...any) {
	str := fun
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var a Args
	a.init_args(args)
	defer a.free()

	C.call(cstr, C.int(a.count), a.list)
}