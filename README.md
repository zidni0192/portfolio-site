# portfolio-site

A personal portfolio and project case studies, built as a small Go-served site.

## Local setup

Requires Go 1.26.2 or newer. Run:

```sh
go test ./...
go run .
```

## Architecture

The initial scaffold is a dependency-free Go command with a testable entry point. The first feature issue will establish the portfolio's page and content model.
