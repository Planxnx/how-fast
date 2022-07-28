package concurrentmap

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/Planxnx/how-fast/benchmark"
	"github.com/lrita/cmap"
	conmap "github.com/orcaman/concurrent-map"
	conmapv2 "github.com/orcaman/concurrent-map/v2"
)

var (
	DATA_SIZE = 1024 // 1KB
	DATA      = make([]byte, DATA_SIZE)
)

func init() {
	rand.Read(DATA)
}

func BenchmarkConcurrentMap(b *testing.B) {
	benchmark.Start(b, benchmarks)
}

var benchmarks = []benchmark.LibBenchmark{
	{
		Name:    "sync_map",
		Package: "sync/map",
		Func: func(b *testing.B) {
			store := sync.Map{}

			b.Run(benchmark.MethodName("write"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						store.Store(i, DATA)
						i++
					}
				})
			})

			b.Run(benchmark.MethodName("read"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						val, ok := store.Load(i)
						if ok {
							_ = val.([]byte)
						}
						i++
					}
				})
			})

			b.Run(benchmark.MethodName("read_while_write"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						go store.Store(i+1, DATA)
						val, ok := store.Load(i)
						if ok {
							_ = val.([]byte)
						}
						i++
					}
				})
			})
		},
	},
	{
		Name:    "concurrent-mapv1",
		Package: "github.com/orcaman/concurrent-map",
		Func: func(b *testing.B) {
			store := conmap.New()

			b.Run(benchmark.MethodName("write"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						store.Set(fmt.Sprint(i), DATA)
						i++
					}
				})
			})

			b.Run(benchmark.MethodName("read"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						val, ok := store.Get(fmt.Sprint(i))
						if ok {
							_ = val.([]byte)
						}
						i++
					}
				})
			})

			b.Run(benchmark.MethodName("read_while_write"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						go store.Set(fmt.Sprint(i+1), DATA)
						val, ok := store.Get(fmt.Sprint(i))
						if ok {
							_ = val.([]byte)
						}
						i++
					}
				})
			})
		},
	},
	{
		Name:    "concurrent-mapv2",
		Package: "github.com/orcaman/concurrent-map/v2",
		Func: func(b *testing.B) {
			store := conmapv2.New[[]byte]()

			b.Run(benchmark.MethodName("write"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						store.Set(fmt.Sprint(i), DATA)
						i++
					}
				})
			})

			b.Run(benchmark.MethodName("read"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						_, _ = store.Get(fmt.Sprint(i))
						i++
					}
				})
			})

			b.Run(benchmark.MethodName("read_while_write"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						go store.Set(fmt.Sprint(i+1), DATA)
						_, _ = store.Get(fmt.Sprint(i))
						i++
					}
				})
			})
		},
	},
	{
		Name:    "cmap_non-generic",
		Package: "github.com/tidwall/cmap",
		Func: func(b *testing.B) {
			store := cmap.Cmap{}

			b.Run(benchmark.MethodName("write"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						store.Store(i, DATA)
						i++
					}
				})
			})

			b.Run(benchmark.MethodName("read"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						val, ok := store.Load(i)
						if ok {
							_ = val.([]byte)
						}
						i++
					}
				})
			})

			b.Run(benchmark.MethodName("read_while_write"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						go store.Store(i+1, DATA)
						val, ok := store.Load(i)
						if ok {
							_ = val.([]byte)
						}
						i++
					}
				})
			})
		},
	},
	{
		Name:    "cmap_generic",
		Package: "github.com/tidwall/cmap",
		Func: func(b *testing.B) {
			store := cmap.Map[int, []byte]{}

			b.Run(benchmark.MethodName("write"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						store.Store(i, DATA)
						i++
					}
				})
			})

			b.Run(benchmark.MethodName("read"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						_, _ = store.Load(i)
						i++
					}
				})
			})

			b.Run(benchmark.MethodName("read_while_write"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					var i int
					for pb.Next() {
						go store.Store(i+1, DATA)
						_, _ = store.Load(i)
						i++
					}
				})
			})
		},
	},
}
