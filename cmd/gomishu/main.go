package main

import (
	"fmt"

	"sync"

	"github.com/revenue-hack/nishitokyo-gomishu"
)

func main() {
	ch := make(chan nisitokyo_gomishu.Receive, 100)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			rec, ok := <-ch
			if !ok {
				wg.Done()
				break
			}
			if rec.Err != nil {
				fmt.Println("errrrrrrrrrrrr")
			}
			fmt.Printf("con: %s, month: %d, day: %d\n", rec.Cell.Content(), rec.Cell.Month(), rec.Cell.Day())
		}
	}()

	nisitokyo_gomishu.Scrape(ch)

	wg.Wait()
	fmt.Println("done")
}
