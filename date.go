package moslemSalatSchedule

import "time"

func (i *impl) SetDate(date time.Time) MoslemSalatSchedule {
	i.option.CalcOpt = i.option.CalcOpt.SetDate(date)

	return i
}

func (i *impl) Now() MoslemSalatSchedule {
	i.option.CalcOpt = i.option.CalcOpt.Now()

	return i
}
