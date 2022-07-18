# Go CLI Github

[![Release](https://github.com/smlx/go-cli-github/actions/workflows/release.yaml/badge.svg)](https://github.com/smlx/go-cli-github/actions/workflows/release.yaml)
[![Coverage](https://coveralls.io/repos/github/smlx/go-cli-github/badge.svg?branch=main)](https://coveralls.io/github/smlx/go-cli-github?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/smlx/go-cli-github)](https://goreportcard.com/report/github.com/smlx/go-cli-github)

This repo is an example with basic workflows for a Go CLI tool hosted on Github.
It adds basic PR building, dependabot integration, testing, coverage etc.

### How to use

1. Copy the contents of this repo into a new directory. Update the `release` workflow branch from `main` to `foo` to disable it, commit all the files, and push to `main` on a new repo.
2. Rename `cmd/go-cli-github` and update the links at the top of the README. Send a PR for this change, and merge it once green.
3. Go to repository Settings > General:
  * Disable wiki and projects
  * Allow only merge commits for Pull Requests
  * Allow auto-merge
  * Automatically delete head branches
4. Go to repository Settings > Branches and add branch protection to `main`, and enable:
  * Require a PR before merging
    * Dismiss stale pull request approvals
  * Require status checks to pass before merging
    * Require branches to be up-to-date before merging.
    * Required status checks:
      * lint
      * commitlint
      * build
      * go-test
  * Include administrators
5. When ready to release, rename the target branch in the release workflow from `foo` to `main`, and send a PR.
6. That's it.
