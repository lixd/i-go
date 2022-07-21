#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
cd "${KUBE_ROOT}/hack" || exit 1

if ! command -v goimports &> /dev/null
then
    echo "goimports could not be found on your machine, please install it first"
    exit
fi

cd "${KUBE_ROOT}" || exit 1

IFS=$'\n' read -r -d '' -a files < <( find . -type f -name '*.go' ! -name 'mock_*.go' ! -name '*.pb.go' && printf '\0' )

"goimports" -w -local github.com/kubeclipper/kubeclipper "${files[@]}"
