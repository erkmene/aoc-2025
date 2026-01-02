package utils

import (
	"fmt"
	"time"
)

type TimeRecord struct {
	Start time.Time
	Laps  []time.Time
}

func NewTimeRecord() *TimeRecord {
	return &TimeRecord{Start: time.Now(), Laps: make([]time.Time, 0)}
}

func (t *TimeRecord) Lap() (time.Duration, time.Duration) {
	now := time.Now()
	totalDuration := now.Sub(t.Start)
	var lastLap time.Time
	if len(t.Laps) > 0 {
		lastLap = t.Laps[len(t.Laps)-1]
	} else {
		lastLap = t.Start
	}
	lastLapDuration := now.Sub(lastLap)
	t.Laps = append(t.Laps, now)
	return totalDuration, lastLapDuration
}

func (t *TimeRecord) String() string {
	return fmt.Sprintf("Start: %s, Laps: %v", t.Start.Format(time.RFC3339), t.Laps)
}

func (t *TimeRecord) LapAndLog(message string) {
	totalDuration, lastLapDuration := t.Lap()
	fmt.Printf("%s: last lap %s, total %s\n", message, lastLapDuration.String(), totalDuration.String())
}
