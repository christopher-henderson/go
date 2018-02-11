package main

import (
	"log"
	"time"
	"os"
	"runtime"
	"fmt"
)

type Queen struct {
	Column int
	Row int
}

func NQueens(N int, C int, M *runtime.MemStats) {
	search Queen{0, 0}; Queen; C {
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
	runtime.ReadMemStats(M)
}

func main() {
	log.SetFlags(log.Lshortfile)
	f, err := os.Create("timings.csv")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	var mem runtime.MemStats
	for size := 1; size <= 15; size++ {
		for cpu := 1; cpu <= runtime.NumCPU(); cpu++ {
			var nanoseconds int64
			runs := 40
			var totalMallocs uint64 = 0
			for run := 0; run < runs; run++ {
				runtime.ReadMemStats(&mem)
				mallocs := mem.Mallocs
				s := time.Now()
				NQueens(size, cpu, &mem)
				nanoseconds += time.Now().Sub(s).Nanoseconds()
				totalMallocs += mem.Mallocs - mallocs
				runtime.GC()
			}
			f.WriteString(fmt.Sprintf("%v, %v, %v, %v\n", size, cpu, time.Duration(int64(nanoseconds/int64(runs))).Milliseconds(), totalMallocs/uint64(runs)))
		}
	}
}
