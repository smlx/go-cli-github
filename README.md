# Go CLI Github

[![Release](https://github.com/smlx/go-cli-github/actions/workflows/release.yaml/badge.svg)](https://github.com/smlx/go-cli-github/actions/workflows/release.yaml)
[![Coverage](https://coveralls.io/repos/github/smlx/go-cli-github/badge.svg?branch=main)](https://coveralls.io/github/smlx/go-cli-github?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/smlx/go-cli-github)](https://goreportcard.com/report/github.com/smlx/go-cli-github)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/smlx/go-cli-github/badge)](https://securityscorecards.dev/viewer/?uri=github.com/smlx/go-cli-github)

This repository is a template for a Go CLI tool or service.
It is quite opinionated about security and release engineering, but hopefully in a good way.

It comes pre-configured for integration with GitHub-specific features such as [Dependabot security tooling](https://docs.github.com/en/code-security/dependabot), [CodeQL](https://codeql.github.com/), and [branch protection](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches/about-protected-branches).
It also automatically builds and tests your code using [GitHub Actions](https://docs.github.com/en/actions).

## Features

* Use [GoReleaser](https://goreleaser.com/) to automatically build and create GitHub Releases and Docker images on merge to `main`.

    * This uses the [Conventional Commits Versioner](https://github.com/smlx/ccv) to automatically version each release.

* Lint your commit messages, and your Go, GitHub Action, and Dockerfile code.
* Test Pull Requests using `go test`.
* Build Docker images from Pull Requests for manual testing and review.
* Static code analysis using [CodeQL](https://codeql.github.com/) and [Go Report Card](https://goreportcard.com/).
* Test coverage analysis using [Coveralls](https://coveralls.io/).
* Security analysis using [OpenSSF](https://securityscorecards.dev).

## How to use

First set up the Github repo

1. Create a new empty Github repository.

Then push some code to main:

1. Install [gonew](https://go.dev/blog/gonew) and run this command, replacing the last argument with the name of your new module:

    ```bash
    gonew github.com/smlx/go-cli-github github.com/smlx/newproject
    ```

1. Create the git repo and push to `main` (which will become the default branch):

    ```bash
    cd newproject
    git init .
    git branch -M main
    git remote add origin git@github.com:smlx/newproject.git
    git add .
    git commit -a
    git push -u origin main
    ```

Then customize the code for your repository:

1. Check out a new branch to set up the repo `git checkout -b setup`

1. Update the code for your project:

    * rename `deploy/go-cli-github` to `deploy/$YOUR_COMMAND`
    * update `deploy/$YOUR_COMMAND/Dockerfile`
    * rename `cmd/go-cli-github` to `cmd/$YOUR_COMMAND`
    * update module in `cmd/$YOUR_COMMAND/*.go`, `internal/server/serve_test.go`
    * update `.goreleaser.yml` to build `cmd/$YOUR_COMMAND`
    * update the links at the top of the README
    * update the `build`, `release`, and `tag-to-release` workflows, replacing `go-cli-github` with `$YOUR_COMMAND`.

1. Commit and push:

    ```bash
    git commit -a
    git push -u origin setup
    ```
1. Open a PR, ensure all the actions are green, then merge the PR.

Configure the repository:

1. Go to repository Settings > General:

    * Disable wiki and projects (unless you plan to use them!)
    * Allow only merge commits for Pull Requests
    * Allow auto-merge
    * Automatically delete head branches

1. Go to repository Settings > Code security and analysis, and enable:

    * Dependabot alerts
    * Dependabot security updates
        * Secret scanning
            * Push protection
    * Private vulnerability reporting

1. Go to repository Settings > Actions > General:

    * Set Workflow permissions to "Read repository contents and package permissions"

1. Go to repository Settings > Rules > Rulesets, and import the `protect-default-branch.json` ruleset.

1. That's it.
