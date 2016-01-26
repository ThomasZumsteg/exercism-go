# Paasio

Write a program that reports network IO statistics

You are writing a [PaaS][], and you need a way to bill customers based
on network and filesystem usage.

Create a wrapper for network connections and files that can report IO
statistics. The wrapper must report:

- The total number of bytes read/written.
- The total number of read/write operations.

[PaaS]: http://en.wikipedia.org/wiki/Platform_as_a_service

To run the tests simply run the command `go test` in the exercise directory.

If the test suite contains benchmarks, you can run these with the `-bench`
flag:

    go test -bench .

For more detailed info about the Go track see the [help
page](http://exercism.io/languages/go).

## Source

Brian Matsuo [view source](https://github.com/bmatsuo)
