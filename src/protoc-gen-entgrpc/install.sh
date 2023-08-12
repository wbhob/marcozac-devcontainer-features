#!/bin/bash -i

set -e

go install "entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@$VERSION"
