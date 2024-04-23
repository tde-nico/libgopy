# libgopy

A golang library for python execution

## Usage

Remember to put the compiled library in the same folder where you compile (needed for linking), and also needed at runtime

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
