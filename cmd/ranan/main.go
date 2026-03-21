package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kaidi-ahas/ranan/internal/audio"
	"github.com/kaidi-ahas/ranan/internal/pitch"
)

func main() {
	fmt.Println("Listening... (Ctrl+C to stop)")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <- stop:
			fmt.Println("\nStopped")
			return
		default:
				frame, err := audio.CaptureFrame()
				if err != nil {
					log.Printf("failed to capture audio: %v", err)
					continue
				}

				analysis := pitch.Analyse(frame)

				fmt.Printf("Frequency: %.2f Hz | Note: %s%d | Cents: %.2f\n",
					analysis.Pitch.Frequency,
					analysis.Note.Name, 
					analysis.Note.Octave,
					analysis.Note.Cents,
			)
		}
	}
}