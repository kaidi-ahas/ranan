package pitch

import (
	"github.com/kaidi-ahas/ranan/internal/music"
)


type Analysis struct {
	Pitch Result
	Note music.Note
}

func Analyse(frame Frame) Analysis {
	pitch := Autocorrelation(frame)
	note := music.FromFrequency(pitch.Frequency)

	return Analysis{
		Pitch: pitch,
		Note: note,
	}
}