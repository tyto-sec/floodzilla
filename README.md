# Floodzilla - Go Program for DoS Attack Simulation

## Overview

Floodzilla is a Go-based tool designed to simulate various Denial-of-Service (DoS) attacks on a target system. This program provides multiple attack options, including TCP Connection Flood, SYN Flood, HTTP Flood, ICMP Flood, and Ping of Death. It also includes a verbose mode to display detailed request information during the execution of attacks.

**Disclaimer:** This program is intended for educational and testing purposes only. Misuse of this tool is illegal and can cause damage to systems or networks. Use it responsibly and only on systems you have permission to test.

## Features

- **TCP Connection Flood**: Overwhelm the target server by initiating multiple TCP connection requests.
- **SYN Flood**: Flood the target with SYN packets to exhaust available resources.
- **HTTP Flood**: Simulate an HTTP request flood on the target.
- **ICMP Flood**: Saturate the target with ICMP Echo (ping) requests.
- **Ping of Death**: Send malformed or oversized ping packets to crash or disrupt the target.
- **Verbose Mode**: Display detailed logs about each request.

## Installation

1. **Install Go**: Make sure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).
2. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/floodzilla.git
   cd floodzilla
   ```
3. **Build the project**:
   ```bash
   go build
   ```

## Usage

You can execute the binary with different flags depending on the type of attack you want to simulate. Each attack has specific parameters that you need to configure, such as the target address and port.

```bash
./floodzilla [options]
```

### Options

| Flag         | Description                                                  | Example Usage                                                                 |
|--------------|--------------------------------------------------------------|-------------------------------------------------------------------------------|
| `-tcp-flood` | Perform a TCP connection flood attack.                        | `./floodzilla -tcp-flood -target 192.168.0.1 -port 80 -verbose`               |
| `-syn-flood` | Perform a SYN flood attack.                                   | `./floodzilla -syn-flood -target 192.168.0.1 -port 80 -verbose`               |
| `-http-flood`| Perform an HTTP flood attack.                                 | `./floodzilla -http-flood -target http://192.168.0.1 -verbose`               |
| `-icmp-flood`| Perform an ICMP (ping) flood attack.                          | `./floodzilla -icmp-flood -target 192.168.12.1 -verbose`                      |
| `-ping-death`| Perform a Ping of Death attack.                               | `./floodzilla -ping-death -target 192.168.12.1 -verbose`                      |
| `-verbose`   | Enable verbose mode to show detailed request information.     | `./floodzilla -tcp-flood -target 192.168.0.1 -port 80 -verbose`               |
| `-target`    | Target IP or URL of the system to attack.                     | `./floodzilla -tcp-flood -target 192.168.0.1 -port 80`                        |
| `-port`      | Port to attack (default is 80).                               | `./floodzilla -tcp-flood -target 192.168.0.1 -port 80`                        |
| `-n`         | Number of concurrent goroutines (default is 10).              | `./floodzilla -tcp-flood -target 192.168.0.1 -port 80 -n 50`                  |

## License

This project is licensed under the GNU General Public LICENSE. See the [LICENSE](LICENSE) file for details.

## Warning

**Use this tool responsibly!** Conducting denial-of-service attacks without proper authorization is illegal and can lead to severe consequences. Only use this tool in a legal and ethical manner, such as in a lab environment or with explicit permission from the target system owner.

---

**Contributions**: Feel free to contribute to this project by submitting a pull request.