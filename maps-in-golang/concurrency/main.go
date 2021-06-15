package main

import (
	"fmt"
	"sync"
)

type MapMutex struct {
	sync.RWMutex
	mp map[string]bool
}

var counter = struct {
	sync.RWMutex
	m map[string]int
}{m: make(map[string]int)}

func main() {
	fmt.Println("let's learn concrrency handling in maps in golnag")

	temp := MapMutex{
		mp: make(map[string]bool),
	}

	// To read from the counter, take the read lock
	counter.RLock()
	n := counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("some_key:", n)
	// or
	temp.RLock()
	val := temp.mp["some_key"]
	temp.RUnlock()
	fmt.Println(val)

	//To write to the counter, take the write lock
	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()
	//or
	temp.Lock()
	temp.mp["some_key"] = false
	temp.Unlock()

}
