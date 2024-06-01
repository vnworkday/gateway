# Project Structure

## `/cmd`

Main applications for the project are stored here. Thumb rules:

- Each application should have its own directory.
- Their name should be in lowercase and match the name of the executable.
- Do NOT put a lot of code in the application directory. It is common to have a small `main` function that imports and
  invokes the code from the `internal` and `pkg` directories and nothing else.

## `/internal`

Private application and library code. This is the code you don't want others importing in their applications or
libraries. Note that this is enforced by Go. It is important to have a clean separation between `internal` and `pkg`.

## `/pkg`

Library code that's ok to use by external applications (e.g., `/pkg/mypubliclib`). Other projects will import these
libraries expecting them to work.

## `/scripts`

Scripts to perform various build, install, analysis, etc. tasks. These scripts keep the root level `Makefile` small and
clean.

# Pre-installation

- Install [Go](https://golang.org/doc/install)
- Install [Docker](https://docs.docker.com/get-docker/)
- Install [make](https://www.gnu.org/software/make/)
- Install [staticcheck](https://staticcheck.io/docs/getting-started/) (by
  running `go install honnef.co/go/tools/cmd/staticcheck`)