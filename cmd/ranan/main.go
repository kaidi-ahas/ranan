package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/kaidi-ahas/ranan/internal/audio"
	"github.com/kaidi-ahas/ranan/internal/music"
	"github.com/kaidi-ahas/ranan/internal/pitch"
	"github.com/kaidi-ahas/ranan/internal/serial"
)

const (
	bufferSize  = 5
	arduinoPort = "/dev/cu.usbmodem11301"
)

func main() {
	tunerMode := flag.Bool("tuner", false, "enable tuner mode")
	threshold := flag.Float64("threshold", 10.0, "in-tune threshold in cents")
	flag.Parse()

	mic, err := audio.NewMicrophone()
	if err != nil {
		log.Fatalf("failed to open microphone: %v", err)
	}
	defer mic.Close()

	port, err := serial.Open(arduinoPort)
	if err != nil {
		log.Printf("warning: Arduino not connected, serial disabled: %v", err)
		port = nil
	}
	if port != nil {
		defer port.Close()
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	buf := pitch.NewBuffer(bufferSize)

	if *tunerMode {
		runTunerMode(mic, buf, port, stop, *threshold)
	} else {
		runFreeMode(mic, buf, port, stop)
	}
}

func runFreeMode(mic *audio.Microphone, buf *pitch.Buffer, port *serial.Port, stop chan os.Signal) {
	fmt.Println("Listening... (Ctrl+C to stop)")

	for {
		select {
		case <-stop:
			fmt.Println("\nStopped.")
			return
		default:
			note, ok := readNote(mic, buf)
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
				if err := port.Send(note.Name, note.Octave, note.Cents); err != nil {
					log.Printf("failed to send to Arduino: %v", err)
				}
			}
		}
	}
}

func runTunerMode(mic *audio.Microphone, buf *pitch.Buffer, port *serial.Port, stop chan os.Signal, threshold float64) {
	inputCh := make(chan string)

	readInput := func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputCh <- scanner.Text()
		}
		close(inputCh)
	}

	go readInput()

	promptNext := make(chan struct{}, 1)
	promptNext <- struct{}{}

	for {
		select {
		case <-stop:
			fmt.Println("\nStopped.")
			return
		case <-promptNext:
			fmt.Print("Enter target note (e.g. A4) or type 'q' to quit: ")
			input, ok := <-inputCh
			if !ok {
				return
			}

			if strings.TrimSpace(input) == "q" {
				fmt.Println("Stopped.")
				return
			}

			input = strings.TrimSpace(input)
			name, octave, err := music.ParseNote(input)
			if err != nil {
				fmt.Printf("Invalid note: %v\n", err)
				promptNext <- struct{}{}
				continue
			}

			targetFreq, err := music.ToFrequency(name, octave)
			if err != nil {
				fmt.Printf("Invalid note: %v\n", err)
				promptNext <- struct{}{}
				continue
			}

			targetLabel := fmt.Sprintf("%s%d", name, octave)
			fmt.Printf("Tuning to %s (%.2f Hz)...\n", targetLabel, targetFreq)

			go func(label string, freq float64) {
				for {
					select {
					case <-stop:
						return
					default:
						note, ok := readNote(mic, buf)
						if !ok {
							continue
						}

						centsFromTarget := centsBetween(note.Frequency, freq)

						fmt.Printf("\rTarget: %s | Frequency: %.2f Hz | Cents: %+.2f   ",
							label,
							note.Frequency,
							centsFromTarget,
						)

						if port != nil {
							if err := port.Send(note.Name, note.Octave, centsFromTarget); err != nil {
								log.Printf("failed to send to Arduino: %v", err)
							}
						}

						if centsFromTarget >= -threshold && centsFromTarget <= threshold {
							fmt.Printf("\nNote %s is in tune. Press Enter to confirm and continue.", label)
							<-inputCh
							promptNext <- struct{}{}
							return
						}
					}
				}
			}(targetLabel, targetFreq)
		}
	}
}

func readNote(mic *audio.Microphone, buf *pitch.Buffer) (music.Note, bool) {
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

func centsBetween(detected, target float64) float64 {
	return 1200.0 * log2(target/detected)
}

func log2(x float64) float64 {
	return math.Log2(x)
}