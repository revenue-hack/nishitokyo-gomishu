package nisitokyo_gomishu

import "fmt"

type Date struct {
	era              string
	year, month, day int
}

type Dater interface {
	Era() string
	Year() int
	Month() int
	Day() int
}

func (d *Date) Era() string {
	return d.era
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

func NewDate(era string, y, m, d int) (Dater, error) {
	if era == "" || !(0 < m && m < 13) {
		return nil, fmt.Errorf("invalid attrs era is %s\tmonth is %d\n", era, m)
	}
	return &Date{era, y, m, d}, nil
}
