# Port Scanner

A lightweight TCP port scanner written in Go.

This tool scans a target host (IP or domain) across a single port or range of ports using concurrent workers, then prints discovered open ports.

## Features

- TCP connect scan using `net.DialTimeout`
- Configurable target host
- Supports single port (`80`) or range (`1-1024`)
- Configurable concurrency via worker threads
- Real-time output for open ports

## Requirements

- Go 1.20+ (module currently targets Go 1.24.2)

## Run

From the project root:

```bash
go run .
```

Default behavior:

- Host: `127.0.0.1`
- Ports: `1-1024`
- Threads: `100`

## CLI Options

| Flag | Description | Default |
|------|-------------|---------|
| `-host` | Target IP or domain (example: `scanme.nmap.org`) | `127.0.0.1` |
| `-ports` | Port range (`1-1024`) or single port (`80`) | `1-1024` |
| `-threads` | Number of concurrent workers | `100` |

## Examples

Scan default localhost range:

```bash
go run .
```

Scan a specific host and common ports:

```bash
go run . -host scanme.nmap.org -ports 20-100
```

Scan one port only:

```bash
go run . -host 192.168.1.10 -ports 22
```

Increase concurrency:

```bash
go run . -host example.com -ports 1-5000 -threads 300
```

## Example Output

```text
Scanning 127.0.0.1 on ports 1 to 1024 using 100 threads...
Port 22 is OPEN on 127.0.0.1
Port 80 is OPEN on 127.0.0.1
Scan complete! Open ports: [22 80]
```

## Notes

- Only scans TCP ports.
- Port discovery order may vary because scanning is concurrent.
- Very high thread counts can increase CPU/network load.

## Disclaimer

Use this tool only on systems you own or have explicit permission to test.