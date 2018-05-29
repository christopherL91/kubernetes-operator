#!/bin/bash

function main {
    local DOCKER_REPO_ROOT="/go/src/github.com/christopherL91/k8s-operator"
    local IMAGE=${IMAGE:-"gcr.io/coreos-k8s-scale-testing/codegen:1.10"}

    docker run --rm \
    -v "$PWD":"$DOCKER_REPO_ROOT" \
    -w "$DOCKER_REPO_ROOT" \
    "$IMAGE" \
        "/go/src/k8s.io/code-generator/generate-groups.sh"  \
        "all" \
        "github.com/christopherL91/k8s-operator/pkg/client" \
        "github.com/christopherL91/k8s-operator/pkg/apis" \
        "foo:v1alpha1" \
        --go-header-file "./hack/boilerplate.go.txt" \
    $@
}

main $@
exit $?