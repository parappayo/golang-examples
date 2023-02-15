# golang-examples

Some examples of using the Google Go programming language

This is mostly here because I'm quickly going to forget most of everything I ever learned about Go.

If you want to write Go for fun, consider looking into [Nim](https://nim-lang.org/) instead.

## Learning Notes

* `x := 1` is syntactic sugar for `var x int = 1` and will shadow variables in scope.
* Slices are like C++ vectors. `reallocs.go` demonstrates how to properly reserve capacity using `make()` to avoid slow reallocations at runtime.
* `list/list.go` demos some simple generics.
