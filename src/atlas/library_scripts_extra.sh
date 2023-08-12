#!/bin/bash -i

_license_url="https://raw.githubusercontent.com/ariga/atlas/master/LICENSE"
_eula_url="https://ariga.io/legal/atlas/eula"

source ./library_scripts.sh

install_go() {
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
    mv "$($go_location env GOPATH)/bin/atlas" "$_lib_path"
}

download_license() {
    clean_download "$_license_url" "$_lib_path/LICENSE"
}

download_eula() {
    clean_download "$_eula_url" "$_lib_path/EULA.pdf"
}
