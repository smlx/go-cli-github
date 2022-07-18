# Go CLI Github

[![Release](https://github.com/smlx/go-cli-github/actions/workflows/release.yaml/badge.svg)](https://github.com/smlx/go-cli-github/actions/workflows/release.yaml)
[![Coverage](https://coveralls.io/repos/github/smlx/go-cli-github/badge.svg?branch=main)](https://coveralls.io/github/smlx/go-cli-github?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/smlx/go-cli-github)](https://goreportcard.com/report/github.com/smlx/go-cli-github)

This repo is an example with basic workflows for a Go CLI tool hosted on Github.
It adds basic PR building, dependabot integration, testing, coverage etc.

### How to use

1. Copy the contents of this repo into a new repo.
2. Rename `cmd/go-cli-github` and update the links at the top of the README. Send a PR for this change, and merge it once green.
3. Protect the main branch with requirements on lint/test as desired.
4. That's it.
