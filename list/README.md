# List Package

The usual caveats apply: do NOT run this in prod, it is stupid to use some hacky demo code for a linked list.

To see it in action run `go test` from this dir.

I was inspired to write this up after completing the [Go linked list challenge](https://exercism.org/tracks/go/exercises/simple-linked-list) on Exercism. I wanted to try writing a list type that has,

* a single type for lists and nodes, unlike the Exercism exercise interface which has separate types
* generic list node type so that the value it holds can be any type

Also I cleaned up the interface to suit my personal preference.
