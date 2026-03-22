package pitch

type Buffer struct {
	results []Result
	size    int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		results: make([]Result, 0, size),
		size:    size,
	}
}

func (b *Buffer) Add(result Result) {
	if len(b.results) >= b.size {
		b.results = b.results[1:]
	}
	b.results = append(b.results, result)
}

func (b *Buffer) Average() float64 {
	if len(b.results) == 0 {
		return 0.0
	}

	sum := 0.0
	for _, r := range b.results {
		sum += r.Frequency
	}

	return sum / float64(len(b.results))
}