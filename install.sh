#!/usr/bin/env bash
CURRENT_FILE=$(basename $0)

BIN_PATH=${1:-${HOME}/bin}

for file in $(find . -maxdepth 1 -type f -executable ! -path "./${CURRENT_FILE}" -printf '%P\n'); do
    ln -fs "$(pwd)/${file}" "${BIN_PATH}/${file}"
done
