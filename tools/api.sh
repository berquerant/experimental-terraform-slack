#!/bin/bash

case "$1" in
    "s" | "srv" | "server")
        shift
        ./tmp/api-server "$@"
        ;;
    "c" | "client")
        shift
        ./tmp/api-client "$@"
        ;;
    *)
        ./tmp/api-server "$@"
esac
