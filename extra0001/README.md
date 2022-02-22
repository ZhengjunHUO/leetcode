## Test
```bash
$ go test -v
=== RUN   TestXorProduct
--- PASS: TestXorProduct (0.00s)
PASS
ok  	github.com/ZhengjunHUO/leetcode/extra	0.092s
```

## Benchmark. compare 2 methods
```bash
$ go test -bench . -benchmem
goos: darwin
goarch: arm64
pkg: github.com/ZhengjunHUO/leetcode/extra
BenchmarkXorproduct-8     	1000000000	         0.3149 ns/op	       0 B/op	       0 allocs/op
BenchmarkXorProductBF-8   	   38329	     31300 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/ZhengjunHUO/leetcode/extra	1.968s
```
