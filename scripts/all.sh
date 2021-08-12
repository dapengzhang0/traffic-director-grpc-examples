#!/bin/bash

# This script creates all resources needed to demonstrate routing.
# - health check, for checking backend health
# - backend services for all services
# - url-map and other traffic director configuration
#
# It takes a parameter to specify the language for the servers. Run as
# `./all.sh <language>`.

set -x

TEST_PREFIX=${TEST_PREFIX-""}
TEST_PORT=${TEST_PORT-"80"}

# $1 = language, one of "java" or "go"
if ! [[ $1 =~ ^(go|java)$ ]]; then
    echo "language $1 is undefined, pick one from [go, java]"
    exit 123
fi

./create_health_check.sh
./create_service.sh $1 account 50053 account
./create_service.sh $1 stats   50052 stats         "--account_server=xds:///${TEST_PREFIX}account.grpcwallet.io:${TEST_PORT}"
./create_service.sh $1 stats   50052 stats-premium "--account_server=xds:///${TEST_PREFIX}account.grpcwallet.io:${TEST_PORT} --premium_only=true"
./create_service.sh $1 wallet  50051 wallet-v1     "--account_server=xds:///${TEST_PREFIX}account.grpcwallet.io:${TEST_PORT} --stats_server=xds:///${TEST_PREFIX}stats.grpcwallet.io:${TEST_PORT} --v1_behavior=true"
./create_service.sh $1 wallet  50051 wallet-v2     "--account_server=xds:///${TEST_PREFIX}account.grpcwallet.io:${TEST_PORT} --stats_server=xds:///${TEST_PREFIX}stats.grpcwallet.io:${TEST_PORT}"
./config_traffic_director.sh
