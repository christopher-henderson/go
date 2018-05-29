#H1 Graph Search as a Feature in Imperative/Procedural Programming Languages

#H2 Syntax

```go
search FCG; type[; concurrency] {
children:
	// Function body which returns a channel of the same type
	// as the FCG.
[accept:
	// Function body which returns a boolean on whether the solution
	// thus far should be accepted.
]
[reject:
	// Function which returns a boolean on whether the candidate
	// node, in the presence of the solution thus far, should
	// be rejected.
]
}
```

#H2 Jankiness

`search`, `children`, `accept`, and `reject` are now reserved keywords in the language. This means that the Golang API is _broken_. If ever your project, or its dependencies, use these names in their source code then you must change those names before invoking this compiler. Forking an external project merely for otherwise inocuous name choices, however, is unsatisfactory.

Fortunately, the Golang _*ABI*_ is completely intact. That is, the binaries produced by the GoSearch compiler are pure, indistinguishable, Go 1.10.2 and any code which wishes to use the GoSearch construct can compile a linkable binary and import it in an otherwise vanilla Golang project.

Let us consider the following which conducts a DFS on `patient zero`'s friends list and serves up ads to them all.

```go
package freemoney

import (
	"log"
	"github.com/obnoxious/ads" // Whoops! Plain old Go that `search`es for ads to show.
)

func PrintMoney() {
	search SocialMediaUser{"patient zero"}; SocialMediaUser {
	children:
		c := make(chan SocialMediaUser, 0)
		go func() {
			for _, friend := range node.Friends {
				log.Println("Printing money shoveling ads onto", friend)
				ads.Shovel(friend)
				c <- friend
			}
		}
		return c
	}
}
```

Well nuts, _nothing_ can compile this. The plain Golang compiler cannot handle the `search` construct, and the `ads` package uses the reserved keyword `search` internally in order to look for ads to show. In order to get this working we're going to move the `ads` dependency out of this package, compile it using GoSearch, and convey the results of the DFS to pure Go.

```go
package freemoney

import (
	"log"
)

func PrintMoney(out <-chan SocialMediaUser) {
	search SocialMediaUser{"patient zero"}; SocialMediaUser {
	children:
		c := make(chan SocialMediaUser, 0)
		go func() {
			for _, friend := range node.Friends {
				log.Println("Printing money shoveling ads onto", friend)
				out <- friend
				c <- friend
			}
		}
		return c
	}
}
```

The Golang compiler can link in precompiled binaries, but it does require a source code stub be present under $GOPATH/src. The stub requires, at minimum, the `//go:binary-only-package` pragma near the top, the package declaration, and, as of Go 1.10, the explicit `import` list.

```go
//go:binary-only-package

package freemoney

import (
        "log"
)
```

We need to compile the above GoSearch source file into a library file and place it in the appropriate `$GOPATH/pkg` directory. We also need to drop the stub into the appropriate `$GOPATH/src` directory.

```bash
$GOSEARCH/bin/go build -i -o $GOPATH/pkg/github.com/gigasmart/unicorn/freemoney.a freemoney.go
echo "//go:binary-only-package

package freemoney

import (
        \"log\"
)" >> $GOPATH/src/github/gigasmart/unicorn/freemoney/freemoney.go
```

Now we can write pure Golang which uses the compiled result of a GoSearch source file.

```go
package main

import (
	"github.com/gigasmart/unicorn/freemoney"
	"github.com/obnoxious/ads"
)

func main() {
	humanMoneyBags := make(chan SocialMediaUser, 0)
	go freemoney.PrintMoney(humanMoneyBags)
	for helplessPeople := range <-humanMoneyBags {
		ads.Shovel(helplessPeople)
	}
}
```






# H2 Motivating Example: Webcrawling

```go
seen := make(map[string]bool, 0)
domain := "golang.org"
// Conduct a DFS crawl of golang.org using eight workers.
//
// For simplicity's sake, thread safety is ignored here for the seen map,
// although it tends to work out fine since HTTP requests spread out
// read/writes enough to make contention rare.
search domain; string; 8 {
children:
	// Make a channel for which the engine can use to retrieve
	// the children of this node.
	c := make(chan Page, 0)
	hrefs := ExtractAllHrefs(node)
	seen[node] = true
	go func() {
		defer close(c)
		for _, href := range hrefs {
			if seen[href] {
				continue
			}
			if notIn(domain, href) {
				continue
			}
			c <- href 
		}
	}()
	return c
}
```