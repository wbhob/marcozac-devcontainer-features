# Dev Container Features

## Contents

This repository contains following features:
- [buf](./src/buf/README.md): Install the [Buf CLI](https://buf.build/docs/installation)

## Usage

To use this features, add them in your `devcontainer.json` as in the example below.

```jsonc
{
    "image": "mcr.microsoft.com/devcontainers/base:debian",
    "features": {
        "ghcr.io/marcozac/devcontainer-features/buf:1": {}
    }
}
```

## Repo Structure

As in the [`devcontainers/features`](https://github.com/devcontainers/features) repository, each feature is located within a `src` sub-folder.

```
├── src
│   ├── buf
│   │   ├── devcontainer-feature.json
│   │   └── install.sh
```
