#!/usr/bin/env bash

# Install pre commit & required dependencies.
set -eu
set -o pipefail

echo "--Installing git-hooks & dependencies---"
if ! hash pre-commit > /dev/null; then
    echo "[[ pre-commit not found installing via brew ]]\n"
    brew install pre-commit
fi

pre-commit install
chmod 755 .git/hooks/pre-commit
