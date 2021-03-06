package medtronic

import (
	"reflect"
	"testing"
	"time"
)

func TestBasalRates(t *testing.T) {
	cases := []struct {
		data  []byte
		sched BasalRateSchedule
	}{
		{
			[]byte{0x28, 0x00, 0x00, 0x30, 0x00, 0x12},
			[]BasalRate{
				{parseTD("00:00"), 1000},
				{parseTD("09:00"), 1200},
			},
		},
		{
			[]byte{0x20, 0x00, 0x00, 0x26, 0x00, 0x0D, 0x2C, 0x00, 0x13, 0x26, 0x00, 0x1C},
			[]BasalRate{
				{parseTD("00:00"), 800},
				{parseTD("06:30"), 950},
				{parseTD("09:30"), 1100},
				{parseTD("14:00"), 950},
			},
		},
		{
			[]byte{0x28, 0x00, 0x00, 0x28, 0x00, 0x06, 0x2C, 0x00, 0x0C, 0x30, 0x00, 0x14, 0x30, 0x00, 0x2C, 0x00, 0x00, 0x00},
			[]BasalRate{
				{parseTD("00:00"), 1000},
				{parseTD("03:00"), 1000},
				{parseTD("06:00"), 1100},
				{parseTD("10:00"), 1200},
				{parseTD("22:00"), 1200},
			},
		},
		{
			[]byte{0x00, 0x00, 0x00, 0x04, 0x00, 0x02, 0x08, 0x00, 0x04, 0x0C, 0x00, 0x06, 0x10, 0x00, 0x08, 0x14, 0x00, 0x0A, 0x18, 0x00, 0x0C, 0x1C, 0x00, 0x0E, 0x20, 0x00, 0x10, 0x24, 0x00, 0x12, 0x28, 0x00, 0x14, 0x2C, 0x00, 0x16, 0x30, 0x00, 0x18, 0x34, 0x00, 0x1A, 0x38, 0x00, 0x1C, 0x3C, 0x00, 0x1E, 0x40, 0x00, 0x20, 0x44, 0x00, 0x22, 0x48, 0x00, 0x24, 0x4C, 0x00, 0x26, 0x50, 0x00, 0x28, 0x54, 0x00, 0x2A, 0x58, 0x00, 0x2C, 0x5C, 0x00, 0x2E, 0x00, 0x00, 0x00},
			[]BasalRate{
				{parseTD("00:00"), 0},
				{parseTD("01:00"), 100},
				{parseTD("02:00"), 200},
				{parseTD("03:00"), 300},
				{parseTD("04:00"), 400},
				{parseTD("05:00"), 500},
				{parseTD("06:00"), 600},
				{parseTD("07:00"), 700},
				{parseTD("08:00"), 800},
				{parseTD("09:00"), 900},
				{parseTD("10:00"), 1000},
				{parseTD("11:00"), 1100},
				{parseTD("12:00"), 1200},
				{parseTD("13:00"), 1300},
				{parseTD("14:00"), 1400},
				{parseTD("15:00"), 1500},
				{parseTD("16:00"), 1600},
				{parseTD("17:00"), 1700},
				{parseTD("18:00"), 1800},
				{parseTD("19:00"), 1900},
				{parseTD("20:00"), 2000},
				{parseTD("21:00"), 2100},
				{parseTD("22:00"), 2200},
				{parseTD("23:00"), 2300},
			},
		},
	}
	for _, c := range cases {
		s := decodeBasalRateSchedule(c.data)
		if !reflect.DeepEqual(s, c.sched) {
			t.Errorf("decodeBasalRateSchedule(% X) == %+v, want %+v", c.data, s, c.sched)
		}
	}
}

func TestBasalRateAt(t *testing.T) {
	cases := []struct {
		sched  BasalRateSchedule
		at     time.Time
		target BasalRate
	}{
		{
			[]BasalRate{
				{parseTD("00:00"), 1000},
			},
			parseTime("2016-11-06T23:00:00"),
			BasalRate{parseTD("00:00"), 1000},
		},
	}
	for _, c := range cases {
		target := c.sched.BasalRateAt(c.at)
		if !reflect.DeepEqual(target, c.target) {
			t.Errorf("%v.BasalRateAt(%v) == %+v, want %+v", c.sched, c.at, target, c.target)
		}
	}
}

func parseTD(s string) TimeOfDay {
	t, err := parseTimeOfDay(s)
	if err != nil {
		panic(err)
	}
	return t
}
