package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kaidi-ahas/ranan/internal/audio"
	"github.com/kaidi-ahas/ranan/internal/music"
	"github.com/kaidi-ahas/ranan/internal/pitch"
	"github.com/kaidi-ahas/ranan/internal/serial"
)

func RunFreeMode(mic *audio.Microphone, buf *pitch.Buffer, port *serial.Port, stop chan os.Signal) {
	fmt.Println("Listening... (Ctrl+C to stop)")

	for {
		select {
		case <-stop:
			fmt.Println("\nStopped.")
			return
		default:
			note, ok := ReadNote(mic, buf)
			if !ok {
				continue
			}

			fmt.Printf("Frequency: %.2f Hz | Note: %s%d | Cents: %.2f\n",
				note.Frequency,
				note.Name,
				note.Octave,
				note.Cents,
			)

			if port != nil {
				status := music.TuningStatus(note.Cents, 10.0)
				if err := port.Send(note.Name, note.Octave, status); err != nil {
					log.Printf("failed to send to Arduino: %v", err)
				}
			}
		}
	}
}

func ReadNote(mic *audio.Microphone, buf *pitch.Buffer) (music.Note, bool) {
	frame, err := mic.Read()
	if err != nil {
		log.Printf("failed to read audio: %v", err)
		return music.Note{}, false
	}

	result := pitch.Autocorrelation(frame)
	if result.Frequency == 0.0 {
		return music.Note{}, false
	}

	buf.Add(result)
	avgFreq := buf.Average()
	if avgFreq == 0.0 {
		return music.Note{}, false
	}

	return music.FromFrequency(avgFreq), true
}