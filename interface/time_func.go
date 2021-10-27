current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

// func (t Time) Date() (year int, month Month, day int)
fmt.Println(current.Date()) // 2020 May 15

// func (t Time) Year() int
fmt.Println(current.Year()) // 2020

// func (t Time) Month() Month
fmt.Println(current.Month()) // May

// func (t Time) Day() int
fmt.Println(current.Day()) // 15

// func (t Time) Clock() (hour, min, sec int)
fmt.Println(current.Clock()) // 17 45 12

// func (t Time) Hour() int
fmt.Println(current.Hour()) //17

// func (t Time) Minute() int
fmt.Println(current.Minute()) // 45

// func (t Time) Second() int
fmt.Println(current.Second()) // 12

// func (t Time) Unix() int64
fmt.Println(current.Unix()) // 1589546712

// func (t Time) Weekday() Weekday
fmt.Println(current.Weekday()) // Friday

// func (t Time) YearDay() int
fmt.Println(current.YearDay()) // 136

// func (t Time) Format(layout string) string
current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
fmt.Println(current.Format("02-01-2006 15:04:05")) // 15-05-2020 17:45:12


//comparison
firstTime := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
secondTime := time.Date(2020, time.May, 15, 16, 45, 12, 0, time.Local)

// func (t Time) After(u Time) bool
// true если позже
fmt.Println(firstTime.After(secondTime)) // true

// func (t Time) Before(u Time) bool
// true если раньше
fmt.Println(firstTime.Before(secondTime)) // false

// func (t Time) Equal(u Time) bool
// true если равны
fmt.Println(firstTime.Equal(secondTime)) // false


//change structure
now := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

// func (t Time) Add(d Duration) Time
// изменяет дату в соответствии с параметром - "продолжительностью"
future := now.Add(time.Hour * 12) // перемещаемся на 12 часов вперед

// func (t Time) AddDate(years int, months int, days int) Time
// изменяет дату в соответствии с параметрами - количеством лет, месяцев и дней
past := now.AddDate(-1, -2, -3) // перемещаемся на 1 год, два месяца и 3 дня назад

// func (t Time) Sub(u Time) Duration
// вычисляет время, прошедшее между двумя датами
fmt.Println(future.Sub(past)) // 10332h0m0s

//time format constants
const (
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)