package domain

import (
	"errors"
	"fmt"
	"time"
)

type SetTimer interface {
	Set(time.Duration) *time.Timer
}

const (
	DurationNotParsable string = "DurationNotParsable"
)

type SetTimerUseCase struct {
	SoundPlayer SoundPlayer
	Timer       SetTimer
}

type SetTimerCommand struct {
	Duration string
}

func (setTimerUseCase *SetTimerUseCase) HandleSetTimerCommand(command *SetTimerCommand) error {

	duration, error := time.ParseDuration(command.Duration)

	if error != nil {
		return errors.New(DurationNotParsable)
	}

	fmt.Println("It will ring in", command.Duration)

	timer := setTimerUseCase.Timer.Set(duration)

	<-timer.C
	setTimerUseCase.SoundPlayer.Play()

	return nil
}
