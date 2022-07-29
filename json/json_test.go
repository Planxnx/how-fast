package json

import (
	"encoding/json"
	"testing"

	"github.com/Planxnx/how-fast/benchmark"
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
		Name:      "goccy_go-json",
		Package:   "github.com/goccy/go-json",
		Marshal:   gojson.Marshal,
		Unmarshal: gojson.Unmarshal,
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
}

func BenchmarkEncoder(b *testing.B) {
	smallPayload := NewSmallPayload()
	largePayload := NewLargePayload()

	for _, bench := range benchmarks {
		if bench.Marshal == nil {
			continue
		}
		b.Run(benchmark.LibName(bench.Name), func(b *testing.B) {
			b.Run(benchmark.FieldName("size", "small"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if _, err := bench.Marshal(smallPayload); err != nil {
						b.Error(err)
					}
				}
			})
			b.Run(benchmark.FieldName("size", "large"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if _, err := bench.Marshal(largePayload); err != nil {
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
		b.Run(benchmark.LibName(bench.Name), func(b *testing.B) {
			b.Run(benchmark.FieldName("size", "small"), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					payload := SmallPayload{}
					if err := bench.Unmarshal(DataSmallPayload, &payload); err != nil {
						b.Error(err)
					}
				}
			})
			b.Run(benchmark.FieldName("size", "large"), func(b *testing.B) {
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
