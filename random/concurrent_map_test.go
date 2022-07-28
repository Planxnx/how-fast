package concurrentmap

import (
	"math"
	"testing"
	"time"

	"github.com/Planxnx/how-fast/benchmark"
	"github.com/valyala/fastrand"

	crand "crypto/rand"
	"math/big"
	"math/rand"
)

func BenchmarkRandom(b *testing.B) {
	benchmark.Start(b, benchmarks)
}

var benchmarks = []benchmark.LibBenchmark{
	{
		Name:    "math_rand",
		Package: "math/rand",
		Func: func(b *testing.B) {
			rand.Seed(time.Now().UnixNano())
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = rand.Uint32()
				}
			})
		},
	},
	{
		Name:    "crypto_rand",
		Package: "crypto/rand",
		Func: func(b *testing.B) {
			maxUint32 := big.NewInt(math.MaxUint32)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, err := crand.Int(crand.Reader, maxUint32)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		},
	},
	{
		Name:    "fastrand",
		Package: "github.com/valyala/fastrand",
		Func: func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = fastrand.Uint32n(math.MaxUint32)
				}
			})
		},
	},
}
