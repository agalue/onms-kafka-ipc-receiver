#!/bin/bash -e
# @author Alejandro Galue <agalue@opennms.org>

GROUP_ID=${GROUP_ID-sink-go-client}

function join { local IFS="$1"; shift; echo "$*"; }

IFS=$'\n'
CONSUMER=(-parameter "acks=1")
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
if [ ! -z "${BOOTSTRAP_SERVERS}" ]; then
  OPTIONS+=(-bootstrap "${BOOTSTRAP_SERVERS}")
fi
if [ ! -z "${GROUP_ID}" ]; then
  OPTIONS+=(-group-id "${GROUP_ID}")
fi
if [ ! -z "${TOPIC}" ]; then
  OPTIONS+=(-topic "${TOPIC}")
fi
if [ "${IS_FLOW}" == "true" ]; then
  OPTIONS+=(-is-flow)
fi

exec /sink-receiver ${OPTIONS[@]} ${CONSUMER[@]}
