package pitch

import "testing"

func TestBufferAveragesFrequencies(t *testing.T) {
	buf := NewBuffer(3)

	buf.Add(Result{Frequency: 438.0})
	buf.Add(Result{Frequency: 440.0})
	buf.Add(Result{Frequency: 442.0})

	avg := buf.Average()

	if avg < 439.5 || avg > 440.5 {
		t.Errorf("expected average near 440.0, got %.2f", avg)
	}
}

func TestBufferReturnsZeroWhenEmpty(t *testing.T) {
	buf := NewBuffer(3)

	avg := buf.Average()

	if avg != 0.0 {
		t.Errorf("expected 0.0 for empty buffer, got %.2f", avg)
	}
}