#!/bin/bash

# This script configures traffic director to route traffic to different
# backends. For detailed routing config, see url_map_template.yaml.

set -x

PROJECT_ID=$(gcloud config list --format 'value(core.project)')
BS_PREFIX=projects/${PROJECT_ID}/global/backendServices/${TEST_PREFIX}-grpcwallet
gcloud compute url-maps import ${TEST_PREFIX}-grpcwallet-url-map --source=<(sed -e "s/\$BS_PREFIX/${BS_PREFIX}/;s/\${TEST_PREFIX}/${TEST_PREFIX}/;s/\${TEST_PORT}/${TEST_PORT}/" url_map_template.yaml)

gcloud compute target-grpc-proxies create ${TEST_PREFIX}-grpcwallet-proxy \
    --url-map ${TEST_PREFIX}-grpcwallet-url-map \
    --validate-for-proxyless

gcloud compute forwarding-rules create ${TEST_PREFIX}-grpcwallet-forwarding-rule \
    --global \
    --load-balancing-scheme=INTERNAL_SELF_MANAGED \
    --address=0.0.0.0 --address-region=us-central1 \
    --target-grpc-proxy=${TEST_PREFIX}-grpcwallet-proxy \
    --ports 80 \
    --network default
