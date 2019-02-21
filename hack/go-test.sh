#!/bin/sh
# Example:  ./hack/go-test.sh

if [ "$IS_CONTAINER" != "" ]; then
  go test ./cmd/... ./data/... ./pkg/... "${@}"
else
  podman run --rm \
    --env IS_CONTAINER=TRUE \
    --volume "${PWD}:/go/src/github.com/openshift-metalkube/kni-installer:z" \
    --workdir /go/src/github.com/openshift-metalkube/kni-installer \
    docker.io/openshift/origin-release:golang-1.10 \
    ./hack/go-test.sh "${@}"
fi
