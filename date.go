package moslemSalatSchedule

import "time"

func (i *impl) SetDate(date time.Time) MoslemSalatSchedule {
	i.option.SetDate(date)

	return i
}

func (i *impl) Now() MoslemSalatSchedule {
	i.option.Now()

	return i
}
