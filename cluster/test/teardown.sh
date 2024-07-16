#!/usr/bin/env bash
set -aeuo pipefail

echo "Running teardown.sh"

${KUBECTL} delete providerconfigs.rest.crossplane.io --all --wait

echo "teardown.sh complete"