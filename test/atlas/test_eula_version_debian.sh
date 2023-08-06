#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check Atlas version: versioned EULA release - License" atlas version | grep -v community
check "Check Atlas version: versioned EULA release - Version" atlas version | grep "v0.12.0"

reportResults
