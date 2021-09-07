package rhelper

import "time"

const (
	SQLTimeFormat string = "2006-01-02 15:04:05.999"
)

func FromSQLDateJakarta(tm string) (t *time.Time) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return FromSQLDate(tm, loc)
}

func FromSQLDateUTC(tm string) (t *time.Time) {
	return FromSQLDate(tm, time.UTC)
}

func FromSQLDate(tm string, loc *time.Location) (t *time.Time) {
	tme, err := time.Parse(SQLTimeFormat, tm)
	if err != nil {
		t = nil
		return
	}

	if loc != nil {
		tme = tme.In(loc)
	}

	t = &tme
	return
}

func ToSQLDateUTC(tm time.Time) (dt string) {
	return ToSQLDate(tm, time.UTC)
}

func ToSQLDateJakarta(tm time.Time) (dt string) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return ToSQLDate(tm, loc)
}

func ToSQLDate(tm time.Time, loc *time.Location) (dt string) {
	if loc != nil {
		tm = tm.In(loc)
	}

	dt = tm.Format(SQLTimeFormat)
	return
}
