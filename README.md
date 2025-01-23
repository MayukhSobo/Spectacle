# Spectacle 👓
An interactive TUI that powers your endeavour in ETCD

[![Build Status](https://github.com/MayukhSobo/Spectacle/actions/workflows/build.yml/badge.svg)](https://github.com/MayukhSobo/Spectacle/actions)
[![Lint Status](https://github.com/MayukhSobo/Spectacle/actions/workflows/linting.yaml/badge.svg)](https://github.com/MayukhSobo/Spectacle/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/MayukhSobo/Spectacle)](https://goreportcard.com/report/github.com/MayukhSobo/Spectacle)
[![GoDoc](https://godoc.org/github.com/MayukhSobo/Spectacle?status.svg)](https://godoc.org/github.com/MayukhSobo/Spectacle)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/MayukhSobo/Spectacle)](https://go.dev/)
[![Release](https://img.shields.io/github/v/release/MayukhSobo/Spectacle)](https://github.com/MayukhSobo/Spectacle/releases)

<p align="center">
  <img width="200" alt="Spectacle Logo" src="assets/logo.png">
</p>

## ✨ Features

- 🚀 Fast and lightweight TUI built with Bubble Tea
- 🔍 Intuitive navigation and data exploration
- 💾 Secure connection management
- 🎨 Beautiful terminal interface styled with Lipgloss
- 🔐 Support for TLS and authentication
- 🌐 Works with single node and cluster setups
- ⚡ Real-time updates and monitoring
- 🔄 Session persistence

## 🚀 Installation

### Using Go

```bash
go install github.com/MayukhSobo/Spectacle@latest
```

### From Source

```bash
git clone https://github.com/MayukhSobo/Spectacle.git
cd Spectacle
make build
```

## 📖 Usage

Simply run:

```bash
spectacle
```

### 🎮 Navigation
- `↑` `↓` Arrow keys to navigate
- `Enter` to select/expand
- `ESC` to go back
- `q` to quit
- `/` to search
- More keybindings available in-app

## 🛠 Development

### Prerequisites
- Go 1.21+
- Docker (for local development)
- Pre-commit hooks
- golangci-lint
- make

### Project Structure
```
Spectacle/
├── cmd/          # CLI commands
├── internal/     # Internal packages
│   └── app/     # TUI components
│       ├── home/       # Home screen
│       ├── savedconns/ # Saved connections
│       └── common/     # Shared components
├── logger/       # Logging utilities
└── dev/         # Development utilities
```

### Local Setup
1. Clone the repository
   ```bash
   git clone https://github.com/MayukhSobo/Spectacle.git
   cd Spectacle
   ```

2. Install dependencies
   ```bash
   go mod download
   ```

3. Install development tools
   ```bash
   make setup-dev
   ```

4. Start local ETCD
   ```bash
   docker-compose -f dev/docker-compose.yaml up -d
   ```

5. Run tests
   ```bash
   make test
   ```

### Code Quality
- Pre-commit hooks ensure code quality
- golangci-lint for static code analysis
- Unit tests for core functionality
- Integration tests with ETCD

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) for terminal styling
- [ETCD](https://etcd.io/) for the amazing distributed key-value store
