package main

import (
	"sync"
	"log"
	"time"
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

var vlarge = Maze{
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

type Set struct {
	maze Maze
	set  map[*int]bool
	sync.Mutex
}

func (s *Set) Contains(node []int) bool {
	k := s.maze.Get(node)
	s.Lock()
	_, ok := s.set[k]
	s.Unlock()
	return ok
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
	delete(s.set, k)
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

func solveMaze(maze Maze, entrance []int) {
	cycles := NewMap(maze)
	directions := [][]int{[]int{-1, 0}, // north
		[]int{1, 0},  // south
		[]int{0, 1},  // east
		[]int{0, -1}} // west
	search entrance; []int {
	children:
		cycle := cycles.Get(gid.ID, solution)
		cycle.Put(node)
		c := make(chan []int, 0)
		go func() {
			defer close(c)
			// var prev []int
			for _, d := range directions {
				child := []int{node[0] + d[0], node[1] + d[1]}
				if cycle.Contains(child) {
					continue
				}
				c <- child
				// cycle.Delete(prev)
				// prev = child
			}
			// We must wait until the final child is handled
			// before we remove it from the cycle detection map.
			// As such, let's give something we know the reject
			// block will pass on.
			c <- entrance
			cycle.Delete(node)
		}()
		return c
	reject:
		value := maze.Get(node)
		if value == nil || *value == 0 {
			return true
		}
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
			log.Println(solution)
			return true
		}
		return false
	}
}

func solveSimple() {
	entrance := []int{1, -1}
	solveMaze(simple, entrance)
}

func solveCycle() {
	entrance := []int{1, -1}
	solveMaze(cycle, entrance)
}

func solveLoop() {
	entrance := []int{1, -1}
	solveMaze(loop, entrance)
}

func solveLarge() {
	entrance := []int{3, -1}
	solveMaze(vlarge, entrance)
}

func testLarge() {
	
}

func main() {
	log.SetFlags(log.Lshortfile)
	s := time.Now()
	solveLarge()
	log.Println(time.Now().Sub(s))
}
