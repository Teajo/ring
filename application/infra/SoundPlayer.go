package infra

import (
	_ "embed"
	"io"
	"log"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type SoundPlayer struct {
	AlarmFilepath string
}

//go:embed sounds/alarm1.mp3
var sound []byte

func (s *SoundPlayer) Play() {

	r := io.NopCloser(strings.NewReader(string(sound)))

	streamer, format, err := mp3.Decode(r)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
