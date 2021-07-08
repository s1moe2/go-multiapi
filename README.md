# Multi API Servers

This repository contains a sample approach to a projects that defines multiple API servers with shared components.

## Structure

The root level contains the `domain` package, which keeps high-level domain related concepts (such as models and services interfaces).

`postgres` contains PostgreSQL implementations of the repository interfaces declared on the `domain` package.

`mock` contains mock implementations of the repository interfaces declared on the `domain` package.

`pkg` aggregates common utilities and internal libraries.

`cmd` hold multiple API implementations, each having their own `main` function, meaning they all can be started individually.


### Tests

Admin API unit tests: `go test ./cmd/admin/... -v`
