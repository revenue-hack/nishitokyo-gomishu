package nisitokyo_gomishu

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var rep = regexp.MustCompile(`([0-9]+)年([0-9]+)月`)

func parseDate(dateStr string) (y, m, d int, err error) {
	result := rep.FindStringSubmatch(dateStr)

	y, err = strconv.Atoi(result[1])
	if err != nil {
		return
	}

	m, err = strconv.Atoi(result[2])
	if err != nil {
		return
	}
	return
}

var invalidArgError = func(detail interface{}) error {
	return fmt.Errorf("parse error, invalid arg is %v\n", detail)
}

var parseError = func(err error) error {
	return fmt.Errorf("parse error is %v\n", err)
}

func parseContent(content string) (int, string, error) {
	sepStrs := strings.Split(content, "日")

	if l := len(sepStrs); l != 2 || l < 0 {
		return 0, "", invalidArgError(sepStrs)
	}

	d, err := strconv.Atoi(sepStrs[0])
	if err != nil {
		return 0, "", parseError(err)
	}

	return d, sepStrs[1], nil
}
