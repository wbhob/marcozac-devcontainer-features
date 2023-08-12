#!/bin/bash -i

set -e

source dev-container-features-test-lib

check "Check protoc-gen-entgrpc fixed version" type protoc-gen-entgrpc

reportResults
