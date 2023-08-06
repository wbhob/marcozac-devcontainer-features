#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check Atlas EULA release" atlas version | grep -v community
check "Check Atlas version" atlas version | grep "v0.12.0"

reportResults
