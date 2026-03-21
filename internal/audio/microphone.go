package audio

import (
	"github.com/gordonklaus/portaudio"
	"github.com/kaidi-ahas/ranan/internal/pitch"
)

const (
	sampleRate = 44100
	frameSize  = 2048
)

func CaptureFrame() (pitch.Frame, error) {
	samples := make([]float32, frameSize)

	err := portaudio.Initialize()
	if err != nil {
		return pitch.Frame{}, err
	}
	defer portaudio.Terminate()

	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, frameSize, samples)
	if err != nil {
		return pitch.Frame{}, err
	}
	defer stream.Close()

	err = stream.Start()
	if err != nil {
		return pitch.Frame{}, err
	}

	err = stream.Read()
	if err != nil {
		return pitch.Frame{}, err
	}

	err = stream.Stop()
	if err != nil {
		return pitch.Frame{}, err
	}

	converted := make([]float64, frameSize)
	for i, s := range samples {
		converted[i] = float64(s)
	}

	return pitch.Frame{
		Samples:    converted,
		SampleRate: sampleRate,
	}, nil
}