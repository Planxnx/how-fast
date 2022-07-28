# concurrent-map

```
BenchmarkSyncMap/Write-8   	 						 3543854	       288.1 ns/op	      63 B/op	       3 allocs/op
BenchmarkSyncMap/Read-8    							99305956	        13.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMap/ReadWhileSet-8         	 		 2421800	       420.4 ns/op	     100 B/op	       4 allocs/op
BenchmarkConcurrentMap/V1/Write-8       			13809352	       124.0 ns/op	      62 B/op	       3 allocs/op
BenchmarkConcurrentMap/V1/Read-8        			23943990	        93.62 ns/op	      15 B/op	       1 allocs/op
BenchmarkConcurrentMap/V1/ReadWhileSet-8         	 3723399	       286.7 ns/op	     120 B/op	       5 allocs/op
BenchmarkConcurrentMap/V2/Write-8                	14476321	        75.82 ns/op	      29 B/op	       2 allocs/op
BenchmarkConcurrentMap/V2/Read-8                 	28803369	        60.70 ns/op	      15 B/op	       1 allocs/op
BenchmarkConcurrentMap/V2/ReadWhileSet-8         	 6118351	       227.7 ns/op	     111 B/op	       4 allocs/op
BenchmarkCMap/Non-Generic/Write-8                	13734174	       101.2 ns/op	      71 B/op	       2 allocs/op
BenchmarkCMap/Non-Generic/Read-8                 	35319914	        37.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkCMap/Non-Generic/ReadWhileSet-8         	 2831402	       477.3 ns/op	      79 B/op	       2 allocs/op
BenchmarkCMap/Generic/Write-8                    	19595368	        67.80 ns/op	      42 B/op	       1 allocs/op
BenchmarkCMap/Generic/Read-8                     	30293053	        40.85 ns/op	       8 B/op	       1 allocs/op
BenchmarkCMap/Generic/ReadWhileSet-8             	 2611936	       440.4 ns/op	      80 B/op	       3 allocs/op
```
