package angle

import "gitlab.com/naufalfmm/moslem-salat-schedule/angle/consts"

func (d Angle) prepareConvertMinuteSecond() Angle {
	if d.second >= consts.TimeFormatConverter {
		d.second = d.second - consts.TimeFormatConverter
		d.minute = d.minute + consts.DecimalOne
	}

	if d.minute >= consts.TimeFormatConverter {
		d.minute = d.minute - consts.TimeFormatConverter
		d.degree = d.degree + consts.DecimalOne
	}

	return d
}
