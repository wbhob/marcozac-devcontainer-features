#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check Atlas version" atlas version | grep "development"

reportResults
