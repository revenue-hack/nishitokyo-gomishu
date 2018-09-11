package nisitokyo_gomishu

import (
	"reflect"
	"testing"
)

func TestParseDate(t *testing.T) {
	str := "平成29年10月"
	date, err := parseDate(str)
	if err != nil {
		t.Errorf("invalid string date err is %v\n", err)
	}

	reflect.DeepEqual(date, Date{"平成", 29, 10, 0})
}
