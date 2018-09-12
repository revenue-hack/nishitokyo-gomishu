package main

import (
	"fmt"

	"github.com/revenue-hack/nishitokyo-gomishu"
)

func main() {
	ch := make(chan nisitokyo_gomishu.Receive, 100)

	nisitokyo_gomishu.Scrape(ch)

	for {
		rec, ok := <-ch
		if !ok {
			break
		}
		if rec.Err != nil {
			fmt.Println("errrrrrrrrrrrr")
		}
		fmt.Printf("con: %s, year: %d\n", rec.Cell.Content(), rec.Cell.Year())
	}
	fmt.Println("done")
}
