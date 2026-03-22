package serial

import (
	"fmt"

	"github.com/tarm/serial"
)


type Port struct {
	port *serial.Port
}

func Open(device string) (*Port, error) {
	config := &serial.Config{
		Name: device,
		Baud: 9600,
	}

	port, err := serial.OpenPort(config)
	if err != nil {
		return nil, fmt.Errorf("failed to open serial port %s: %w", device, err)
	}

	return &Port{port: port}, nil
}

func (p *Port) Send(note string, octave int, cents float64) error {
	message := fmt.Sprintf("%s%d, %.2f\n", note, octave, cents)

	_, err := p.port.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to write to serial port: %w", err)
	}

	return nil
}

func (p *Port) Close() error {
	return p.port.Close()
}