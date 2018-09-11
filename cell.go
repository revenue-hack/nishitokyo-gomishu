package nisitokyo_gomishu

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Cell struct {
	Date
	content string
}

type Celler interface {
	Dater
	Content() string
}

func (c *Cell) Content() string {
	return c.content
}

/*
func (c *Cell) NewCell(dateStr string) (Celler, error) {
}
*/

func scrapeCell(i int, s *goquery.Selection) {
	if i == 0 {
		fmt.Printf("day of week: %s\n", s.Find("th").Text())
	} else {
		s.Find(".top").Each(func(_ int, s2 *goquery.Selection) {
			//text := s2.Text()
		})
	}
}
