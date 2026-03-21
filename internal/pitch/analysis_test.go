package pitch

import (
	"math"
	"testing"
)

func TestAnalyseDetectsNoteFromSineWave(t *testing.T) {
	sampleRate := 44100
	frequency := 440.0
	frameSize := 2048

	samples := make([]float64, frameSize)
	for i:= 0; i < frameSize; i++ {
		samples[i] = math.Sin(2*math.Pi*frequency*float64(i) / float64(sampleRate))
	}

	frame := Frame{
		Samples: samples,
		SampleRate: sampleRate,
	}

	analysis := Analyse(frame)

	if math.Abs(analysis.Pitch.Frequency-frequency) > 5 {
		t.Errorf("expected frequency ~440 Hz, got %.2f", analysis.Pitch.Frequency)
	}

	if analysis.Note.Name != "A" {
		t.Errorf("expected note A, got %s", analysis.Note.Name)
	}

	if analysis.Note.Octave != 4 {
		t.Errorf("expected octave 4, got %d", analysis.Note.Octave)
	}
}