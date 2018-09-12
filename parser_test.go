package nisitokyo_gomishu

import (
	"testing"
)

func TestParseDate(t *testing.T) {
	str := "2018年10月"
	y, m, d, err := parseDate(str)
	if err != nil {
		t.Errorf("invalid string date err is %v\n", err)
	}

	if y != 2018 || m != 10 || d != 0 {
		t.Errorf("unmatch y is %d, m is %d, d is %d\n", y, m, d)
	}
}

func TestParseContent(t *testing.T) {
	for _, inputs := range []struct {
		content string
		wantD   int
		wantC   string
		wantE   bool
	}{
		{
			"7日びん・缶・ペットボトル・スプレー缶・ライター",
			7,
			"びん・缶・ペットボトル・スプレー缶・ライター",
			false,
		},
		{
			"16日古紙・古布類・プラスチック容器包装類",
			16,
			"古紙・古布類・プラスチック容器包装類",
			false,
		},
		{
			"25日休み",
			25,
			"休み",
			false,
		},
		{
			"28日可燃ごみ",
			28,
			"可燃ごみ",
			false,
		},
		{
			"",
			0,
			"",
			true,
		},
	} {
		d, c, err := parseContent(inputs.content)

		if d != inputs.wantD {
			t.Errorf("fail want is %d, result is %d\n", inputs.wantD, d)
		}
		if c != inputs.wantC {
			t.Errorf("fail want is %s, result is %s\n", inputs.wantC, c)
		}
		if (err != nil) != inputs.wantE {
			t.Errorf("fail want is %t, result is %v\n", inputs.wantE, err != nil)
		}
	}
}
