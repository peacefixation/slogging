# Slogger

## Setup

```bash
asdf plugin add golang
asdf install golang 1.24.2
asdf local golang 1.24.2
asdf plugin add golangci-lint
asdf install golangci-lint 1.64.2
asdf local golangci-lint 1.64.2
```

## Develop

```bash
make lint
make run
```

## Docs

https://www.dash0.com/guides/logging-in-go-with-slog

https://golangci-lint.run/docs/linters/configuration/#sloglint

