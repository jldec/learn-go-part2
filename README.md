# learn-go-part2 (WIP)

[![CI](https://github.com/jldec/learn-go-part2/workflows/CI/badge.svg)](https://github.com/jldec/learn-go-part2/actions)

Learning Go concurrency by writing a web server for static HTML.

### [static](static/static.go)
Static HTML server based on [http.FileServer](https://pkg.go.dev/net/http#FileServer)

### [counter](counter/counter.go)
Simple thread-safe request counter implemented using channels

### [sqrt](sqrt/sqrt.go)
Playing around with floats and such, inspired by this [exercise](https://tour.golang.org/flowcontrol/8) in the Go Tour.

### Benchmarks

```
$ go test -benchmem -bench '^(BenchmarkCounter)$' ./counter

goos: darwin
goarch: amd64
pkg: github.com/jldec/learn-go-part2/counter
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkCounter-12    	 4509134	       250.7 ns/op	       0 B/op	       0 allocs/op
--- BENCH: BenchmarkCounter-12
    counter_test.go:43: 1 iterations, counter = 1
    counter_test.go:43: 100 iterations, counter = 100
    counter_test.go:43: 10000 iterations, counter = 10000
    counter_test.go:43: 1000000 iterations, counter = 1000000
    counter_test.go:43: 4509134 iterations, counter = 4509134
PASS
ok  	github.com/jldec/learn-go-part2/counter	1.582s
```