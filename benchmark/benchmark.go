package benchmark

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

func LibName(lib string) string {
	return "lib=" + lib
}

func MethodName(method string) string {
	return "method=" + method
}

func Start(b *testing.B, benchmarks []LibBenchmark) {
	for _, bench := range benchmarks {
		b.Run(LibName(bench.Name), bench.Func)
	}
}
