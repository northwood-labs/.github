# something outside a snippet

# @config-manager:start testing
# This should definitely never show up in the downstream file.
# @config-manager:end testing

#-------------------------------------------------------------------------------
# @config-manager:start includes
include ./standard.mk
# @config-manager:end includes

# @config-manager:start binary_name
#===============================================================================
# After bootstrapping, remove the `@config-manager` markers from around these
# tasks, then fill them out as appropriate for your project.
#===============================================================================
BINARY_NAME=name-of-binary
# @config-manager:end binary_name

#-------------------------------------------------------------------------------
# @config-manager:start go_globals
# Running `make` will show the list of subcommands that will run.

# go install golang.org/dl/go{VERSION}@latest
# go{VERSION} download
GO_VER=1.25
GOBIN=$(shell ./find-go-bin.sh)
# @config-manager:end go_globals

#-------------------------------------------------------------------------------
# @config-manager:start standard_installation
# Installation

.PHONY: install-tools
## install-tools: [tools]* Install/upgrade the required dependencies.
install-tools: install-tools-mac install-tools-go install-tools-py

.PHONY: install-tools-py
## install-tools-py: [tools] Install/upgrade the required Go packages.
install-tools-py:
	@ $(HEADER) "=====> Installing Python tools..."
	poetry lock && poetry install

.PHONY: install-tools-go
## install-tools-go: [tools] Install/upgrade the required Go packages.
install-tools-go:
	@ $(HEADER) "=====> Installing Go tools..."
	$(GO) get -modfile=go.tools.mod tool

.PHONY: install-tools-mac
## install-tools-mac: [tools] Install/upgrade the required tools for macOS.
install-tools-mac:
	@ $(HEADER) "=====> Installing required packages for macOS (Homebrew)..."
	brew bundle install --upgrade

	@ $(HEADER) "=====> Installing Chag (may require password)..."
	curl -sSLf https://raw.githubusercontent.com/mtdowling/chag/master/install.sh | sudo bash

	@ $(BORDER) "To update to the latest versions, run:" "  make install-tools-mac"

