#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check gofumpt latest version" gofumpt --version

reportResults
