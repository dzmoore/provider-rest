#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
echo "Creating the provider config with cluster admin permissions in cluster..."
SA=$(${KUBECTL} -n crossplane-system get sa -o name | grep provider-rest | sed -e 's|serviceaccount\/|crossplane-system:|g')
${KUBECTL} create clusterrolebinding provider-rest-admin-binding --clusterrole cluster-admin --serviceaccount="${SA}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

cat <<EOF | ${KUBECTL} apply -f -
apiVersion: rest.crossplane.io/v1alpha1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: None
EOF
