# Moslem Salat Times

Package "Moslem Salat Times" calculate the salat times based on location coordinates

## Features
- Return the salat times by periodically options, such as daily, weekly, monthly (started by the specific date), and current monthly (started by the first day of the month)
- Able to return all the salat times or each salat
- Have 8 times by 5 salat times and 3 others, that are midnight, fajr, sunrise, dhuhr, asr, sunset, maghrib, and isha.
- Calculate based on options, that are coordinates, elevation, fajr and isha zenith options, mazhab, and higher latitude method.
- Have 11 fajr and isha zenith options, that are KEMENAG Indonesia, Egyptian General Authority of Survey, ISNA, Moonsighting Committee Worldwide, Muslim World League, Umm Al-Qura University, University of Islamic Sciences Karachi, JAKIM, MUIS, DIYANET, and UOIF

## Quick Start
Install the library by
```sh
go get github.com/naufalfmm/moslem-salat-times
```

```go
func main() {
    // Initialize
	mss, err := moslemSalatTimes.New(
		schedule.WithLatitudeLongitude(angle.NewDegreeFromFloat(-6.30286), angle.NewDegreeFromFloat(107.018512)),
		schedule.WithElevation(16.),
		schedule.WithSunZenith(sunZenithEnum.KEMENAG),
		schedule.WithMazhab(mazhabEnum.Standard),
		schedule.WithTimezone(time.Local),
		schedule.WithRoundingTimeOption(roundingTimeOptionEnum.Default),
		schedule.SetNow(),
		schedule.WithPeriodical(periodicalEnum.Weekly),
	)
	if err != nil {
		panic(err)
	}

    // Get all salat times
	allTimes, err := mss.AllTimes(mss.GetOption())
	if err != nil {
		panic(err)
	}

	for _, allTime := range allTimes {
		fmt.Println(allTime.Date.Format("02-01-2006"))
        fmt.Println("----------------------------------")
		for _, salatTime := range allTime.SalatTimes {
			fmt.Println(salatTime.Salat.Name(), salatTime.Time.Format("15:04:05"))
		}
		fmt.Println("----------------------------------")
		fmt.Println()
	}

    // Get dhuhr time by changing periodical to monthly and rounding to minute ceil
	dhuhrTimes, err := mss.Dhuhr(mss.GetOption().
        SetPeriodical(periodicalEnum.Monthly).
        SetRoundingTimeOption(roundingTimeOptionEnum.MinuteCeil))
	if err != nil {
		panic(err)
	}

	for _, dhuhrTime := range dhuhrTimes {
		fmt.Println(dhuhrTime.Salat.Name(), dhuhrTime.Date.Format("02-01-2006"), dhuhrTime.Time.Format("15:04:05"))
	}
}
```