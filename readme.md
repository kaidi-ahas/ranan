# Ranan

Ranan is a backend service for vocal pitch detection and singing accuracy analysis.

The goal of the project is to build an API capable of generating or accepting simple sheet music (a sequence of notes with durations), analyzing sung notes, and providing feedback about pitch accuracy.

---

## Features (planned)

- Generate simple sheet music sequences
- Accept simple sheet music input
- Pitch detection
- Frequency → musical note conversion
- Cent deviation calculation
- Singing accuracy feedback

---

## Running the API

Start the server:

`go run ./cmd/api`

The API will run on:

`http://localhost:8080`

Health check:

`GET /health`

Expected response:

`OK`