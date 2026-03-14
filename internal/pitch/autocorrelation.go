package pitch

func Autocorrelation(frame Frame) Result {
	samples := frame.Samples
	sampleRate := float64(frame.SampleRate)

	maxCorr := 0.0
	bestLag := 0

	maxLag := len(samples) / 2

	for lag := 1; lag < maxLag; lag++ {
		corr := 0.0

		for i := 0; i < len(samples)-lag; i++ {
			corr += samples[i] * samples[i+lag]
		}

		if corr > maxCorr {
			maxCorr = corr
			bestLag = lag
		}
	}

	if bestLag == 0 {
		return Result{}
	}

	freq := sampleRate / float64(bestLag)

	return Result{
		Frequency: freq,
	}
}