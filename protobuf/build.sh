#!/bin/bash
# @author Alejandro Galue <agalue@opennms.org>

type protoc >/dev/null 2>&1 || { echo >&2 "protoc required but it's not installed; aborting."; exit 1; }

for module in rpc sink telemetry netflow flowdocument; do
  mkdir -p $module
  protoc --proto_path=./ --go_out=./ $module.proto
done
