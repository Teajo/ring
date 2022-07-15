package handlers

import (
	"fmt"
	"time"
)

type DisplayCountdown interface {
	display(*RemainingTime)
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
	c.displayCountdown(remainingTime)

	for range time.Tick(1 * time.Second) {
		remainingTime := c.getRemainingTime(deadline)
		c.DisplayCountdown.display(remainingTime)

		if remainingTime.T <= 0 {
			callback()
			break
		}
	}
}

func (c *Countdown) displayCountdown(ct *RemainingTime) {
	fmt.Printf("\r%02d:%02d:%02d", ct.H, ct.M, ct.S)
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
