#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check Atlas version: latest EULA release" atlas version | grep -v community
check "Check Atlas EULA" cat "/usr/local/lib/atlas/EULA.pdf" | grep "EULA"

reportResults
