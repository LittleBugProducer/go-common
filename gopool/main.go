package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/bytedance/gopkg/util/gopool"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		gopool.Go(func() {
			time.Sleep(time.Second * 5)
			fmt.Println("hello")
			wg.Done()
		})
	}
	wg.Wait()
}
