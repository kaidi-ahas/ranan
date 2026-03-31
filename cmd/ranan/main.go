package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kaidi-ahas/ranan/internal/audio"
	"github.com/kaidi-ahas/ranan/internal/pitch"
)

const bufferSize  = 5

func main() {
	tunerMode := flag.Bool("tuner", false, "enable tuner mode")
	threshold := flag.Float64("threshold", 10.0, "in-tune threshold in cents")
	portFlag := flag.String("port", "", "Arduino serial port (auto-detected if not set)")
	flag.Parse()

	mic, err := audio.NewMicrophone()
	if err != nil {
		log.Fatalf("failed to open microphone: %v", err)
	}
	defer mic.Close()

	port := OpenSerial(*portFlag)
	if port != nil {
		defer port.Close()
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	buf := pitch.NewBuffer(bufferSize)

	if *tunerMode {
		RunTunerMode(mic, buf, port, stop, *threshold)
	} else {
		RunFreeMode(mic, buf, port, stop)
	}
}