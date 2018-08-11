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

Where `Concurrency Level` is an optional integer value > 0 and defines the number of concurrent subbranch searches allowed by the search engine. `accept` and `reject` are optional code blocks intended for defining bactracking algorithms.

The type of node that this algorithm is searching through is required merely due to a technical difficulty in implementing this feature on my own with no real access to the Go compiler maintainers.

## Where did this come from?
This code was a constructive demonstration of my master's thesis. The thesis pointed out that languages (such as Prolog) offered search as a first class citizen, but that these languages were often considered obscure and too scientific (...such as Prolog). The research attempts to bring easy to implement, and efficient, graph search to imperative/procedural programming languages in such a way that no programmer will ever dread such algorithms again.
