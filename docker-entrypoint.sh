#!/bin/bash -e
# @author Alejandro Galue <agalue@opennms.org>

GROUP_ID=${GROUP_ID-sink-go-client}

function join { local IFS="$1"; shift; echo "$*"; }

IFS=$'\n'
CONSUMER=()
for VAR in $(env)
do
  env_var=$(echo "$VAR" | cut -d= -f1)
  if [[ $env_var =~ ^KAFKA_ ]]; then
    key=$(echo "$env_var" | cut -d_ -f2- | tr '[:upper:]' '[:lower:]' | tr _ .)
    val=${!env_var}
    if [[ $key == "manager."* ]]; then
      echo "[Skipping] '$key'"
    else
      echo "[Configuring] '$key'='$val'"
      CONSUMER+=(-parameter "$key=$val")
    fi
  fi
done

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
if [ "${IS_TELEMETRY}" == "true" ]; then
  OPTIONS+=(-is-telemetry)
fi
if [ "${IS_NETFLOW}" == "true" ]; then
  OPTIONS+=(-is-netflow)
fi

echo "Starting onms-kafka-ipc-receiver with: ${OPTIONS[@]} ${CONSUMER[@]}"
exec /onms-kafka-ipc-receiver ${OPTIONS[@]} ${CONSUMER[@]}
