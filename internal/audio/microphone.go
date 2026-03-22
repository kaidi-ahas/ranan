package audio

import (
	"github.com/gordonklaus/portaudio"
	"github.com/kaidi-ahas/ranan/internal/pitch"
)

const (
	sampleRate = 44100
	frameSize  = 2048
)

type Microphone struct {
	stream *portaudio.Stream
	samples []float32
}

func NewMicrophone() (*Microphone, error) {
	err := portaudio.Initialize()
	if err != nil {
		return nil, err
	}

	samples := make([]float32, frameSize)

	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, frameSize, samples)
	if err != nil {
		portaudio.Terminate()
		return nil, err
	}

	err = stream.Start()
	if err != nil {
		stream.Close()
		portaudio.Terminate()
		return nil, err
	}

	return &Microphone{
		stream: stream,
		samples: samples,
	}, nil
}

func (m *Microphone) Read() (pitch.Frame, error) {
	err := m.stream.Read()
	if err != nil {
		return pitch.Frame{}, err
	}

	converted := make([]float64, frameSize)
	for i, s := range m.samples {
		converted[i] = float64(s)
	}

	return pitch.Frame{
		Samples: converted,
		SampleRate: sampleRate,
	}, nil
}

func (m *Microphone) Close() {
	m.stream.Stop()
	m.stream.Close()
	portaudio.Terminate()
}