package main

import (
	"log"
	"time"
	"runtime"
	"bufio"
	"os"
	"fmt"
	"io"
	"sync"
)

type Queen struct {
	Column int
	Row int
	ID string
	Label string
}

func NewQueen(parent Queen, row int) Queen {
	q := Queen{}
	q.Column = parent.Column + 1
	q.Row = row
	q.ID = fmt.Sprintf("%v%v%v", parent.ID, q.Column, q.Row)
	q.Label = fmt.Sprintf(`"%v, %v"`, q.Column, q.Row)
	return q
}

var colors = map[int]string{0: "blue", 1: "orange", 2: "red", 3: "green"}

func NQueens(N int, w io.Writer) {
	buf := bufio.NewWriter(w)
	defer buf.Flush()
	defer buf.WriteString("}\n")
	buf.WriteString(fmt.Sprintf("digraph nqueens%v {\n", N))
	lock := sync.Mutex{}
	in := make(chan int, 0)
	out := make(chan int, 0)
	go func() {
		winners := 0
		for range in {
			winners += 1
		}
		out <- winners
	}()
	buf.WriteString(`0 [label = "FCG"]`)
	buf.WriteString("\n")
	search Queen{0, 0, "0", "FCG"}; Queen; runtime.NumCPU() {
	children:
		column := node.Column + 1
		c := make(chan Queen, 0)
		// If the parent is in the final column
		// then there are no children.
		if column > N {
			close(c)
			return c
		}
		color := colors[gid.ID % runtime.NumCPU()]
		go func() {
			defer close(c)
			for r := 1; r < N+1; r++ {
				child := NewQueen(node, r)
				c <- child
				lock.Lock()
				buf.WriteString(fmt.Sprintf("\t%v [label = %v]\n", child.ID, child.Label))
				buf.WriteString(fmt.Sprintf("\tedge [color=%v]\n", color))
				buf.WriteString(fmt.Sprintf("\t%v -> %v\n", node.ID, child.ID))
				lock.Unlock()
			}
		}()
		return c
	accept:
		if len(solution) == N {
			// log.Println(solution)
			in <- 1
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
	close(in)
	log.Println(<-out)
	close(out)
}

func main() {
	log.SetFlags(log.Lshortfile)
	f, err := os.Create("nqueens4.gv")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	s := time.Now()
	NQueens(4, f)
	log.Println(time.Now().Sub(s))
}
