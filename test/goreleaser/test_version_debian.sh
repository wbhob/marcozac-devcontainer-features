#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check GoReleaser fixed version" goreleaser --version | grep "1.20.0"

reportResults
