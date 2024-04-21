package main

import (
	"fmt"
	"test/libgopy"
)

func main() {
	libgopy.Load("test_script")

	resf := libgopy.Call_f64("func6")
	fmt.Printf("resf: %v\n", resf)
	resi := libgopy.Call_i64("func1")
	fmt.Printf("resi: %v\n", resi)
	resb := libgopy.Call_byte("func5")
	fmt.Printf("resb: %v -> %v\n", resb, string(resb))

	libgopy.Call_f64("func7", 6.5, 10.0, 9.7, 8.2)
	libgopy.Call_i64("func7", 6, 10, 9, 8)
	libgopy.Call_byte("func7", "Hello", "World", "Go", "Python")
	libgopy.Call_byte("func7", []byte("Hello"), []byte("World"), []byte("Go"), []byte("Python"))
	libgopy.Call_f64("func7", 6.5, 10, "Hello", []byte("World"), int64(3))

	libgopy.Call("func7", 6.5, 10, "Hello", []byte("World"), int64(3))

	libgopy.Finalize()
}
