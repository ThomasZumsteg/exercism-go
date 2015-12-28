# Robot Name

Write a program that manages robot factory settings.

When robots come off the factory floor, they have no name.

The first time you boot them up, a random name is generated, such as
RX837 or BC811.

Every once in a while we need to reset a robot to its factory settings,
which means that their name gets wiped. The next time you ask, it will
respond with a new random name.

The names must be random: they should not follow a predictable sequence.
Random names means a risk of collisions. Your solution should not allow
the use of the same name twice when avoidable. In some exercism language
tracks there are tests to ensure that the same name is never used twice.

To run the tests simply run the command `go test` in the exercise directory.

If the test suite contains benchmarks, you can run these with the `-bench`
flag:

    go test -bench .

For more detailed info about the Go track see the [help
page](http://help.exercism.io/getting-started-with-go.html).

## Source

A debugging session with Paul Blackwell at gSchool. [view source](http://gschool.it)
