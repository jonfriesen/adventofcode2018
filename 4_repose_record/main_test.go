package main

import (
	"testing"
	"time"
)

func Test_getTimeStampFromLine(t *testing.T) {
	ts := getTimeStampFromLine("[1518-11-05 00:55] wakes up")

	if ts.Year() != 1518 || ts.Month() != time.November || ts.Day() != 5 {
		t.Errorf("Time expected %s but got %v", "1518-11-05 00:55", ts)
	}
}

func TestShift_sumSleep(t *testing.T) {
	shift := Shift{
		startTime: time.Date(1518, 11, 01, 00, 00, 00, 00, time.UTC),
		endTime:   time.Date(1518, 11, 01, 00, 59, 00, 00, time.UTC),
		events: []ShiftEvent{
			ShiftEvent{
				eTime: time.Date(1518, 11, 01, 00, 05, 00, 00, time.UTC),
				eType: Sleep,
			},
			ShiftEvent{
				eTime: time.Date(1518, 11, 01, 00, 25, 00, 00, time.UTC),
				eType: Awake,
			},
			ShiftEvent{
				eTime: time.Date(1518, 11, 01, 00, 30, 00, 00, time.UTC),
				eType: Sleep,
			},
			ShiftEvent{
				eTime: time.Date(1518, 11, 01, 00, 55, 00, 00, time.UTC),
				eType: Awake,
			},
		},
	}

	sum := shift.sumSleep()

	if sum != 45 {
		t.Error("Expected 45, but got ", sum)
	}

	shift = Shift{
		startTime: time.Date(1518, 11, 01, 00, 00, 00, 00, time.UTC),
		endTime:   time.Date(1518, 11, 01, 00, 59, 00, 00, time.UTC),
		events: []ShiftEvent{
			ShiftEvent{
				eTime: time.Date(1518, 11, 01, 00, 05, 00, 00, time.UTC),
				eType: Sleep,
			},
			ShiftEvent{
				eTime: time.Date(1518, 11, 01, 00, 25, 00, 00, time.UTC),
				eType: Awake,
			},
			ShiftEvent{
				eTime: time.Date(1518, 11, 01, 00, 30, 00, 00, time.UTC),
				eType: Sleep,
			},
		},
	}

	sum = shift.sumSleep()

	if sum != 49 {
		t.Error("Expected 49, but got ", sum)
	}

}
