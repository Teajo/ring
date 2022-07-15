package main

import (
	"fmt"
	"os"
	"ring/application/domain"

	application "ring/application/domain"
	handlers "ring/application/domain/handlers"

	infra "ring/application/infra"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No duration given, bye")
		return
	}

	args := os.Args[1:]
	duration := args[0]

	if duration == "" {
		fmt.Println("No duration given, bye")
		return
	}

	soundPlayer := &infra.SoundPlayer{}
	timer := &infra.TimerSetter{}
	displayCountdown := &infra.DisplayCountdown{}

	countDown := &handlers.Countdown{
		DisplayCountdown: displayCountdown,
	}

	timerUseCase := &application.SetTimerUseCase{
		SoundPlayer: soundPlayer,
		Timer:       timer,
		Countdown:   countDown,
	}

	err := timerUseCase.HandleSetTimerCommand(&domain.SetTimerCommand{
		Duration: duration,
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}
