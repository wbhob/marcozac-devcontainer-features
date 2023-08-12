#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check Atlas version: versioned Community Edition" atlas version | grep "development"
check "Check Atlas license: Apache 2" cat "/usr/local/lib/atlas/LICENSE" | grep "Apache License"

reportResults
