package utils

import "testing"

/*
	Naming Convention: `Benchmark_<Group>/lib=<LibrayName>/method=<Method>`
	Reference: https://github.com/golang/proposal/blob/master/design/14313-benchmark-format.md
*/

type LibBenchmark struct {
	Name    string
	Package string
	Func    func(b *testing.B)
}

func FieldName(filed, name string) string {
	return filed + "=" + name
}

func LibName(lib string) string {
	return FieldName("lib", lib)
}

func MethodName(method string) string {
	return FieldName("method", method)
}

func Start(b *testing.B, benchmarks []LibBenchmark) {
	for _, bench := range benchmarks {
		b.Run(LibName(bench.Name), bench.Func)
	}
}
