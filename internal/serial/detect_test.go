package serial

import (
	"os"
	"testing"
)

func TestDetectPort_FindsMatch(t *testing.T) {
	dir := t.TempDir()
	fakePort := dir + "/cu.usbmodem1234"
	f, err := os.Create(fakePort)
	if err != nil {
		t.Fatalf("failed to create fake port file: %v", err)
	}
	f.Close()

	port, err := DetectPort(dir + "/cu.usbmodem*")
	if err != nil {
		t.Fatalf("expected %s, got %s", fakePort, port)
	}
	if port != fakePort {
		t.Errorf("expected %s, got %s", fakePort, port)
	}
}

func TestDetectPort_NoMatch(t *testing.T) {
	dir := t.TempDir()

	_, err := DetectPort(dir + "/cu.usbmodem*")
	if err == nil {
		t.Error("expected error when no port found, got nil")
	}
}