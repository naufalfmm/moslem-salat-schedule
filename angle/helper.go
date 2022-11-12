package angle

import "gitlab.com/naufalfmm/moslem-salat-schedule/angle/consts"

func (d Angle) prepareConvertMinuteSecond() Angle {
	if d.second.GreaterThanOrEqual(consts.TimeFormatConverter) {
		d.second = d.second.Sub(consts.TimeFormatConverter)
		d.minute = d.minute.Add(consts.DecimalOne)
	}

	if d.minute.GreaterThanOrEqual(consts.TimeFormatConverter) {
		d.minute = d.minute.Sub(consts.TimeFormatConverter)
		d.degree = d.degree.Add(consts.DecimalOne)
	}

	return d
}
