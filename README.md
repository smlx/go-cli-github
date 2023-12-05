# Go CLI Github

[![Release](https://github.com/smlx/go-cli-github/actions/workflows/release.yaml/badge.svg)](https://github.com/smlx/go-cli-github/actions/workflows/release.yaml)
[![Coverage](https://coveralls.io/repos/github/smlx/go-cli-github/badge.svg?branch=main)](https://coveralls.io/github/smlx/go-cli-github?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/smlx/go-cli-github)](https://goreportcard.com/report/github.com/smlx/go-cli-github)

This repo is an example with basic workflows for a Go CLI tool hosted on Github.
It adds basic PR building, dependabot integration, testing, coverage etc.

### How to use

1. Copy the contents of this repo into a new directory.
   ```bash
   git clone git@github.com:smlx/go-cli-github.git $PROJECT_DIR && rm -rf ./$PROJECT_DIR/.git
   ```
2. Update the `release` workflow branch from `main` to `foo` to disable it, commit all the files, and push to `main` on a new repo.
2. Update for your project, send a PR and merge it once green:
    * rename `deploy/go-cli-github` to `deploy/$YOUR_COMMAND`
    * update `deploy/$YOUR_COMMAND/Dockerfile`
    * rename `cmd/go-cli-github` to `cmd/$YOUR_COMMAND`
    * update `.goreleaser.yml` to build `cmd/$YOUR_COMMAND`
    * update the links at the top of the README
    * update the `build`, `release`, and `tag-to-release` workflows, replacing `go-cli-github` with `$YOUR_COMMAND`
    * update module in `go.mod` and in `cmd/$YOUR_COMMAND/*.go`, `internal/server/serve_test.go`
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
            * CodeQL
            * build
            * buildimage
            * commitlint
            * go-test
            * lint
    * Include administrators
5. Go to repository Settings > Code security and analysis, and enable:
    * Dependabot alerts
    * Dependabot security updates
    * Secret scanning
        * Push protection
6. When ready to release, rename the target branch in the release workflow from `foo` to `main`, and send a PR.
7. That's it.
