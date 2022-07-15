package infra

import (
	"fmt"
	"ring/application/domain/handlers"
)

type DisplayCountdown struct {
}

func (d *DisplayCountdown) Display(rt *handlers.RemainingTime) {
	fmt.Printf("\r%02d:%02d:%02d", rt.H, rt.M, rt.S)
}
