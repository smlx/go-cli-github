# Go CLI GitHub

[![Go Reference](https://pkg.go.dev/badge/github.com/smlx/go-cli-github.svg)](https://pkg.go.dev/github.com/smlx/go-cli-github)
[![Release](https://github.com/smlx/go-cli-github/actions/workflows/release.yaml/badge.svg)](https://github.com/smlx/go-cli-github/actions/workflows/release.yaml)
[![coverage](https://raw.githubusercontent.com/smlx/go-cli-github/badges/.badges/main/coverage.svg)](https://github.com/smlx/go-cli-github/actions/workflows/coverage.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/smlx/go-cli-github)](https://goreportcard.com/report/github.com/smlx/go-cli-github)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/smlx/go-cli-github/badge)](https://securityscorecards.dev/viewer/?uri=github.com/smlx/go-cli-github)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/8168/badge)](https://www.bestpractices.dev/projects/8168)

This repository is a template for a Go CLI tool or service.
It is quite opinionated about security and release engineering, but hopefully in a good way.

It comes pre-configured for integration with GitHub-specific features such as [Dependabot security tooling](https://docs.github.com/en/code-security/dependabot), [CodeQL](https://codeql.github.com/), and [branch protection](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches/about-protected-branches).
It also automatically builds and tests your code using [GitHub Actions](https://docs.github.com/en/actions).

## Features

* Use [GoReleaser](https://goreleaser.com/) to automatically build and create GitHub Releases and container images on merge to `main`.

    * This uses the [Conventional Commits Versioner](https://github.com/smlx/ccv) to automatically version each release.

* Lint your commit messages, Go code, GitHub Actions, and Dockerfiles.
* Test Pull Requests using `go test`.
* Build container images from Pull Requests and push them to the GitHub container registry for manual testing and review.
* Static code analysis using [CodeQL](https://codeql.github.com/) and [Go Report Card](https://goreportcard.com/).
* Coverage analysis using the [go-test-coverage action](https://github.com/vladopajic/go-test-coverage).
* Security analysis using [OpenSSF](https://securityscorecards.dev).
* Signed binary and container release artifacts using [artifact attestations](https://docs.github.com/en/actions/security-guides/using-artifact-attestations-to-establish-provenance-for-builds).
* SBOM generation for both release artifacts and container images, with image SBOMs pushed to the container registry.

## How to use

First set up the GitHub repo

1. Create a new empty GitHub repository.

Then push some code to main:

1. Install [gonew](https://go.dev/blog/gonew) and run this command, replacing the last argument with the name of your new module:

    ```bash
    gonew github.com/smlx/go-cli-github@main github.com/smlx/newproject
    ```

1. Create the git repo and push to `main` (which will become the default branch):

    ```bash
    cd newproject
    git init .
    git branch -M main
    git remote add origin git@github.com:smlx/newproject.git
    git add .
    git commit -am 'chore: create repository from template'
    git push -u origin main
    ```

1. Create the `badges` branch for storing the README coverage badge.

    ```bash
    git checkout --orphan badges
    git rm -rf .
    rm -f .gitignore
    echo 'This branch exists only to store the coverage badge in the README on `main`.' > README.md
    git add README.md
    git commit -m 'chore: initialize the badges branch'
    git push origin badges
    ```

Then customize the code for your repository:

1. Check out a new branch to set up the repo `git checkout -b setup main`

1. Update the code for your project:

    * rename `cmd/go-cli-github` to `cmd/$YOUR_COMMAND`
    * update `.github/workflows/build.yaml`, replacing `go-cli-github` with `$YOUR_COMMAND`.
    * update `.goreleaser.yaml` to build `cmd/$YOUR_COMMAND`
    * update the links at the top of `README.md`
    * update the contact email in `SECURITY.md`

1. Commit and push:

    ```bash
    git add .
    git commit -am 'chore: update template for new project'
    git push -u origin setup
    ```
1. Open a PR, wait until all the checks go green, then merge the PR.

Configure the repository:

1. Go to repository Settings > General:

    1. Releases

        * Enable release immutability

    1. Features

        * Disable wiki and projects (unless you plan to use them!)

    1. Pull Requests

        * Allow merge commits only for Pull Requests
        * Allow auto-merge
        * Automatically delete head branches

1. Go to repository Settings > Advanced Security, and enable:

    * Private vulnerability reporting

    * Dependabot

        * Dependabot alerts
        * Dependabot security updates
        * Grouped security updates
        * Dependabot on Actions runners

    * Code Scanning

        * CodeQL analysis > Set up > Default

    * Secret Protection

        * Push protection

1. Go to repository Settings > Rules > Rulesets, and import the `protect-default-branch.json` ruleset.

That's it.

## How to contribute

Issues are welcome.

PRs are also welcome, but keep in mind that this is a very opinionated template, so not all changes will be accepted.
PRs also need to ensure that test coverage remains high, and best practices are followed.
