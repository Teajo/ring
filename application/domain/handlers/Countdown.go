package handlers

import (
	"fmt"
	"time"
)

type Countdown struct {
}

type remainingTime struct {
	t int
	h int
	m int
	s int
}

func (c *Countdown) Start(deadline time.Time, callback func()) {
	remainingTime := c.getRemainingTime(deadline)
	c.displayCountdown(remainingTime)

	for range time.Tick(1 * time.Second) {
		remainingTime := c.getRemainingTime(deadline)
		c.displayCountdown(remainingTime)

		if remainingTime.t <= 0 {
			callback()
			break
		}
	}
}

func (c *Countdown) displayCountdown(ct *remainingTime) {
	fmt.Printf("\r%02d:%02d:%02d", ct.h, ct.m, ct.s)
}

func (c *Countdown) getRemainingTime(t time.Time) *remainingTime {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	hours := int(total / (60 * 60))
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return &remainingTime{
		t: total,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
