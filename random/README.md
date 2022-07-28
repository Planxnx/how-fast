# Random

![benchmark](./benchmark.svg)

```shell
$ go test -benchmem -run=^$ -bench . ./random > ./random/benchmark.txt
$ benchdraw --x=lib --group=method --y="ns/op" --input=./random/benchmark.txt --output=./random/benchmark.svg
```
