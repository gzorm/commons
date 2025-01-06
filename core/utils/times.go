package utils

import (
	"fmt"
	"time"

	"github.com/gzorm/commons/core/timex"
)

// An ElapsedTimer is a timer to track the elapsed time.
type ElapsedTimer struct {
	start time.Duration
}

// NewElapsedTimer returns an ElapsedTimer.
func NewElapsedTimer() *ElapsedTimer {
	return &ElapsedTimer{
		start: timex.Now(),
	}
}

// Duration returns the elapsed time.
func (et *ElapsedTimer) Duration() time.Duration {
	return timex.Since(et.start)
}

// Elapsed returns the string representation of elapsed time.
func (et *ElapsedTimer) Elapsed() string {
	return timex.Since(et.start).String()
}

// ElapsedMs returns the elapsed time of string on milliseconds.
func (et *ElapsedTimer) ElapsedMs() string {
	return fmt.Sprintf("%.1fms", float32(timex.Since(et.start))/float32(time.Millisecond))
}

// CurrentMicros returns the current microseconds.
func CurrentMicros() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// CurrentMillis returns the current milliseconds.
func CurrentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GetTodayStartAndEndTime
func GetTodayStartAndEndTime() (int64, int64) {
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := midnight.Add(time.Hour*24-time.Second).UnixNano() / 1e9 // 减去一秒得到当天的最后一秒

	startOfDay := midnight.Unix()
	return startOfDay, endOfDay
}
