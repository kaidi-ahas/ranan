package music

import (
	"testing"
)

func TestFromFrequency_A4(t *testing.T) {
	note := FromFrequency(440.0)

	if note.Name != "A" {
		t.Errorf("expected note name A, got %s", note.Name)
	}

	if note.Octave != 4 {
		t.Errorf("expected octave 4, got %d", note.Octave)
	}

	if note.Cents > 0.5 || note.Cents < -0.5 {
		t.Errorf("expected cents near 0, got %f", note.Cents)
	}
}

func TestFromFrequency_C4(t *testing.T) {
	note := FromFrequency(261.63)

	if note.Name != "C" {
		t.Errorf("expected note name C, got %s", note.Name)
	}

	if note.Octave != 4 {
		t.Errorf("expected octave 4, got %d", note.Octave)
	}

	if note.Cents > 0.5 || note.Cents < -0.5 {
		t.Errorf("expected cents near 0, got %f", note.Cents)
	}
}

func TestFromFrequency_SharpDeviation(t *testing.T) {
	note := FromFrequency(445.0)

	if note.Name != "A" {
		t.Errorf("expected note name A, got %s", note.Name)
	}

	if note.Octave != 4 {
		t.Errorf("expected octave 4, got %d", note.Octave)
	}

	if note.Cents < 18.0 || note.Cents > 20.0 {
		t.Errorf("expected cents near +19, got %f", note.Cents)
	}
}