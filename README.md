# Go-BLS Signature Aggregation

This repository is a go-wrapper around @herumi's BLS implementation using the underlying [herumi/mcl](https://github.com/herumi/mcl) cryptography library written in C++. The go wrapper is also based on @herumi's original BLS C++ project [herumi/bls](https://github.com/herumi/bls). The code was modified by Prysmatic Labs in order to integrate it into Ethereum's Serenity upgrade including Proof of Stake and Sharding.

## Installing

```bash
git clone https://github.com/prysmaticlabs/go-bls
```

#### Running With Go

Make sure you have the latest version of Go installed along with [libgmp](https://gmplib.org/) required for precision arithmetic. Then, you can run the local bls tests using the go tool:

```bash
go test -bench .
goos: linux
goarch: amd64
pkg: github.com/prysmaticlabs/go-bls
BenchmarkPubkeyFromSeckey-4   	   10000	    221235 ns/op
BenchmarkSigning-4            	    5000	    274591 ns/op
BenchmarkValidation-4         	    1000	   1305703 ns/op
PASS
ok  	github.com/prysmaticlabs/go-bls	7.814s
```

#### Running With Bazel

Install Google's Bazel build tool [here](https://docs.bazel.build/versions/master/install-ubuntu.html) for your architecture. Then, run tests as follows:

```bash
bazel test //...
INFO: Analysed 3 targets (1 packages loaded).
INFO: Found 2 targets and 1 test target...
INFO: Elapsed time: 6.543s, Critical Path: 5.53s
INFO: 17 processes: 17 linux-sandbox.
INFO: Build completed successfully, 18 total actions
//:go_default_test                                                       PASSED in 2.0s

Executed 1 out of 1 test: 1 test passes.
There were tests whose specified size is too big. Use the --test_verbose_timeout_warnings command line option INFO: Build completed successfully, 18 total actions
```

## License

The original BLS code was written by @herumi under the BSD-3 software license.