.PHONY: install-hooks
## install-hooks: [tools]* Install/upgrade the Git hooks used for ensuring consistency.
install-hooks:
	@ $(HEADER) "=====> Installing Git hooks..."
	cp -vf .githooks/commit-msg.sh .git/hooks/commit-msg
	chmod +x .git/hooks/*
	pre-commit install

	@ $(BORDER) "Learn more about 'pre-commit' at:" "  https://pre-commit.com" " " "Learn more about 'gommit' at:" "  https://github.com/antham/gommit"
# @config-manager:end standard_installation

#-------------------------------------------------------------------------------
# @config-manager:start go_compile
# Go Compile

.PHONY: tidy
## tidy: [build] Updates go.mod and downloads dependencies.
tidy:
	@ $(HEADER) "=====> Tidy and download the Go dependencies..."
	$(GO) mod tidy -go=$(GO_VER) -v
	$(GO) mod tidy -go=$(GO_VER) -v -modfile=go.tools.mod

.PHONY: godeps
## godeps: [build] Updates go.mod and downloads dependencies.
godeps:
	@ $(HEADER) "=====> Upgrade the minor versions of Go dependencies..."
	find . -type f -name "go.mod" | xargs -I% dirname "%" | xargs -I@ bash -c 'cd "@" && $(GO) mod tidy -go=$(GO_VER) && $(GO) get -u -t -v ./...'
	find . -type f -name "go.tools.mod" | xargs -I% dirname "%" | xargs -I@ bash -c 'cd "@" && $(GO) mod tidy -go=$(GO_VER) -modfile=go.tools.mod && $(GO) get -u -t -v -modfile=go.tools.mod ./...'

	@ echo ""
	@ $(YELLOW) "Run 'make tidy' to clean up the go.mod file."

.PHONY: build
## build: [build]* Builds and installs the Terraform provider locally.
build: tidy
	@ $(HEADER) "=====> Building and installing the provider..."
	$(GO) install -a -ldflags="-s -w" .
	@ ls -lahF $(GOBIN)/$(BINARY_NAME)
# @config-manager:end go_compile

#-------------------------------------------------------------------------------
# Clean

# @config-manager:start go_clean
.PHONY: clean-go
## clean-go: [clean] Clean Go's module cache.
clean-go:
	@ $(HEADER) "=====> Cleaning Go cache..."
	$(GO) clean -i -r -x -testcache -modcache -cache
# @config-manager:end go_clean

.PHONY: clean-tests
## clean-tests: [clean] Cleans all test artifacts.
clean-tests:
	@ $(HEADER) "=====> Cleaning artifacts from tests..."
	- find . -type d -name ".terraform" | xargs -I% rm -fv "%"
	- find . -type d -name "terratest-*" | xargs -I% rm -fv "%"
	- find . -type f -name "terraform.tfstate*" | xargs -I% rm -fv "%"
	- find ./examples -type d -name "\.*" | xargs -I% rm -fv "%"

# @config-manager:start go_clean_bench
.PHONY: clean-bench
## clean-bench: [clean] Cleans all benchmarking-related files.
clean-bench:
	@ $(HEADER) "=====> Cleaning artifacts from benchmarks..."
	- find . -type f -name "__*.out" | xargs -I% rm -fv "%"
	- find . -type f -name "*.test" | xargs -I% rm -fv "%"
# @config-manager:end go_clean_bench

# @config-manager:start tf_clean
.PHONY: clean-tf
## clean-tf: [clean] Clean Terra/Fu leftovers.
clean-tf:
	@ $(HEADER) "=====> Cleaning Terra/Fu artifacts..."
	find . -type d -name "terraform.d" | xargs -I% rm -Rfv "%"
	find . -type d -name ".terraform" | xargs -I% rm -Rfv "%"
	find . -type f -name ".terraform.lock.hcl" | xargs -I% rm -fv "%"
# @config-manager:end tf_clean

.PHONY: clean
## clean: [clean]* Runs ALL cleaning tasks (except the Go cache).
clean: clean-bench clean-tf clean-tests clean-ds

#-------------------------------------------------------------------------------
# Documentation

.PHONY: docs
## docs: [docs]* Runs primary documentation tasks.
docs: docs-provider docs-cli

.PHONY: docs-provider
## docs-provider: [docs] Generate Terraform Registry documentation.
docs-provider: clean-ds
	@ $(HEADER) "=====> Generating Terraform provider documentation..."
	$(GO) generate -v ./...

	@ $(HEADER) "=====> Remove tfplugindocs comments..."
	find docs/ -type f -name "*.md" | xargs -I% sed -i 's|<!-- arguments generated by tfplugindocs -->||g' "%"
	find docs/ -type f -name "*.md" | xargs -I% sed -i 's|<!-- variadic argument generated by tfplugindocs -->||g' "%"

	@ echo " "
	# If this fails, there's a bigger issue.
	$(GOTOOLS) lefthook run pre-commit --commands markdownlint

.PHONY: docs-cli
## docs-cli: [docs] Preview the Go library documentation on the CLI.
docs-cli:
	@ $(HEADER) "=====> Displaying Go CLI documentation..."
	$(GO) doc -C corefunc/ -all

# @config-manager:start go_docs_serve
.PHONY: docs-serve
## docs-serve: [docs] Preview the Go library documentation as displayed on pkg.go.dev.
docs-serve:
	@ $(HEADER) "=====> Displaying Go HTTP documentation..."
	$(GO) doc -http
# @config-manager:end go_docs_serve

# @config-manager:start go_binsize
.PHONY: binsize
## binsize: [docs] Analyze the size of the binary by Go package.
binsize:
	@ $(HEADER) "=====> Displaying Go HTTP documentation..."
	$(GO) tool nm -size "$(GOBIN)/$(BINARY_NAME)" | go-binsize-treemap > binsize.svg
	rsvg-convert --width=2000 --format=png --output="binsize.png" "binsize.svg"
# @config-manager:end go_binsize

#-------------------------------------------------------------------------------
# Linting

# @config-manager:start precommit
.PHONY: pre-commit
## pre-commit: [lint]* Runs `pre-commit` against all files.
pre-commit:
	@ $(HEADER) "=====> Running pre-commit..."
	pre-commit run --all-files
# @config-manager:end precommit

# @config-manager:start go_license
.PHONY: license
## license: [lint]* Checks the licenses of all files and dependencies.
license:
	@ $(HEADER) "=====> Checking license usage..."
	@ - trivy fs --config trivy-license.yaml --format json . 2>/dev/null > .licenses.cache.json
	@ cat .licenses.cache.json | jq -Mr '[.Results[] | select(.Packages) | .Packages[] | select(.Licenses) | .Licenses[]] | to_entries | group_by(.value)[] | {Name: .[0].value, Count: length} | "\(.Name): \(.Count)"'

	@ $(HEADER) "=====> Checking license headers..."
	@ echo "Missing/outdated:"
	@ - licensei header
# @config-manager:end go_license

.PHONY: lint
## lint: [lint]* Runs ALL linting/validation tasks.
lint: license pre-commit

#-------------------------------------------------------------------------------
# Testing
# https://github.com/golang/go/wiki/TableDrivenTests
# https://go.dev/doc/tutorial/fuzz
# https://pkg.go.dev/testing
# https://pkg.go.dev/golang.org/x/perf/cmd/benchstat

.PHONY: test
## test: [test]* Runs ALL tests.
test: unit examples acc mutate terratest bats

# @config-manager:start go_tests_stub
#===============================================================================
# After bootstrapping, remove the `@config-manager` markers from around these
# tasks, then fill them out as appropriate for your project.
#===============================================================================
PHONY: list-tests
## list-tests: [test] Lists all of the tests that are available to run.
list-tests:
	@ $(HEADER) "=====> Unit tests..."

	@ $(HEADER) "=====> Example tests..."

	@ $(HEADER) "=====> Fuzzing tests..."

	@ $(HEADER) "=====> Quick benchmarks..."

	@ $(HEADER) "=====> Full benchmarks..."

.PHONY: unit
## unit: [test] Runs unit tests. Set NAME= (without 'Test') to run a specific test by name.
unit:
	@ $(HEADER) "=====> Running unit tests..."
	$(GOTOOLS) gotestsum --format testname -- -run=Test$(NAME) -count=1 -parallel=$(shell nproc) -timeout 30s -coverpkg=./... -coverprofile=__coverage.out -v ./...
	@ go-cover-treemap -coverprofile __coverage.out > unit-coverage.svg
	@ rsvg-convert --width=2000 --format=png --output="unit-coverage.png" "unit-coverage.svg"

.PHONY: mutate
## mutate: [test] Runs mutation tests.
mutate:
	@ $(HEADER) "=====> Running mutation tests..."
	$(GO) test -tags=mutation -count=1 -parallel=$(shell nproc) -timeout 30s

.PHONY: examples
## examples: [test] Runs tests for examples. Set NAME= (without 'Example') to run a specific test by name.
examples:
	@ $(HEADER) "=====> Running tests for examples..."
	$(GOTOOLS) gotestsum --format testname -- -run=Example$(NAME) -count=1 -parallel=$(shell nproc) -timeout 30s -coverpkg=./... -coverprofile=__coverage.out -v ./...

.PHONY: fuzz
## fuzz: [test]* Runs the fuzzer for 1 minute per test.
fuzz:
	@ $(HEADER) "=====> Running the fuzzer (https://go.dev/doc/tutorial/fuzz)..."
	$(GO) test -list=Fuzz ./... | grep ^Fuzz | xargs -I% $(GO) test -run='^$$' -fuzz=% -fuzztime 1m -v ./

.PHONY: quickbench
## quickbench: [test]* Runs the benchmarks with minimal data for a quick check.
quickbench:
	@ $(HEADER) "=====> Running "quick" benchmark..."
	$(GO) test -bench=. -run=^Benchmark$(NAME) -timeout 60m ./

.PHONY: bench
## bench: [test]* Runs the benchmarks with enough data for analysis with benchstat.
bench:
	@ $(HEADER) "=====> Running "full" benchmark..."
	$(GO) test -bench=. -run=^Benchmark$(NAME) -count=6 -timeout 60m -benchmem -cpuprofile=__cpu.out -memprofile=__mem.out -trace=__trace.out ./ | tee __bench-$(shell date --utc "+%Y%m%dT%H%M%SZ").out

.PHONY: pgo
## pgo: [test] Runs the benchmarks with enough data for use with Profile-Guided Optimization.
pgo:
	@ $(HEADER) "=====> Running benchmark for PGO data..."
	$(GO) test -run=^TestAcc -count=6 -timeout 60m -cpuprofile=default.pgo -parallel=$(shell nproc) ./...
# @config-manager:end go_tests_stub

# @config-manager:start go_pprof
.PHONY: view-cov-cli
## view-cov-cli: [test] After running test or unittest, this will view the coverage report on the CLI.
view-cov-cli:
	@ $(HEADER) "=====> Viewing code coverage on the CLI..."
	$(GOTOOLS) gocovsh --profile=__coverage.out

.PHONY: view-cov-html
## view-cov-html: [test] After running test or unittest, this will launch a browser to view the coverage report.
view-cov-html:
	@ $(HEADER) "=====> Viewing code coverage as HTML..."
	$(GO) tool cover -html=__coverage.out

.PHONY: view-cpupprof
## view-cpupprof: [test] After running bench, this will launch a browser to view the CPU profiler results.
view-cpupprof:
	@ $(HEADER) "=====> Viewing CPU profiling in the browser..."
	$(GO) tool pprof -http :8080 __cpu.out

.PHONY: view-mempprof
## view-mempprof: [test] After running bench, this will launch a browser to view the memory profiler results.
view-mempprof:
	@ $(HEADER) "=====> Viewing memory profiling in the browser..."
	$(GO) tool pprof -http :8080 __mem.out

.PHONY: view-trace
## view-trace: [test] After running bench, this will launch a browser to view the trace results.
view-trace:
	@ $(HEADER) "=====> Viewing trace profiling in the browser..."
	$(GO) tool trace __trace.out
# @config-manager:end go_pprof

#-------------------------------------------------------------------------------
# Git Tasks

# @config-manager:start git_tag
.PHONY: tag
## tag: [release]* Signs and tags the release.
tag:
	@ $(HEADER) "=====> Signing and tagging the release..."
	@ if [ $$(git status -s -uall | wc -l) != 1 ]; then $(ERROR) "Git workspace must be clean."; exit 1; fi;

	@ $(WHITE) "This release will be tagged as: $(NEXT_VERSION)"
	@ echo "---------------------------------------------------------------------"
	@ read -p "Press any key to continue, or press Control+C to cancel. " x;

	@ echo " "
	@ chag update $(NEXT_VERSION)
	@ echo " "

	@ $(HEADER) "These are the contents of the CHANGELOG for this release. Are these correct?"
	@ echo "---------------------------------------------------------------------"
	@ chag contents
	@ echo "---------------------------------------------------------------------"
	@ echo "Are these release notes correct? If not, cancel and update CHANGELOG.md."
	@ read -p "Press any key to continue, or press Control+C to cancel. " x;

	@ echo " "

	git add .
	git commit -a -m "relprep: Preparing the $(NEXT_VERSION) release." --no-verify
	chag tag --sign
# @config-manager:end git_tag
