#!/bin/sh
if [ "$IS_CONTAINER" != "" ]; then
  go vet "${@}"
else
  podman run --rm \
    --env IS_CONTAINER=TRUE \
    --volume "${PWD}:/go/src/github.com/openshift-metalkube/kni-installer:z" \
    --workdir /go/src/github.com/openshift-metalkube/kni-installer \
    docker.io/openshift/origin-release:golang-1.10 \
    ./hack/go-vet.sh "${@}"
fi;
