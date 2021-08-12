#!/bin/bash
set -x

# This scripts deletes all resources created by this example.

TEST_PREFIX=${TEST_PREFIX-""}

# TD resources
gcloud compute forwarding-rules delete ${TEST_PREFIX}grpcwallet-forwarding-rule --global -q
gcloud compute target-grpc-proxies delete ${TEST_PREFIX}grpcwallet-proxy -q
gcloud compute url-maps delete ${TEST_PREFIX}grpcwallet-url-map -q

# per service
services=(
    "grpcwallet-account"
    "grpcwallet-stats"
    "grpcwallet-stats-premium"
    "grpcwallet-wallet-v1"
    "grpcwallet-wallet-v2"
)
for s in "${services[@]}"; do
    gcloud compute backend-services delete "${TEST_PREFIX}${s}"-service --global -q
    gcloud compute instance-groups managed delete "${TEST_PREFIX}${s}"-mig-us-central1 --zone us-central1-a -q
    gcloud compute instance-templates delete "${TEST_PREFIX}${s}"-template -q
done

# health check & firewall
gcloud compute firewall-rules delete ${TEST_PREFIX}grpcwallet-allow-health-checks -q
gcloud compute health-checks delete ${TEST_PREFIX}grpcwallet-health-check -q
