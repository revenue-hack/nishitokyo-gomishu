package nisitokyo_gomishu

const (
	Sunday DayOfWeek = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type DayOfWeek int

type DayOfWeeker interface {
	String() string
	Integer() int
}

func (d *DayOfWeek) Integer() int {
	return int(*d)
}

func (d *DayOfWeek) String() string {
	switch *d {
	case Sunday:
		return "日曜日"
	case Monday:
		return "月曜日"
	case Tuesday:
		return "火曜日"
	case Wednesday:
		return "水曜日"
	case Thursday:
		return "木曜日"
	case Friday:
		return "金曜日"
	case Saturday:
		return "土曜日"
	default:
		return ""
	}
}

func CalcDayOfWeek(index int) DayOfWeek {
	nr := index % 7
	return DayOfWeek(nr)
}
