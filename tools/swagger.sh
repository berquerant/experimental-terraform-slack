#!/bin/bash

run() {
    go run github.com/go-swagger/go-swagger/cmd/swagger "$@"
}

generate() {
    spec="$1"
    target="$2"
    principal="$3"
    if [ -z "${target}" ] ; then
        return 1
    fi
    mkdir -p "${target}"
    run generate server \
        --spec "${spec}" \
        --target "${target}" \
        --principal "${principal}"
    run generate client \
        --spec "${spec}" \
        --target "${target}" \
        --principal "${principal}"
}

validate() {
    run validate "$1"
}

usage() {
    echo "${0} [generate|validate|run|do]"
    echo "${0} generate SPEC TARGET PRINCIPAL"
    echo "${0} validate SPEC"
    echo "${0} run SPEC TARGET PRINCIPAL"
    echo "${0} do opt..."
}

main() {
    set -ex

    cmd="$1"
    spec="$2"
    target="$3"
    principal="$4"

    case "${cmd}" in
        "generate")
            generate "${spec}" "${target}" "${principal}"
            ;;
        "validate")
            validate "${spec}"
            ;;
        "run")
            validate "${spec}" && generate "${spec}" "${target}" "${principal}"
            ;;
        "do")
            shift
            run "$@"
            ;;
        *)
            return 1
            ;;
    esac
}

if ! main "$@" ; then
    usage
    exit 1
fi
