package music

import "math"

var noteNames = []string{
    "C", "C#", "D", "D#", "E", "F",
    "F#", "G", "G#", "A", "A#", "B",
}

type Note struct {
    Name      string
    Octave    int
    Frequency float64
    Cents     float64
}

func FromFrequency(freq float64) Note {
    midi := 69.0 + 12.0*math.Log2(freq/440.0)

    midiRounded := int(math.Round(midi))

    name := noteNames[midiRounded%12]
    octave := (midiRounded / 12) - 1

    refFreq := 440.0 * math.Pow(2.0, float64(midiRounded-69)/12.0)
    cents := 1200.0 * math.Log2(freq/refFreq)

    return Note{
        Name:      name,
        Octave:    octave,
        Frequency: freq,
        Cents:     cents,
    }
}