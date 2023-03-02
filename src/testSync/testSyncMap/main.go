package main

import (
	"strconv"
	"sync"
)

var m = make(map[string]int)

func setMap(key string, value int) {
	m[key] = value
}

func getMap(key string) int {
	return m[key]
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			setMap(key, n)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
