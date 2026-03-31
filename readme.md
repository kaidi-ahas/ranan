# Ranan

Ranan is a pitch detection, musical feedback, and tuning engine written in Go.

The system captures microphone audio, detects the fundamental frequency, converts it to a musical note, calculates cent deviation from equal temperament tuning, and provides real-time feedback via CLI and Arduino LED indicator.

---

## Features

- Live microphone audio capture
- Pitch detection via autocorrelation
- Frame buffering for stable pitch readings
- Frequency → MIDI note conversion
- Note name and octave indentification
- Cent deviation calculation from equal temperament
- Tuner mode with sequential note input and in-place display
- Serial communication to Arduino for hardware LED tuner display
- Auto-detection of Arduino serial port


---

## Architecture
```
cmd/api/          — HTTP server
cmd/ranan/        — CLI tuner and feedback tool
internal/audio/   — microphone capture
internal/pitch/   — pitch detection and analysis
internal/music/   — frequency to note conversion and tuning logic
internal/serial   — Arduino serial communication
arduino/tuner/    - Arduino LED tuner firmware
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

Connect an Arduino UNO via USB and upload the firmware from `arduino/tuner/tuner.ino` using Arduino IDE.

The CLI auto-detects the Arduino port on startup:
```
Arduino detected on /dev/cu.usbmodem11301
```

To override the detected port use the `--port` flag:
```
go run cmd/ranan/main.go --port=/dev/usbmodem11301
```

If no Arduino is connected, serial is disabled automatically and the program continues without it.

The Arduino receives messages in this format:
```
A4,intune
A4,flat
A4,sharp
```

### LED behavior

| Status | LED |
|--------|-----|
| flat (cents < −threshold) | Left red LED |
| intune (within threshold) | Green LED |
| sharp (cents > +threshold) | Right red LED |

---

## Running the CLI

Free-running mode - detects and displays whatever note is played:
```bash
go run ./cmd/ranan/
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

## Tuner Mode

Tuner mode prompts for a target note and gives real-time feedback against it:
```
go run cmd/ranan/main.go --tuner
```

The program updates a single line in place:
```
Target: A4 | Frequency: 441.20 Hz | Cents: +4.71
```

When the pitch is within threshold, the program confirms:
```
Note A4 is in tune. Press Enter to confirm and continue.
```

To use a custom threshold in cents:
```
go run cmd/ranan/main.go --tuner --threshold=5
```

Default threshold is 10 cents. The Arduino LEDs reflect the same threshold as the CLI.

To quit, type `q` at the note prompt and press Enter.


## Running Tests
```bash
go test ./...
```