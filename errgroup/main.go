package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group

	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		return nil
	})

	g.Go(func() error {
		time.Sleep(10 * time.Second)
		return errors.New("failed to exec #2")
	})

	g.Go(func() error {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("occurs error:", err)
		return
	}
	fmt.Println("finished successfully")
}
