GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' Makefile | column -t -s ':' |  sed -e 's/^/ /'

## info: Show project information
info:
	@echo "Git Tag:           ${GIT_TAG}"
	@echo "Git Commit:        ${GIT_COMMIT}"
	@echo "Git Tree State:    ${GIT_DIRTY}"

## generate: Generate code for the project
generate:
	@go generate ./...

## test: Run tests
test:
	@"$(CURDIR)/scripts/test.sh"

## lint: Run linters
lint:
	@"$(CURDIR)/scripts/lint.sh"

## gen: Generate CRUD code
gen:
	@"$(CURDIR)/scripts/gen.sh" -s $(service) -r $(resource)

## pre-commit: ⚠️ Run all required checks before commit
pre-commit:
	@make check
	@echo "--------------------------------------------------------------------------------"
	@make test

.NOTPARALLEL:

.PHONY: help info generate test lint gen pre-commit