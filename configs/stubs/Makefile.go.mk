#-------------------------------------------------------------------------------
# @config-manager:start includes
# @config-manager:end includes

# @config-manager:start binary_name
# @config-manager:end binary_name

#-------------------------------------------------------------------------------
# @config-manager:start go_globals
# @config-manager:end go_globals

#-------------------------------------------------------------------------------
# @config-manager:start standard_installation
# @config-manager:end standard_installation

#-------------------------------------------------------------------------------
# @config-manager:start go_compile
# @config-manager:end go_compile

#-------------------------------------------------------------------------------
# Clean

# @config-manager:start go_clean
# @config-manager:end go_clean

# @config-manager:start go_clean_bench
# @config-manager:end go_clean_bench

# @config-manager:start tf_clean
# @config-manager:end tf_clean

.PHONY: clean
## clean: [clean]* Runs ALL cleaning tasks (except the Go cache).
clean:

#-------------------------------------------------------------------------------
# Documentation

.PHONY: docs
## docs: [docs]* Runs primary documentation tasks.
docs:

# @config-manager:start go_docs_serve
# @config-manager:end go_docs_serve

# @config-manager:start go_binsize
# @config-manager:end go_binsize

#-------------------------------------------------------------------------------
# Linting

# @config-manager:start precommit
# @config-manager:end precommit

# @config-manager:start go_license
# @config-manager:end go_license

.PHONY: lint
## lint: [lint]* Runs ALL linting/validation tasks.
lint:

#-------------------------------------------------------------------------------
# Testing
# https://github.com/golang/go/wiki/TableDrivenTests
# https://go.dev/doc/tutorial/fuzz
# https://pkg.go.dev/testing
# https://pkg.go.dev/golang.org/x/perf/cmd/benchstat

.PHONY: test
## test: [test]* Runs ALL tests.
test:

# @config-manager:start go_tests_stub
# @config-manager:end go_tests_stub

# @config-manager:start go_pprof
# @config-manager:end go_pprof

#-------------------------------------------------------------------------------
# Git Tasks

# @config-manager:start git_tag
# @config-manager:end git_tag
