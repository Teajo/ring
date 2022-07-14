package main

import (
	"fmt"
	"os"
	"ring/application/domain"
	application "ring/application/domain"
	infra "ring/application/infra"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No duration given, bye")
		return
	}

	args := os.Args[1:]

	if args[0] == "" {
		fmt.Println("No duration given, bye")
		return
	}

	duration, error := time.ParseDuration(args[0])

	if error != nil {
		fmt.Println("Failed to parse given duration")
		return
	}

	soundPlayer := &infra.SoundPlayer{AlarmFilepath: "sounds/alarm1.mp3"}
	timer := &infra.TimerSetter{}
	timerUseCase := &application.SetTimerUseCase{
		SoundPlayer: soundPlayer,
		Timer:       timer,
	}

	timerUseCase.HandleSetTimerCommand(&domain.SetTimerCommand{
		Duration: duration,
	})
}
