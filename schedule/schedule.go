package schedule

import "github.com/naufalfmm/moslem-salat-schedule/option"

type Schedule struct {
	Opt CommOpt
}

func (s Schedule) GetOption() option.Option {
	opt := s.Opt.ToOption()
	return &opt
}
