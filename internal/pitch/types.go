package pitch

type Frame struct {
	Samples []float64
	SampleRate int
}

type Result struct {
	Frequency float64
	Confidence float64
}