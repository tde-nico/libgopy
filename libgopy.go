package libgopy

/*
#cgo LDFLAGS: ./libgopy.so
#include "./lib/includes/libgopy.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func Init() {
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

		//t := reflect.TypeOf(arg).String()
		switch v := arg.(type) {
		case int64:
			integer := C.i64(v)
			tmp.value = unsafe.Pointer(&integer)
			tmp.t = C.TYPE_INT
		case int32:
			integer := C.i64(v)
			tmp.value = unsafe.Pointer(&integer)
			tmp.t = C.TYPE_INT
		case int16:
			integer := C.i64(v)
			tmp.value = unsafe.Pointer(&integer)
			tmp.t = C.TYPE_INT
		case int8:
			integer := C.i64(v)
			tmp.value = unsafe.Pointer(&integer)
			tmp.t = C.TYPE_INT
		case int:
			integer := C.i64(v)
			tmp.value = unsafe.Pointer(&integer)
			tmp.t = C.TYPE_INT
		case uint64:
			integer := C.u64(v)
			tmp.value = unsafe.Pointer(&integer)
			tmp.t = C.TYPE_UINT
		case uint32:
			integer := C.u64(v)
			tmp.value = unsafe.Pointer(&integer)
			tmp.t = C.TYPE_UINT
		case uint16:
			integer := C.u64(v)
			tmp.value = unsafe.Pointer(&integer)
			tmp.t = C.TYPE_UINT
		case uint8:
			integer := C.u64(v)
			tmp.value = unsafe.Pointer(&integer)
			tmp.t = C.TYPE_UINT
		case uint:
			integer := C.u64(v)
			tmp.value = unsafe.Pointer(&integer)
			tmp.t = C.TYPE_UINT
		case float64:
			double := C.double(v)
			tmp.value = unsafe.Pointer(&double)
			tmp.t = C.TYPE_FLOAT
		case float32:
			double := C.double(v)
			tmp.value = unsafe.Pointer(&double)
			tmp.t = C.TYPE_FLOAT
		case []uint8:
			var cbytes C.t_pybytes
			cbytes.bytes = (*C.uchar)(C.CBytes(v))
			cbytes.size = C.uint(len(v))
			tmp.value = unsafe.Pointer(&cbytes)
			tmp.t = C.TYPE_BYTES
		case string:
			var cbytes C.t_pybytes
			cbytes.bytes = (*C.uchar)(unsafe.Pointer(C.CString(v)))
			cbytes.size = C.uint(len(v))
			tmp.value = unsafe.Pointer(&cbytes)
			tmp.t = C.TYPE_BYTES
		default:
			fmt.Printf("Unknown type: %v\n", v)
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
		case C.TYPE_BYTES:
			C.free(unsafe.Pointer((*C.t_pybytes)(arg.value).bytes))
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

func Call_u64(fun string, args ...any) uint64 {
	str := fun
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var a Args
	a.init_args(args)
	defer a.free()

	res := C.call_u64(cstr, C.int(a.count), a.list)

	return uint64(res)
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
