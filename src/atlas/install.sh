#!/bin/bash -i

set -e

ATLAS_VERSION="${VERSION:-"latest"}"
COMMUNITY_EDITION="${COMMUNITYEDITION:-"true"}"
GO_VERSION="${GOVERSION:-"1.20.7"}"
_lib_path="/usr/local/lib/atlas"

source ./library_scripts_extra.sh

mkdir -p "$_lib_path"

if [ "$COMMUNITY_EDITION" == "true" ]; then
    if [ "$ATLAS_VERSION" != "latest" ]; then
        echo "Community edition binaries are only available for the latest version"
        echo "Building from source..."
        build_from_source
    else
        source ./atlas.sh -y --output "$_lib_path/atlas" --community
    fi
    download_license
else
    source ./atlas.sh --output "$_lib_path/atlas" -y
    download_eula
fi

ln -sf "$_lib_path/atlas" "/usr/local/bin/atlas"
