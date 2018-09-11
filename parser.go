package nisitokyo_gomishu

import (
	"regexp"
	"strconv"
)

var rep = regexp.MustCompile(`(\W+)([0-9]+)年([0-9]+)月`)

func parseDate(dateStr string) (Dater, error) {
	result := rep.FindStringSubmatch(dateStr)

	era := result[1]
	y, err := strconv.Atoi(result[2])
	if err != nil {
		return nil, err
	}

	m, err := strconv.Atoi(result[3])
	if err != nil {
		return nil, err
	}

	return NewDate(era, y, m, 0)
}
