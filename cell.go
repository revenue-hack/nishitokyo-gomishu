package nisitokyo_gomishu

type Cell struct {
	Date
	DayOfWeek
	content string
}

type Celler interface {
	DayOfWeeker
	Dater
	Content() string
	SetContent(content string)
}

func (c *Cell) Content() string {
	return c.content
}

func (c *Cell) SetContent(content string) {
	c.content = content
}

func NewCell(index int, dateStr, content string) (Celler, error) {
	cell := new(Cell)
	y, m, _, err := parseDate(dateStr)
	if err != nil {
		return nil, err
	}
	d, detail, err := parseContent(content)
	if err != nil {
		return nil, err
	}
	cell.SetDay(d)
	cell.SetMonth(m)
	cell.SetYear(y)
	cell.SetContent(detail)
	cell.DayOfWeek = CalcDayOfWeek(index)

	return cell, nil
}
