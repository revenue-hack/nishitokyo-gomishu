package nisitokyo_gomishu

import (
	"log"
	"strings"

	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
)

const (
	firstLink = "http://www.city.nishitokyo.lg.jp/kurasi/gomi_recycle/shushu/gomicalander/gomicalender_exel/tanashi.html"
	uaPc      = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.3a6"
)

var args = []string{
	"--headless",
	"--user-agent=" + uaPc,
	"--disable-gpu",
	"window-size=1280x800",
	"--no-sandbox",
	"--privileged",
	"--lang=ja",
}

type Day struct {
	date      string
	dayOfWeek string
	content   string
}

func Scrape() {
	//var wg sync.WaitGroup

	scrapeTable(func(query *goquery.Document) {
		query.Find(".table01").Each(func(_ int, s *goquery.Selection) {
			//wg.Add(1)
			s.Find("tr").Each(func(i int, s2 *goquery.Selection) {
				if i == 0 {
					fmt.Printf("day of week: %s\n", s.Find("th").Text())
				} else {
					fmt.Println("else")
					s2.Find(".top").Each(func(_ int, s2 *goquery.Selection) {
						text := s2.Text()
						if text != "" {
							fmt.Printf("content: %s\n", text)
						}
					})
				}
			})
			//wg.Done()
		})
	}, firstLink)
	//wg.Wait()
	fmt.Println("done")

}

func scrapeTable(parseDom func(query *goquery.Document), url string) {
	log.Println("parseTop start")
	driver := agouti.ChromeDriver()
	err := driver.Start()
	if err != nil {
		log.Fatalf("Fatal to start driver : %v", err)
	}
	defer driver.Stop()

	page, _ := driver.NewPage(agouti.Desired(agouti.Capabilities{"chromeOptions": map[string]interface{}{
		"args": args,
	}}))
	err = page.Navigate(url)

	content, err := page.HTML()
	if err != nil {
		log.Printf("Fatal to get html: %v\n", err)
	}

	doc, err := getDocument(content)
	if err != nil {
		log.Printf("Fatal to get document: %v\n", err)
	}

	parseDom(doc)
}

func getDocument(content string) (*goquery.Document, error) {
	r := strings.NewReader(content)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return doc, nil
}
