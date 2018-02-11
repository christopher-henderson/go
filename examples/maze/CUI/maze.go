package main

import (
	"fmt"
	"sync"
	"time"
	"log"
	tm "github.com/buger/goterm"
)



type Maze [][]int

func (m Maze) Get(node []int) *int {
	if len(node) == 0 {
		return nil
	}
	r := node[0]
	c := node[1]
	if r >= len(m) || c >= len(m) || r < 0 || c < 0 {
		return nil
	}
	return &m[r][c]
}

var simple = Maze{{0, 0, 0, 0, 0, 0},
	{1, 1, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 0},
	{0, 1, 1, 1, 1, 1},
	{0, 1, 0, 1, 0, 0},
	{0, 0, 0, 1, 0, 0},
}

var cycle = Maze{{0, 0, 0, 0, 0, 0},
	{1, 1, 1, 1, 0, 0},
	{0, 1, 0, 1, 0, 0},
	{0, 1, 1, 1, 1, 0},
	{0, 1, 0, 1, 0, 0},
	{0, 0, 0, 1, 0, 0},
}

var loop = Maze{{0, 0, 0, 0, 0, 0},
	{1, 1, 0, 0, 0, 0},
	{0, 1, 1, 0, 0, 0},
	{0, 1, 1, 1, 1, 0},
	{0, 1, 0, 1, 0, 0},
	{0, 0, 0, 1, 0, 0},
}

