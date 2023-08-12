#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check ShellCheck fixed version" shellcheck --version | grep "0.8.0"

reportResults
