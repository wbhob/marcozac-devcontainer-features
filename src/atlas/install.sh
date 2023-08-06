#!/bin/bash -i

set -e

ATLAS_VERSION="${VERSION:-"latest"}"
COMMUNITY_EDITION="${COMMUNITYEDITION:-"true"}"
GO_VERSION="${GOVERSION:-"1.20.7"}"

install_go() {
    source ./library_scripts.sh

    # nanolayer is a cli utility which keeps container layers as small as possible
    # source code: https://github.com/devcontainers-contrib/nanolayer
    # `ensure_nanolayer` is a bash function that will find any existing nanolayer installations,
    # and if missing - will download a temporary copy that automatically get deleted at the end
    # of the script
    ensure_nanolayer nanolayer_location "v0.4.45"

    $nanolayer_location \
        install \
        devcontainer-feature \
        "ghcr.io/devcontainers/features/go:1" \
        --option version="$GO_VERSION" \
        --env "INSTALL_GO_TOOLS=false"
}

build_from_source() {
    local go_location
    go_location="$(which go 2>/dev/null)" || go_location="/usr/local/go/bin/go" && ($go_location version >/dev/null 2>&1 || install_go)
    $go_location install "ariga.io/atlas/cmd/atlas@$ATLAS_VERSION"
    mv "$(/usr/local/go/bin/go env GOPATH)/bin/atlas" /usr/local/bin
}

if [ "$COMMUNITY_EDITION" == "true" ]; then
    if [ "$ATLAS_VERSION" != "latest" ]; then
        echo "Community edition binaries are only available for the latest version"
        echo "Building from source..."
        build_from_source
    else
        source ./atlas.sh -y --community
    fi
else
    source ./atlas.sh -y
fi