var large = Maze{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1},
	{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

var branched = Maze{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
	{0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	{1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0},
	{0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0},
	{0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0},
	{0, 1, 0, 1, 1, 1, 1, 1, 0, 0, 1, 0},
	{0, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1},
	{0, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0},
	{0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0},
	{0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

type Set struct {
	maze Maze
	set  map[*int]bool
	sync.Mutex
}

func (s *Set) Contains(node []int) bool {
	k := s.maze.Get(node)
	s.Lock()
	v, ok := s.set[k]
	s.Unlock()
	return v && ok
}

func (s *Set) Put(node []int) {
	k := s.maze.Get(node)
	s.Lock()
	s.set[k] = true
	s.Unlock()
}

func (s *Set) Delete(node []int) {
	k := s.maze.Get(node)
	s.Lock()
	s.set[k] = false
	s.Unlock()
}

func NewSet(maze Maze) *Set {
	return &Set{maze: maze, set: make(map[*int]bool), Mutex: sync.Mutex{}}
}

type Map struct {
	m    map[int]*Set
	maze Maze
	sync.Mutex
}

func (m *Map) Get(gid int, solution [][]int) *Set {
	m.Lock()
	defer m.Unlock()
	s, ok := m.m[gid]
	if ok {
		return s
	}
	s = NewSet(m.maze)
	m.m[gid] = s
	for _, node := range solution {
		s.set[m.maze.Get(node)] = true
	}
	return s
}

func NewMap(maze Maze) *Map {
	return &Map{maze: maze, m: make(map[int]*Set), Mutex: sync.Mutex{}}
}

func solveMaze(maze Maze, entrance []int, concurrency int) [][][]int {
	terminal := NewTerminal(maze)
	solutions := make([][][]int, 0)
	lock := sync.Mutex{}
	cycles := NewMap(maze)
	directions := [][]int{[]int{-1, 0}, // north
		[]int{1, 0},  // south
		[]int{0, 1},  // east
		[]int{0, -1}} // west
	search entrance; []int; concurrency {
	children:
		terminal.ColorNode(node, colorMap[gid.ID % 4])
		cycle := cycles.Get(gid.ID, solution)
		cycle.Put(node)
		cs := make([][]int, 0)
		for _, d := range directions {
				child := []int{node[0] + d[0], node[1] + d[1]}
				if cycle.Contains(child) {
					continue
				}
				cs = append(cs, child)
			}
		c := make(chan []int, 0)
		go func() {
			defer close(c)
			defer terminal.PopNode(node)
			defer cycle.Delete(node)
			for _, child := range cs {
				c <- child
			}
			// We must wait until the final child is handled
			// before we remove it from the cycle detection map.
			// As such, let's give something we know the reject
			// block will pass on.
			c <- entrance
		}()
		return c
	reject:
		wait :=  time.Millisecond * 100
		terminal.HighlightConsidering(node)
		defer terminal.ClearHighlight(node)
		time.Sleep(wait)
		value := maze.Get(node)
		if value == nil || *value == 0 {
			terminal.HighlightFailed(node)
			time.Sleep(wait)
			return true
		}
		terminal.HighlightSuccess(node)
		time.Sleep(wait)
		return false
	accept:
		// The first entry is indeed in the solution,
		// and it is on the border, but is our starting point,
		// not destination.
		if len(solution) == 1 {
			return false
		}
		final := solution[len(solution) - 1]
		r := final[0]
		c := final[1]
		if r == len(maze) - 1 || c == len(maze) - 1 || r == 0 || c == 0 {
			s := make([][]int, len(solution))
			copy(s, solution)
			lock.Lock()
			solutions = append(solutions, s)
			for _, n := range solution {
				terminal.ColorNode(n, tm.GREEN)
			}
			// log.Println(solution)
			lock.Unlock()
			return true
		}
		return false
	}
	return solutions
}

const (
	RED = iota
	BLUE
	YELLOW
	CYAN
)

var colorMap = map[int]int{
	RED:    tm.RED,
	BLUE:   tm.BLUE,
	YELLOW: tm.YELLOW,
	CYAN:   tm.CYAN}


func main() {
	log.SetFlags(log.Lshortfile)
	solveMaze(branched, []int{3, -1}, 8)
	fmt.Println("Hit Enter to exit.")
	fmt.Scanln()
}






type ColorStack []int

func (c ColorStack) Push(color int) ColorStack {
	return append(c, color)
}

func (c ColorStack) Pop() ColorStack {
	if len(c) == 0 {
		return c
	}
	return c[:len(c) - 1]
}

func (c ColorStack) Current() int {
	if len(c) == 0 {
		return tm.BLACK
	}
	return c[len(c)-1]
}

const considering = tm.BLACK
const failed = tm.RED
const succeeded = tm.GREEN

type Terminal struct {
	maze Maze
	cmap map[*int]ColorStack
	sync.Mutex
}

func NewTerminal(maze Maze) *Terminal {
	t := &Terminal{maze, make(map[*int]ColorStack, 0), sync.Mutex{}}
	m := "\n"
	for _, r := range maze {
		for _, c := range r {
			m = fmt.Sprintf("%s %v", m, c)
		}
		m = fmt.Sprintln(m)
	}
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Println(m)
	tm.Flush()
	return t
}

func (t *Terminal) Write(node []int, str string) {
	x, y := MapXY(node)
	tm.MoveCursor(x, y)
	tm.Print(str)
	tm.Flush()
}

func (t *Terminal) Highlight(node []int, str string, color, background int) {
	t.Write(node, tm.Background(tm.Color(str, color), background))
}

func (t *Terminal) ColorNode(node []int, color int) {
	t.Lock()
	defer t.Unlock()
	value := t.maze.Get(node)
	if value == nil {
		return
	}
	cs, ok := t.cmap[value]
	if !ok {
		log.Panicln("This coloring should not be possible")
	}
	if cs.Current() == tm.GREEN {
		// Leave known paths green.
		return
	}
	cs = cs.Push(color)
	t.cmap[value] = cs
	fmtedValue := fmt.Sprintf("%v", *value)
	t.Write(node, tm.Color(fmtedValue, color))
}

func (t *Terminal) PopNode(node []int) {
	t.Lock()
	defer t.Unlock()
	value := t.maze.Get(node)
	if value == nil {
		return
	}
	cs, ok := t.cmap[value]
	if !ok {
		log.Panicln("This pop should not be possible")
	}
	if cs.Current() == tm.GREEN {
		return
	}
	cs = cs.Pop()
	t.cmap[value] = cs
	color := cs.Current()
	fmtedValue := fmt.Sprintf("%v", *value)
	t.Write(node, tm.Color(fmtedValue, color))
}

func (t *Terminal) HighlightConsidering(node []int) {
	t.Lock()
	defer t.Unlock()
	value := t.maze.Get(node)
	if value == nil {
		return
	}
	var cs ColorStack
	var ok bool
	if cs, ok = t.cmap[value]; !ok {
		cs = ColorStack{}
		t.cmap[value] = cs
	}
	color := cs.Current()
	fmtedValue := fmt.Sprintf("%v", *value)
	t.Highlight(node, fmtedValue, color, considering)
}

func (t *Terminal) HighlightFailed(node []int) {
	t.Lock()
	defer t.Unlock()
	value := t.maze.Get(node)
	if value == nil {
		return
	}
	cs, ok := t.cmap[value]
	if !ok {
		log.Panicln("Impossible failed highlight.")
	}
	color := cs.Current()
	fmtedValue := fmt.Sprintf("%v", *value)
	t.Highlight(node, fmtedValue, color, failed)
}

func (t *Terminal) HighlightSuccess(node []int) {
	t.Lock()
	defer t.Unlock()
	value := t.maze.Get(node)
	if value == nil {
		return
	}
	cs, ok := t.cmap[value]
	if !ok {
		log.Panicln("Impossible success highlight.")
	}
	color := cs.Current()
	fmtedValue := fmt.Sprintf("%v", *value)
	t.Highlight(node, fmtedValue, color, succeeded)
}

func (t *Terminal) ClearHighlight(node []int) {
	t.Lock()
	defer t.Unlock()
	value := t.maze.Get(node)
	if value == nil {
		return
	}
	cs, ok := t.cmap[value]
	if !ok {
		log.Panicln("Impossible clearing highlight.")
	}
	color := cs.Current()
	fmtedValue := fmt.Sprintf("%v", *value)
	t.Write(node, tm.Color(fmtedValue, color))
}

func MapXY(node []int) (int, int) {
	// X, Y
	return (node[1] + 1) * 2, (node[0] + 2)
}

