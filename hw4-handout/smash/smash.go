package smash

import (
	"bufio"
	"io"
	"sync"
)

type word string

func Smash(r io.Reader, smasher func(word) uint32) map[uint32]uint {
	m := make(map[uint32]uint)
	var mu sync.Mutex
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wor := word(scanner.Text())
		wg.Add(1)
		go func(wr word) {
			defer wg.Done()
			hash := smasher(wr)
			mu.Lock()
			m[hash]++
			mu.Unlock()
		}(wor)
	}

	wg.Wait()
	return m
}


// func Smash(r io.Reader, smasher func(word string) uint32) map[uint32]uint {
// 	m := make(map[uint32]uint)
// 	var mu sync.Mutex
// 	var wg sync.WaitGroup

// 	scanner := bufio.NewScanner(r)
// 	scanner.Split(bufio.ScanWords)

// 	for scanner.Scan() {
// 		word := scanner.Text()
// 		wg.Add(1)
// 		go func(w string) {
// 			defer wg.Done()
// 			hash := smasher(w)
// 			mu.Lock()
// 			m[hash]++
// 			mu.Unlock()
// 		}(word)
// 	}

// 	wg.Wait()

// 	return m
// }
