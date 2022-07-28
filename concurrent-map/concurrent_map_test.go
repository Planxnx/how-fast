package concurrentmap

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

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

func BenchmarkSyncMap(b *testing.B) {
	store := sync.Map{}

	b.Run("Write", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			var i int
			for pb.Next() {
				store.Store(i, DATA)
				i++
			}
		})
	})

	b.Run("Read", func(b *testing.B) {
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

	b.Run("ReadWhileSet", func(b *testing.B) {
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
}

// github.com/orcaman/concurrent-map
// github.com/orcaman/concurrent-map/v2
func BenchmarkConcurrentMap(b *testing.B) {

	b.Run("V1", func(b *testing.B) {
		store := conmap.New()

		b.Run("Write", func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				var i int
				for pb.Next() {
					store.Set(fmt.Sprint(i), DATA)
					i++
				}
			})
		})

		b.Run("Read", func(b *testing.B) {
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

		b.Run("ReadWhileSet", func(b *testing.B) {
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
	})

	b.Run("V2", func(b *testing.B) {
		store := conmapv2.New[[]byte]()

		b.Run("Write", func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				var i int
				for pb.Next() {
					store.Set(fmt.Sprint(i), DATA)
					i++
				}
			})
		})

		b.Run("Read", func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				var i int
				for pb.Next() {
					_, _ = store.Get(fmt.Sprint(i))
					i++
				}
			})
		})

		b.Run("ReadWhileSet", func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				var i int
				for pb.Next() {
					go store.Set(fmt.Sprint(i+1), DATA)
					_, _ = store.Get(fmt.Sprint(i))
					i++
				}
			})
		})
	})
}

// github.com/lrita/cmap
func BenchmarkCMap(b *testing.B) {
	b.Run("Non-Generic", func(b *testing.B) {
		store := cmap.Cmap{}

		b.Run("Write", func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				var i int
				for pb.Next() {
					store.Store(i, DATA)
					i++
				}
			})
		})

		b.Run("Read", func(b *testing.B) {
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

		b.Run("ReadWhileSet", func(b *testing.B) {
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
	})
	b.Run("Generic", func(b *testing.B) {
		store := cmap.Map[int, []byte]{}

		b.Run("Write", func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				var i int
				for pb.Next() {
					store.Store(i, DATA)
					i++
				}
			})
		})

		b.Run("Read", func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				var i int
				for pb.Next() {
					_, _ = store.Load(i)
					i++
				}
			})
		})

		b.Run("ReadWhileSet", func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				var i int
				for pb.Next() {
					go store.Store(i+1, DATA)
					_, _ = store.Load(i)
					i++
				}
			})
		})
	})
}
