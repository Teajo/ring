package domain

import (
	"fmt"
	"time"
)

type SetTimer interface {
	Set(time.Duration) *time.Timer
}

type SetTimerUseCaseError string

const (
	DurationTooHigh SetTimerUseCaseError = "DurationTooHigh"
)

type SetTimerUseCase struct {
	SoundPlayer SoundPlayer
	Timer       SetTimer
}

type SetTimerCommand struct {
	Duration time.Duration
}

func (setTimerUseCase *SetTimerUseCase) HandleSetTimerCommand(command *SetTimerCommand) error {

	fmt.Println("It will ring in", command.Duration)

	timer := setTimerUseCase.Timer.Set(command.Duration)

	<-timer.C
	setTimerUseCase.SoundPlayer.Play()

	return nil
}
