# Concurrent Map

![benchmark](./benchmark.svg)

```shell
$ go test -benchmem -run=^$ -bench . ./concurrent-map > ./concurrent-map/benchmark.txt
$ benchdraw --x=lib --group=method --y="ns/op" --input=./concurrent-map/benchmark.txt --output=./concurrent-map/benchmark.svg
```
