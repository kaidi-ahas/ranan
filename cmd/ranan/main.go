package main

import (
	"fmt"
	"log"

	"github.com/kaidi-ahas/ranan/internal/audio"
	"github.com/kaidi-ahas/ranan/internal/pitch"
)

func main() {
	fmt.Println("Listening...")

	frame, err := audio.CaptureFrame()
	if err != nil {
		log.Fatalf("failed to capture audio: %v", err)
	}

	analysis := pitch.Analyse(frame)

	fmt.Printf("Frequency: %.2f Hz\n", analysis.Pitch.Frequency)
	fmt.Printf("Note:       %s%d\n", analysis.Note.Name, analysis.Note.Octave)
	fmt.Printf("Cents:      %.2f\n", analysis.Note.Cents)
}