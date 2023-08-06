#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check Atlas EULA release" atlas version | grep -v community

reportResults
