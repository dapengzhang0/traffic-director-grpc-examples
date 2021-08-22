#!/bin/bash
set -x

# This scripts deletes all resources created by this example.

# TD resources
gcloud compute forwarding-rules delete ${TEST_PREFIX}-grpcwallet-forwarding-rule --global -q
gcloud compute target-grpc-proxies delete ${TEST_PREFIX}-grpcwallet-proxy -q
gcloud compute url-maps delete ${TEST_PREFIX}-grpcwallet-url-map -q

# per service
services=(
    "${TEST_PREFIX}-grpcwallet-account"
    "${TEST_PREFIX}-grpcwallet-stats"
    "${TEST_PREFIX}-grpcwallet-stats-premium"
    "${TEST_PREFIX}-grpcwallet-wallet-v1"
    "${TEST_PREFIX}-grpcwallet-wallet-v2"
)
gcloud compute backend-services delete ${TEST_PREFIX}-grpcwallet-wallet-v1-affinity-service --global -q
for s in "${services[@]}"; do
    gcloud compute backend-services delete "$s"-service --global -q
    gcloud compute instance-groups managed delete "$s"-mig-us-central1 --zone us-central1-a -q
    gcloud compute instance-templates delete "$s"-template -q
done

# health check & firewall
gcloud compute firewall-rules delete ${TEST_PREFIX}-grpcwallet-allow-health-checks -q
gcloud compute health-checks delete ${TEST_PREFIX}-grpcwallet-health-check -q
