#!/bin/bash

set -e

GAZELLE_MODE=$1

ROOTDIR=$(dirname "$0")/..

GOSDK=$(bazel info output_base 2>/dev/null)/external/go_sdk/bin

function generate {
    PACKAGE=$1

    bazel build //antlr:${PACKAGE}
    rm -rf ${ROOTDIR}/antlr/${PACKAGE}
    mkdir ${ROOTDIR}/antlr/${PACKAGE}
    cp -r ${ROOTDIR}/bazel-bin/antlr/${PACKAGE}.go/${PACKAGE}/* ${ROOTDIR}/antlr/${PACKAGE}
    chmod 0644 ${ROOTDIR}/antlr/${PACKAGE}/*
    ${GOSDK}/gofmt -w antlr/${PACKAGE}/*.go
    # Make the generated files deterministic.
    sed -i '/Code generated from/c\// File generated by ANTLR. DO NOT EDIT.' ${ROOTDIR}/antlr/${PACKAGE}/*
}

generate traffic_class
generate sequence

bazel run //:gazelle -- update -mode=${GAZELLE_MODE} -go_naming_convention go_default_library -index=false -external=external ${ROOTDIR}/antlr
