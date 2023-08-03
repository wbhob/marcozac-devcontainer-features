#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "buf --version" buf --version

reportResults
