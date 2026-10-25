# Pulse

**A lightweight concurrent URL checker written in Go.**

Pulse is a command-line tool that checks multiple URLs concurrently and reports their HTTP status codes and response times. It uses Go goroutines, a semaphore-based concurrency limit, HTTP timeouts, and [Cobra](https://github.com/spf13/cobra) for command-line interface management.

## Features

- **Concurrent URL checks** — check multiple URLs without waiting for each request to finish sequentially.
- **Concurrency limit** — limits the number of simultaneous checks to 5.
- **HTTP timeout** — each request has a 5-second timeout.
- **HTTP status reporting** — displays the status and status code returned by the server.
- **Response timing** — measures how long each request takes.
- **Multiple URL arguments** — check several URLs with one command.

## Prerequisites

Before setting up Pulse, make sure you have:

- [Go](https://go.dev/dl/) installed.
- [Git](https://git-scm.com/downloads) installed.
- A terminal or command prompt.

Verify your Go installation:

```bash
go version
```

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/abhegyn/Pulse.git
cd Pulse
```

> If the repository URL or capitalization differs, use the clone URL from the repository's GitHub page.

### 2. Download dependencies

From the project root, run:

```bash
go mod download
```

Go will download the dependencies declared in `go.mod`, including Cobra.

If dependencies have not yet been added to the module, see the official [Cobra installation instructions](https://github.com/spf13/cobra#installing).

### 3. Run Pulse

Run the application directly:

```bash
go run . check https://example.com
```

Or build an executable:

```bash
go build -o pulse .
```

On Linux/macOS:

```bash
./pulse check https://example.com
```

On Windows PowerShell:

```powershell
.\pulse.exe check https://example.com
```

## Usage

### Check a single URL

```bash
go run . check https://example.com
```

### Check multiple URLs

```bash
go run . check https://example.com https://go.dev https://github.com
```

Pulse checks the supplied URLs concurrently and prints each URL's HTTP status and response time.


## Example output

The exact output and response times will vary depending on network conditions and the target servers.

```text
for: https://example.com
status: 200 OK, status code: 200, time taken: 150ms

for: https://go.dev
status: 200 OK, status code: 200, time taken: 230ms
```

## How it works

Pulse uses several Go concepts to perform its checks efficiently:

- **Goroutines** — each URL check runs concurrently.
- **`sync.WaitGroup`** — waits for all launched checks to finish.
- **Buffered channels** — act as a semaphore to limit concurrent checks.
- **`net/http`** — sends HTTP requests and receives responses.
- **HTTP client timeouts** — prevent requests from waiting indefinitely.
- **Cobra** — provides the command and argument structure for the CLI.

## Development

Run the tests:

```bash
go test ./...
```

Run the tests with the race detector:

```bash
go test -race ./...
```

Pulse uses Go's built-in `testing` package. Tests can use `net/http/httptest` to create local HTTP servers, making it possible to test HTTP behavior without relying on external websites.

## Built With

- [Go](https://go.dev/) — programming language.
- [Cobra](https://github.com/spf13/cobra) — CLI framework.
- Go standard library packages: `net/http`, `sync`, and `time`.


## License

No license has been specified yet. Until a license is added to the repository, don't assume the code is available for unrestricted reuse.
