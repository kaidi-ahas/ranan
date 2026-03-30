package music

import (
	"fmt"
	"math"
	"strconv"
)

func ParseNote(input string) (string, int, error) {
	if len(input) < 2 {
		return "", 0, fmt.Errorf("invalid note: %s", input)
	}

	var nameEnd int
	if len(input) > 2 && input[1] == '#' {
		nameEnd = 2
	} else {
		nameEnd = 1
	}

	name := input[:nameEnd]
	octaveStr := input[nameEnd:]

	octave, err := strconv.Atoi(octaveStr)
	if err != nil {
		return "", 0, fmt.Errorf("invalid octave in note: %s", input)
	}
	
	return name, octave, nil
}

func ToFrequency (name string, octave int) (float64, error) {

	index := -1
	
	for i, n := range noteNames {
		if n == name {
			index = i
			break
		}
	}

	if index == -1 {
		return 0, fmt.Errorf("unknown note name: %s", name)
	}

	midi := (octave+1)*12 + index
	freq := 440 * math.Pow(2.0, float64(midi-69)/12.0)

	return freq, nil
}