package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kaidi-ahas/ranan/internal/audio"
	"github.com/kaidi-ahas/ranan/internal/music"
	"github.com/kaidi-ahas/ranan/internal/pitch"
	"github.com/kaidi-ahas/ranan/internal/serial"
)

const (
	bufferSize = 5
	arduinoPort = "/dev/cu.usbmodem11301"
)

func main() {
	fmt.Println("Listening... (Ctrl+C to stop)")

	mic, err := audio.NewMicrophone()
	if err != nil {
		log.Fatalf("failed to open microphone: %v", err)
	}
	defer mic.Close()

	port, err := serial.Open(arduinoPort)
	if err != nil {
		log.Fatalf("failed to open serial port: %v", err)
	}
	defer port.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	buf := pitch.NewBuffer(bufferSize)

	for {
		select {
		case <-stop:
			fmt.Println("\nStopped")
			return
		default:
			frame, err := mic.Read()
			if err != nil {
				log.Printf("failed to read audio: %v", err)
				continue
			}

			result := pitch.Autocorrelation(frame)
			if result.Frequency == 0.0 {
				continue
			}

			buf.Add(result)

			avgFreq := buf.Average()
			if avgFreq == 0.0 {
				continue
			}

			note := music.FromFrequency(avgFreq)

			fmt.Printf("Frequency: %.2f Hz | Note: %s%d | Cents: %.2f\n",
				note.Frequency,
				note.Name,
				note.Octave,
				note.Cents,
			)

			err = port.Send(note.Name, note.Octave, note.Cents)
			if err != nil {
				log.Printf("failed to send to Arduino: %v", err)
			}
		}
	}
}
