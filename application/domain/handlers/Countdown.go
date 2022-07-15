package handlers

import (
	"time"
)

type DisplayCountdown interface {
	Display(*RemainingTime)
}

type Countdown struct {
	DisplayCountdown
}

type RemainingTime struct {
	T int
	H int
	M int
	S int
}

func (c *Countdown) Start(deadline time.Time, callback func()) {
	remainingTime := c.getRemainingTime(deadline)
	c.DisplayCountdown.Display(remainingTime)

	for range time.Tick(1 * time.Second) {
		remainingTime := c.getRemainingTime(deadline)
		c.DisplayCountdown.Display(remainingTime)

		if remainingTime.T <= 0 {
			callback()
			break
		}
	}
}

func (c *Countdown) getRemainingTime(t time.Time) *RemainingTime {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	hours := int(total / (60 * 60))
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return &RemainingTime{
		T: total,
		H: hours,
		M: minutes,
		S: seconds,
	}
}
