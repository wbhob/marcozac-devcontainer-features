#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check Atlas Community Edition" atlas version | grep community

reportResults
