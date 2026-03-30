package music

import (
	"math"
	"testing"
)

func TestParseNote_C4(t *testing.T) {
	name, octave, err := ParseNote("C4")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if name != "C" {
		t.Errorf("expected name C, got %s", name)
	}

	if octave != 4 {
		t.Errorf("expected octave 4, got %d", octave)
	}
}

func TestParseNote_Sharp(t *testing.T) {
	name, octave, err := ParseNote("C#4")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name != "C#" {
		t.Errorf("expected name C#, got %s", name)
	}

	if octave != 4 {
		t.Errorf("expected octave 4, got %d", octave)
	}
}

func TestParseNote_Invalid(t *testing.T) {
	_, _, err := ParseNote("X")
	if err == nil {
		t.Error("expected error for invalid note, got nil")
	}
}

func TestToFrequency_A4(t *testing.T) {
	freq, err := ToFrequency("A", 4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(freq-440) > 0.01 {
		t.Errorf("expected 440.0, got %f", freq)
	}
}

func TestToFrequency_C4(t *testing.T) {
	freq, err := ToFrequency("C", 4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(freq-261.63) > 0.01 {
		t.Errorf("expected ~261.63, got %f", freq)
	}
}

func TestToFrequency_UnknownNote(t *testing.T) {
	_, err := ToFrequency("X", 4)
	if err == nil {
		t.Fatal("expected error for unknown note, got nil")
	}
}



