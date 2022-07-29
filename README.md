# How Fast ?

How fast is each Golang package? And if they were compared, how would it be? (my research and diary on golang library performance)

## Test Environment

- CPU: Apple M1 chip (8 cores)
- Memory: 16 GB LPDDR4X-4266 MHz
- Go: go1.18.4 darwin/arm64
- OS: macOS Monterrey 12.3.1

## Random

![random](./random/benchmark.svg)

## Concurrent Map

![cmap](./concurrent-map/benchmark.svg)

## JSON Serialization

### Encoder

![encoder_large](./json/benchmark_encoder_large.svg)

![encoder_small](./json/benchmark_encoder_small.svg)

### Decoder

![decoder_large](./json/benchmark_decoder_large.svg)

![decoder_small](./json/benchmark_decoder_small.svg)
