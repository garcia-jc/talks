package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Running in", runtime.Version())
	items, wg := []string{"Ale", "Tim", "Jim", "Fer"}, new(sync.WaitGroup)
	for i, item := range items {
		wg.Add(1)
		i, item := i, item // huh?
		go func() {
			fmt.Println("Yo, what's up ", item, "?")
			fmt.Println("Task ", i, "completed")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("work finished")
}
