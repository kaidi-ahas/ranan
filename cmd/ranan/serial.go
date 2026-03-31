package main

import (
	"fmt"
	"log"

	"github.com/kaidi-ahas/ranan/internal/serial"
)


func OpenSerial(portFlag string) *serial.Port {
	device := portFlag

	if device == "" {
		detected, err := serial.DetectPort("/dev/cu.usbmodem*")
		if err != nil {
			log.Printf("warning: no Arduino port detected, serial disabled")
			return nil
		}
		fmt.Printf("Arduino detected on %s\n", detected)
		device = detected
	}

	port, err := serial.Open(device)
	if err != nil {
		log.Printf("warning: failed to open serial port %s, serial disabled: %v", device, err)
		return nil
	}

	return port
}