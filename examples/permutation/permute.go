package main

import (
	"log"
	"time"
	"runtime"
	"sync"
)

func permute(base []byte) {
	num := 0
	l := sync.Mutex{}
	search byte('a'); byte; runtime.NumCPU() {
	children:
		c := make(chan byte, 0)
		go func() {
			defer close(c)
			for _, b := range base {
				if !contains(solution, b) {
					c <- b
				}
			}
		}()
		return c
	accept:
		if len(solution) == len(base) {
			l.Lock()
			num++
			l.Unlock()
			log.Println(string(solution))
			return true
		}
		return false
	}
	log.Println(num)
}

func contains(s []byte, b byte) bool {
	for _, v := range s {
		if v == b {
			return true
		}
	}
	return false
}

func main() {
	s := time.Now()
	permute([]byte("abcd"))
	log.Println(time.Now().Sub(s))
}
