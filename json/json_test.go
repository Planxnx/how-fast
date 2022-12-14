package json

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Planxnx/how-fast/utils"
	"github.com/bytedance/sonic"
	gojson "github.com/goccy/go-json"
	segmentiojson "github.com/segmentio/encoding/json"
	"github.com/wI2L/jettison"
)

func TestXxx(t *testing.T) {
	t.Error()
	data, _ := json.Marshal(NewLargePayload())
	t.Log(string(data))
}

var benchmarks = []jsonBenchmark{
	{
		Name:      "encoding_json",
		Package:   "encoding/json",
		Marshal:   json.Marshal,
		Unmarshal: json.Unmarshal,
	},
	{
		Name:      "encoding_gob",
		Package:   "encoding/gob",
		Marshal:   gobMarshal,
		Unmarshal: gobUnmarshal,
	},
	{
		Name:    "goccy_go-json",
		Package: "github.com/goccy/go-json",
		Marshal: gojson.Marshal,
	},
	{
		Name:    "jettison",
		Package: "github.com/wI2L/jettison",
		Marshal: jettison.Marshal,
	},
	{
		Name:      "segmentio_json",
		Package:   "github.com/segmentio/encoding/json",
		Marshal:   segmentiojson.Marshal,
		Unmarshal: segmentiojson.Unmarshal,
	},
	{
		Name:      "sonic",
		Package:   "github.com/bytedance/sonic",
		Marshal:   sonic.Marshal,
		Unmarshal: sonic.Unmarshal,
	},
}

func BenchmarkEncoder(b *testing.B) {
	smallPayload := NewSmallPayload()
	largePayload := NewLargePayload()
	longPayload := NewMapPayload(1000, 1)
	smallDataPayload := NewMapPayload(8, 1)
	bigDataPayload := NewMapPayload(8, 2*1024*1024)

	for _, bench := range benchmarks {
		if bench.Marshal == nil {
			continue
		}
		b.Run(utils.LibName(bench.Name), func(b *testing.B) {
			b.Run(utils.FieldName("size", "small"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if _, err := bench.Marshal(smallPayload); err != nil {
						b.Error(err)
					}
				}
			})
			b.Run(utils.FieldName("size", "large"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if _, err := bench.Marshal(largePayload); err != nil {
						b.Error(err)
					}
				}
			})
			b.Run(utils.FieldName("size", "long"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if _, err := bench.Marshal(longPayload); err != nil {
						b.Error(err)
					}
				}
			})
			b.Run(utils.FieldName("size", "small_data"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if _, err := bench.Marshal(smallDataPayload); err != nil {
						b.Error(err)
					}
				}
			})
			b.Run(utils.FieldName("size", "big_data"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if _, err := bench.Marshal(bigDataPayload); err != nil {
						b.Error(err)
					}
				}
			})
		})
	}
}

func BenchmarkDecoder(b *testing.B) {
	for _, bench := range benchmarks {
		if bench.Unmarshal == nil {
			continue
		}
		b.Run(utils.LibName(bench.Name), func(b *testing.B) {
			b.Run(utils.FieldName("size", "small"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					payload := SmallPayload{}
					if err := bench.Unmarshal(DataSmallPayload, &payload); err != nil {
						b.Error(err)
					}
				}
			})
			b.Run(utils.FieldName("size", "large"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					payload := LargePayload{}
					if err := bench.Unmarshal(DataLargePayload, &payload); err != nil {
						b.Error(err)
					}
				}
			})
		})
	}
}

type jsonBenchmark struct {
	Name      string
	Package   string
	Marshal   func(v interface{}) ([]byte, error)
	Unmarshal func(data []byte, v interface{}) error
}

func gobMarshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func gobUnmarshal(b []byte, result interface{}) error {
	buf := bytes.NewBuffer(b)
	enc := gob.NewDecoder(buf)

	err := enc.Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func NewMapPayload(n int, valueSizeBytes ...int) map[string][]byte {
	size := 8
	if len(valueSizeBytes) > 0 {
		size = valueSizeBytes[0]
	}
	m := make(map[string][]byte, n)
	for i := 0; i < n; i++ {
		v := make([]byte, size)
		_, _ = rand.Read(v)
		m[fmt.Sprint(i)] = v
	}
	return m
}
