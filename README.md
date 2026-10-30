# portfolio-site

A responsive personal portfolio and project case studies, served by a small Go application. Portfolio content is defined in a typed local Go data model; no external services, credentials, or JavaScript build step are needed.

## Local setup

Requires Go 1.26.2 or newer. Run:

```sh
go test ./...
go vet ./...
go run .
```

Open http://localhost:8080 to view the site. Stop the server with Ctrl-C.

## Architecture

The application uses the standard library's `net/http` server and `html/template` for escaped server-side rendering. `main.go` contains the `Project` and `Portfolio` data types and the local `portfolio` content. `newHandler` accepts a portfolio value, copies and sorts its projects, and renders the responsive semantic HTML page. The tests exercise rendered content, empty-state behavior, escaping, ordering, and unknown paths.

## Adding a project

Add a `Project` entry to the `Projects` slice in `main.go`, supplying a name, short description, repository URL, and technology tags. Keep descriptions concise and links pointed at the public project page. Run `go test ./...` before submitting the change.
