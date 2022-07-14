package infra

import "time"

type TimerSetter struct {
}

func (t *TimerSetter) Set(duration time.Duration) *time.Timer {
	timer := time.NewTimer(duration)
	return timer
}
