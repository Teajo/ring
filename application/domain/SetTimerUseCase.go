package domain

import (
	"errors"
	"ring/application/domain/handlers"
	"time"
)

type SetTimer interface {
	Set(time.Duration) *time.Timer
}

const (
	DurationNotParsable string = "Given duration is not parsable"
)

type SetTimerUseCase struct {
	SoundPlayer SoundPlayer
	Timer       SetTimer
	Countdown   *handlers.Countdown
}

type SetTimerCommand struct {
	Duration string
}

func (useCase *SetTimerUseCase) HandleSetTimerCommand(command *SetTimerCommand) error {

	duration, error := time.ParseDuration(command.Duration)

	if error != nil {
		return errors.New(DurationNotParsable)
	}

	deadline := time.Now().Add(duration)

	useCase.Countdown.Start(deadline, func() {
		useCase.SoundPlayer.Play()
	})

	return nil
}
