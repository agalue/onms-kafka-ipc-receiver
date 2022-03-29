#!/bin/bash -e
# @author Alejandro Galue <agalue@opennms.org>

GROUP_ID=${GROUP_ID-onms-kafka-ipc-receiver}

OPTIONS=()
if [ ! -z "${BOOTSTRAP_SERVER}" ]; then
  OPTIONS+=(-bootstrap "${BOOTSTRAP_SERVER}")
fi
if [ ! -z "${GROUP_ID}" ]; then
  OPTIONS+=(-group-id "${GROUP_ID}")
fi
if [ ! -z "${TOPIC}" ]; then
  OPTIONS+=(-topic "${TOPIC}")
fi
if [ ! -z "${IPC}" ]; then
  OPTIONS+=(-ipc "${IPC}")
fi
if [ ! -z "${PARSER}" ]; then
  OPTIONS+=(-parser "${PARSER}")
fi

echo "Starting onms-kafka-ipc-receiver with: ${OPTIONS[@]}"
exec /onms-kafka-ipc-receiver ${OPTIONS[@]}
