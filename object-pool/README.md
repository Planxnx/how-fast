# object-pool

```
goos: darwin
goarch: arm64
pkg: github.com/Planxnx/how-fast/concurrent-map
BenchmarkNonePool-8        	 4088990	       378.0 ns/op	    1072 B/op	       2 allocs/op
BenchmarkSyncPool-8        	120249339	        10.30 ns/op	      30 B/op	       0 allocs/op
BenchmarkGoCommonsPool-8   	  894732	      1219 ns/op	      65 B/op	       1 allocs/op
```
