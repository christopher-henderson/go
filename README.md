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