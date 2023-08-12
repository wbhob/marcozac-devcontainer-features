#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check GoReleaser latest version" goreleaser --version

reportResults
