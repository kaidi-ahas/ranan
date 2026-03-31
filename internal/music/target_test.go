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

func TestCentsBetween_Unison(t *testing.T) {
	cents := CentsBetween(440.0, 880.0)
	if math.Abs(cents-1200.0) > 0.01 {
		t.Errorf("expected 1200 cents for one octave, got %.2f", cents)
	}
}

func TestTuningStatus_Flat(t *testing.T) {
	status := TuningStatus(-15.0, 10.0)
	if status != "flat" {
		t.Errorf("expected flat, got %s", status)
	}
}

func TestTuningStatus_InTune(t *testing.T) {
	status := TuningStatus(3.0, 10.0)
	if status != "intune" {
		t.Errorf("expected intune, got %s", status)
	}
}

func TestTuningStatus_Sharp(t *testing.T) {
	status := TuningStatus(15.0, 10.0)
	if status != "sharp" {
		t.Errorf("expected sharp, got %s", status)
	}
}