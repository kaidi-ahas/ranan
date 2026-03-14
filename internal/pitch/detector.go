package pitch

type Detector interface {
	Detect(frame Frame) (Result, error)
}