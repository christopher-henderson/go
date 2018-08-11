# The GoSearch Programming Language

GoSearch is a dialect of the Go programming language which implements the `search` keyword.

Given two facts:
* The node from which no decisions have yet been made (henceforth known as the First Choice Generator, or FCG)
* A block of code which returns a generator channel that defines iteration of the children of a given node.

...then any depth-first search may be conducted on the user's behalf.

Given either/or both of the following facts:
* A block of code that accepts a slice of user defined type and a candidate node and returns a boolean indicating whether that node is grounds for rejection of that subbranch.
* A block of code that accepts a slice of user defined type and returns a boolean indicating whether a complete solution has been found.

...then any backracking algorithm may be conducted on the user's behalf.

The structure of a `search` block is as follows:

```go
search FCG; typeof(FCG)[; Concurrency Level] {
	children:
		...
	[
	accept:
		...
	reject:
		...
	]
}
```

Where `Concurrency Level` is an optional integer value > 0 and defines the number of concurrency subbranch searches allowed by the search engine. `accept` and `reject` are optional code blocks intended for defining bactracking algorithms.
