package pitch

import (
	"math"
	"testing"
)

func TestAutocorrelationDetectsFrequency(t *testing.T) {
	sampleRate := 44100
	frequency := 440.0
	frameSize := 2048

	samples := make([]float64, frameSize)

	for i := 0; i < frameSize; i++ {
		samples[i] = math.Sin(2 * math.Pi * frequency * float64(i) / float64(sampleRate))
	}

	frame := Frame{
		Samples: samples,
		SampleRate: sampleRate,
	}

	result := Autocorrelation(frame)

	if math.Abs(result.Frequency-frequency) > 5 {
		t.Errorf("expected ~%.2f Hz, got %.2f Hz", frequency, result.Frequency)
	}
}