#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check ShellCheck latest version" shellcheck --version

reportResults
