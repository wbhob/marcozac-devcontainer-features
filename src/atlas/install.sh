#!/bin/bash -i

set -e

ATLAS_VERSION="${VERSION:-"latest"}"

COMMUNITY_EDITION="${COMMUNITYEDITION:-"true"}"
[ "$COMMUNITY_EDITION" == "true" ] && ATLAS_FLAVOR="community"

source ./atlas.sh -y
