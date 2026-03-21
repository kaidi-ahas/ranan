package main

import (
	"fmt"
	"math"

	"github.com/kaidi-ahas/ranan/internal/pitch"
)

func main() {
	sampleRate := 44100
	frequency := 440.0
	frameSize := 2048

	samples := make([]float64, frameSize)
	for i := 0; i < frameSize; i++ {
		samples[i] = math.Sin(2 * math.Pi * frequency * float64(i) / float64(sampleRate))
	}

	frame := pitch.Frame{
		Samples:    samples,
		SampleRate: sampleRate,
	}

	analysis := pitch.Analyse(frame)

	fmt.Printf("Frequency: %.2f Hz\n", analysis.Pitch.Frequency)
	fmt.Printf("Note:      %s%d\n", analysis.Note.Name, analysis.Note.Octave)
	fmt.Printf("Cents:     %.2f\n", analysis.Note.Cents)
}