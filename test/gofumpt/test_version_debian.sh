#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check gofumpt fixed version" gofumpt --version | grep "0.5.0"

reportResults
