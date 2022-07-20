package objectpool

import (
	"bytes"
	"context"
	"sync"
	"testing"

	commonspool "github.com/jolestar/go-commons-pool/v2"
)

func useBuffer(b *bytes.Buffer) {
	b.Write([]byte("hello"))
	b.WriteString(" world")
}

func BenchmarkNonePool(b *testing.B) {
	getPool := func() any {
		return bytes.NewBuffer(make([]byte, 0, 1024))
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			obj := getPool().(*bytes.Buffer)
			useBuffer(obj)
		}
	})
}

func BenchmarkSyncPool(b *testing.B) {
	pool := sync.Pool{
		New: func() any {
			return bytes.NewBuffer(make([]byte, 0, 1024))
		},
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			obj := pool.Get().(*bytes.Buffer)
			useBuffer(obj)
			pool.Put(obj)
		}
	})
}

// https://github.com/jolestar/go-commons-pool
func BenchmarkGoCommonsPool(b *testing.B) {
	ctx := context.Background()

	pool := commonspool.NewObjectPoolWithDefaultConfig(ctx, commonspool.NewPooledObjectFactorySimple(func(context.Context) (interface{}, error) {
		return bytes.NewBuffer(make([]byte, 0, 1024)), nil
	}))
	defer pool.Close(ctx)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			obj, err := pool.BorrowObject(ctx)
			if err != nil {
				b.Error(err)
				b.Fail()
			}
			useBuffer(obj.(*bytes.Buffer))
			if err := pool.ReturnObject(ctx, obj); err != nil {
				b.Error(err)
				b.Fail()
			}
		}
	})
}
