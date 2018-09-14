package main

import (
	"fmt"

	"sync"

	"log"
	"os"

	"github.com/revenue-hack/nishitokyo-gomishu"
)

var scrapeError = func(rec nisitokyo_gomishu.Receive) {
	log.Printf("scrape fail receive is %v\n", rec.Cell)
}

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
				scrapeError(rec)
				os.Exit(1)
			}
			fmt.Printf("con: %s, month: %d, day: %d\n", rec.Cell.Content(), rec.Cell.Month(), rec.Cell.Day())
		}
	}()

	nisitokyo_gomishu.Scrape(ch)

	wg.Wait()

	fmt.Println("success")
}
