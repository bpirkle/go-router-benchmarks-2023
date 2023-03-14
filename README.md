# go-router-benchmarks-2023

This is a small project to informally compare the performance of several Go routers/frameworks.

This is NOT a complete or highly accurate comparison. It does not attempt to exercise the more advanced features of each router/framework compared, which might be meaningful. Assuming that it is run locally, it is also subject to effects from and limitations of the local environment that might not be representative of a proper production server.

The goal is only to do a rough order-of-magnitude level comparision to confirm there are no glaring differences, not to find small percentage differences.

### Building

This project is slightly fidgety, as it builds four separate binaries. This is controlled by Go build tags at the top of each of the .go files. This allows us to build the different binaries (and also have four main() functions) existing in the same codebase.

After cloning the repo, do the following:

```sh-session
make get
make build
```

These step install dependencies then build the four binaries.

If you prefer, you can also issue all the commands yourself. See the Makefile for the necessary formats, including build tags.
## Benchmarking

Wouldn't it be neat if this was written. I should do that.

The idea is that the Makefile will execute the benchmarking by starting each binary, issuing a number of requests, and then stopping the binary. I am planning to use "wrk" to do the actual benchmarking, so I will need to explain how to obtain that, and figure out how to integrate that into the makefile.