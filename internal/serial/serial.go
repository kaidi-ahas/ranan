package serial

import (
	"fmt"
	"io"

	"github.com/tarm/serial"
)


type Port struct {
	writer io.Writer
	closer io.Closer
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

	return &Port{writer: port, closer: port}, nil
}

func (p *Port) Send(note string, octave int, status string) error {
	message := fmt.Sprintf("%s%d,%s\n", note, octave, status)

	_, err := p.writer.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to write to serial port: %w", err)
	}

	return nil
}

func (p *Port) Close() error {
	return p.closer.Close()
}