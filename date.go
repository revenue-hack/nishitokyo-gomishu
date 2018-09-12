package nisitokyo_gomishu

import "fmt"

type Date struct {
	year, month, day int
}

type Dater interface {
	Year() int
	Month() int
	Day() int
	SetYear(y int)
	SetMonth(m int)
	SetDay(day int)
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

func (d *Date) SetYear(y int) {
	d.year = y
}

func (d *Date) SetMonth(m int) {
	d.month = m
}

func (d *Date) SetDay(day int) {
	d.day = day
}

func NewDate(y, m, d int) (Dater, error) {
	if !(0 < m && m < 13) {
		return nil, fmt.Errorf("invalid attrs month is %d\n", m)
	}
	return &Date{y, m, d}, nil
}
