#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check protoc-gen-entgrpc latest version" type protoc-gen-entgrpc

reportResults
