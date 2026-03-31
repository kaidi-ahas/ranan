package serial

import (
	"strings"
	"testing"
)

type fakeWriter struct {
	written string
}

func (f *fakeWriter) Write(p []byte) (int, error) {
	f.written += string(p)
	return len(p), nil
}

func newTestPort(w *fakeWriter) *Port {
	return &Port{writer: w}
}

func TestSend_Flat(t *testing.T) {
	w := &fakeWriter{}
	p := newTestPort(w)

	err := p.Send("A", 4, "flat")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(w.written, "A4,flat") {
		t.Errorf("expected message to contain A4,flat, got %s", w.written)
	}
}

func TestSent_InTune(t *testing.T) {
	w := &fakeWriter{}
	p := newTestPort(w)

	err := p.Send("A", 4, "intune")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(w.written, "A4,intune") {
		t.Errorf("expected message to contain A4,intune, got %s", w.written)
	}
}

func TestSent_Sharp(t *testing.T) {
	w := &fakeWriter{}
	p := newTestPort(w)

	err := p.Send("A", 4, "sharp")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(w.written, "A4,sharp") {
		t.Errorf("expected message to contain A4,sharp, got %s", w.written)
	}
}