#!/usr/bin/env bash
set -eu -o pipefail

VERSION=${VERSION:-$(git describe --abbrev=0 --broken)}
COMMIT=${COMMIT:-$(git rev-parse --short HEAD)}
# BSD date command doesn't support RFC3339/ISO8601 or subsecond precision.
BUILDDATE=${BUILDDATE:-$(date -u -Ins 2> /dev/null || true)}
BUILDDATE=${BUILDDATE:-$(date -u +%FT%T000000000%z)}

export LDFLAGS="\
  -X \"github.com/Shopify/kubeaudit/cmd.Version=${VERSION}\" \
  -X \"github.com/Shopify/kubeaudit/cmd.Commit=${COMMIT}\" \
  -X \"github.com/Shopify/kubeaudit/cmd.BuildDate=${BUILDDATE}\" \
  ${LDFLAGS:-} \
"
