# Project Structure

## `/cmd`

Main applications for the project are stored here. Thumb rules:

- Each application should have its own directory.
- Do NOT put a lot of code in the application directory. It is common to have a small `main` function that imports and
  invokes the code from the `internal` and `pkg` directories and nothing else.

## `/internal`

Private application and library code. This is the code you don't want others importing in their applications or
libraries. Note that this is enforced by Go.

## `/pkg`

Library code that's ok to use by external applications. Other projects will import these
libraries expecting them to work.

## `/scripts`

Scripts to perform various build, install, analysis, etc. tasks. These scripts keep the root level `Makefile` small and
clean.

## `/docs`

Design and user documents (in addition to your godoc generated documentation).

## `/assets`

Other assets to go along with your repository (images, logos, etc.).

# Prerequisites

- Install [Go 1.22.3](https://golang.org/doc/install)
- Install [Docker](https://docs.docker.com/get-docker/)
- For Windows only, install [Chocolatey](https://chocolatey.org/install) and
  then run `choco install make`

# ⚠️ Pre-commit ⚠️

Make sure you have already run `make pre-commit` before committing your code. This will ensure that your code is
properly formatted and passes all the tests.