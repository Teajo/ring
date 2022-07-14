package domain

import "time"

type SetAlarm interface {
	Set(time.Time) *time.Timer
}

type SetAlarmUseCase struct {
	soundPlayer SoundPlayer
	timer       SetAlarm
}

type SetAlarmCommand struct {
	dateTime time.Time
}

func (setAlarmUseCase *SetAlarmUseCase) setAlarm(command *SetAlarmCommand) error {

	timer := setAlarmUseCase.timer.Set(command.dateTime)

	<-timer.C
	setAlarmUseCase.soundPlayer.Play()

	return nil
}
