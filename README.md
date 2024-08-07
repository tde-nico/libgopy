# libgopy

<!--
#field

#groups
Libraries

#languages
C
C++
Go
Python

#frames and libs

-->

A golang library for python execution

## Usage

Remember to put the compiled library in the same folder where you compile (needed for linking), and also needed at runtime

Functions:
```go
Init() // initialize python
Load(module string) // load a python module and all the contained functions (the functions starting wiht '_' will be considered private and then not loaded)
Call(fun string, args ...any) // calls fun with args as args, returns nothing
Call_f64(fun string, args ...any) // calls fun with args as args, returns float64
Call_i64(fun string, args ...any) // calls fun with args as args, returns int64
Call_u64(fun string, args ...any) // calls fun with args as args, returns uint64
Call_byte(fun string, args ...any) // calls fun with args as args, returns []byte
Finalize() // finalize python
```
current supported types for arguments are:
```go
float64 float32
int64  int32  int16  int8  int
uint64 uint32 uint16 uint8 uint
rune
byte
[]byte
[]uint8
string
```

## Example

`main.go`
```go
package main

import (
    "fmt"

    py "github.com/tde-nico/libgopy"
)

func main() {
	py.Init() // intitialize python

    err := py.Load("test") // loads the python module test (test.py)
    if err != nil {
        panic(err)
    }

    res := py.Call_byte("hello", "world", 5, 6.7) // calls fuction hello as: func hello(string, int, float) []byte
    fmt.Printf("Result: %s\n", res)

    py.Finalize() // finalize python
}
```

`test.py`
```py
def hello(x, y, z):
    print(f"Hello from test.py {x} {y} {z}")
    return b"Hello from test.py in bytes"
```

## Benchmarks

tested with: `go test -bench=. -benchtime=10s -count=10` on Windows 10 WSL 2
```
goos: linux
goarch: amd64
pkg: github.com/tde-nico/libgopy
cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
BenchmarkLibgopy-8       4186684              2870 ns/op
BenchmarkLibgopy-8       4076690              2961 ns/op
BenchmarkLibgopy-8       4256122              2941 ns/op
BenchmarkLibgopy-8       4148332              2925 ns/op
BenchmarkLibgopy-8       4042467              2985 ns/op
BenchmarkLibgopy-8       4039228              2984 ns/op
BenchmarkLibgopy-8       4178586              2871 ns/op
BenchmarkLibgopy-8       4013100              2971 ns/op
BenchmarkLibgopy-8       3895118              2882 ns/op
BenchmarkLibgopy-8       4016323              2982 ns/op
PASS
ok      github.com/tde-nico/libgopy     139.135s
```

## Notes

The library included in this module, is compiled for python3.10.
If you want another version, change the python flags into the makefile and then recompile it

## Features

- [x] multi type argument support
- [x] variadic arguments
- [x] auto funcion import from module
- [x] gloabal variables
