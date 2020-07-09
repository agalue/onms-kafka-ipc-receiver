#!/bin/bash
# @author Alejandro Galue <agalue@opennms.org>

type protoc >/dev/null 2>&1 || { echo >&2 "protoc required but it's not installed; aborting."; exit 1; }

mkdir -p sink
protoc -I . sink.proto --go_out=./sink
