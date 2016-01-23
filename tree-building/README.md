# Tree Building

Refactor a tree building algorithm.

Some web-forums have a tree layout, so posts are presented as a tree. However
the posts are typically stored in a database as an unsorted set of records. Thus
when presenting the posts to the user the tree structure has to be
reconstructed.

Your job will be to refactor a working but slow and ugly piece of code that
implements the tree building logic for highly abstracted records. The records
only contain an id number and a parent id number. The id number is always
between 0 (inclusive) and the length of the record list (exclusive). No record
has an parent id lower than it's own id and no record, except the root record,
has a parent id that's equal to it's own id.

To run the tests simply run the command `go test` in the exercise directory.

If the test suite contains benchmarks, you can run these with the `-bench`
flag:

    go test -bench .

For more detailed info about the Go track see the [help
page](http://exercism.io/languages/go).

## Source

 [view source]()
