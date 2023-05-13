package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go incrementar(&wg, &mutex)
	}
	wg.Wait()
	fmt.Println("Valor final del counter:", counter)
}

func incrementar(wg *sync.WaitGroup, m *sync.Mutex) {
	for j := 0; j < 100; j++ {
		m.Lock()
		counter++
		m.Unlock()
	}
	wg.Done()

}
