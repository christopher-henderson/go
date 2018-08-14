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
This code was a constructive demonstration of my master's thesis. The thesis pointed out that languages (such as Prolog) offered search as a first class citizen, but that these languages were often considered obscure and too scientific (...such as Prolog). The research attempts to bring easy to implement, and efficient, graph search to imperative/procedural programming languages in such a way that no programmer would ever dread such algorithms again.

## Building

```bash
$ git clone git@github.com:christopher-henderson/GoSearch.git
$ cd GoSearch/src
$ ./make.bash
```

This makes the assumption that you have a working Go compiler in your path already to use as a bootstrap.

The target compiler will be available at `GoSearch/bin/go`.

## Motivating Examples
A small collection of examples may be found in the `examples` directory. The following is a complete implementation of the NQueens problem, solved using a number of goroutines equal to the number of CPUs available to the system at runtime:

```go
package main

import (
	"log"
	"time"
	"runtime"
)

type Queen struct {
	Column int
	Row int
}

func NQueens(N int) {
	search Queen{0, 0}; Queen; runtime.NumCPU() {
	children:
		column := node.Column + 1
		c := make(chan Queen, 0)
		// If the parent is in the final column
		// then there are no children.
		if column > N {
			close(c)
			return c
		}
		go func() {
			defer close(c)
			for r := 1; r < N+1; r++ {
				c <- Queen{column, r}
			}
		}()
		return c
	accept:
		if len(solution) == N {
			// stdout is expensive, so you
			// can get a hefty speedup by
			// commenting this out.
			log.Println(solution)
			return true
		}
		return false
	reject:
		row, column := node.Row, node.Column
		for _, q := range solution {
			r, c := q.Row, q.Column
			if row == r ||
				column == c ||
				row+column == r+c ||
				row-column == r-c {
				return true
			}
		}
		return false
	}
}

func main() {
	log.SetFlags(log.Lshortfile)
	s := time.Now()
	NQueens(8)
	log.Println(time.Now().Sub(s))
}

```

## Efficiency
@TODO pull from the paper, although the Rust version is 30x faster on NQueens.

## Compatibility
This compiler is based off of the Go project as of release 1.10.2 ([71bdbf431b79dff61944f22c25c7e085ccfc25d5](https://github.com/christopher-henderson/GoSearch/commit/71bdbf431b79dff61944f22c25c7e085ccfc25d5)). Due to the reservation of `search`, `children`, `accept`, and `reject` in the language the API with standard Go is broken. The _ABI_, however, remains intact. The result is that code written in the GoSearch dialect may be compiled by this project and then later linked into normal Go using the standard compiler.
