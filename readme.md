# Ranan

Ranan is a pitch detection and musical feedback engine written in Go.

The system captures microphone audio, detects the fundamental frequency, converts it to a musical note, and calculates cent deviation from equal temperament tuning.

---

## Features

- Live microphone audio capture
- Pitch detection via autocorrelation
- Frame buffering for stable pitch readings
- Frequency → MIDI note conversion
- Note name and octave indentification
- Cent deviation calculation from equal temperament
- Serial communication to Arduino for hardware LED tuner display

---

## Architecture
```
cmd/api/          — HTTP server
cmd/ranan/        — CLI feedback tool
internal/audio/   — microphone capture
internal/pitch/   — pitch detection and analysis
internal/music/   — frequency to note conversion
```

---

## Requirements

- Go 1.25+
- PortAudio

Install PortAudio on macOS:
```bash
brew install portaudio
brew install pkg-config
```

Install PortAudio on Linux:
```bash
sudo apt-get install portaudio19-dev
```

---

## Arduino Setup

The CLI can send pitch data to an Arduino over USB serial.

The Arduino receives messages in this format:
```
A4,0.79
```

Where the first part is the note name and octave, and the second part is the cent deviation.

To use this feature:
1. Connect Arduino UNO to your Mac via USB
2. Upload the Arduino firmware (see arduino/ directory)
3. Confirm the port name with: `ls /dev/cu.*`
4. Update the `arduinoPort` constant in `cmd/ranan/main.go`

---

## Running the CLI
```bash
go run cmd/ranan/main.go
```

Example output:
```
Listening... (Ctrl+C to stop)
Frequency: 440.00 Hz | Note: A4 | Cents: 0.00
Frequency: 441.20 Hz | Note: A4 | Cents: 4.71
Frequency: 438.80 Hz | Note: A4 | Cents: -2.23
```

Press Ctrl+C to stop.

---

## Running the API
```bash
go run cmd/api/main.go
```

The API runs on:
```
http://localhost:8080
```

Health check:
```
GET /health
```

Expected response:
```
OK
```

## Running Tests
```bash
go test ./...
```