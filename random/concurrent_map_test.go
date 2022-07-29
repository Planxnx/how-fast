package concurrentmap

import (
	"math"
	"runtime"
	"testing"
	"time"

	"github.com/Planxnx/how-fast/benchmark"
	"github.com/valyala/fastrand"

	crand "crypto/rand"
	"math/big"
	"math/rand"
)

func BenchmarkRandom(b *testing.B) {
	runtime.GOMAXPROCS(8)
	benchmark.Start(b, benchmarks)
}

var benchmarks = []benchmark.LibBenchmark{
	{
		Name:    "math_rand",
		Package: "math/rand",
		Func: func(b *testing.B) {
			rand.Seed(time.Now().UnixNano())
			b.Run(benchmark.MethodName("sync"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = rand.Uint32()
				}
			})
			b.Run(benchmark.MethodName("async"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						_ = rand.Uint32()
					}
				})
			})
		},
	},
	{
		Name:    "crypto_rand",
		Package: "crypto/rand",
		Func: func(b *testing.B) {
			maxUint32 := big.NewInt(math.MaxUint32)
			b.Run(benchmark.MethodName("sync"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, err := crand.Int(crand.Reader, maxUint32)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
			b.Run(benchmark.MethodName("async"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						_, err := crand.Int(crand.Reader, maxUint32)
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			})
		},
	},
	{
		Name:    "fastrand",
		Package: "github.com/valyala/fastrand",
		Func: func(b *testing.B) {
			b.Run(benchmark.MethodName("sync"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = fastrand.Uint32n(math.MaxUint32)
				}
			})
			b.Run(benchmark.MethodName("async"), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						_ = fastrand.Uint32n(math.MaxUint32)
					}
				})
			})

		},
	},
}
