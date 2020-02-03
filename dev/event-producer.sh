#!/bin/bash

set -eo pipefail
trap exit INT TERM

####

echo "[info] starting"

EACH=${EACH:=2}
MESSAGE=${MESSAGE:="deploys.test.myservice:1|c"}

echo "[info] sending \"${MESSAGE}\" every ${EACH}sec"

####

while true ; do
  echo "${MESSAGE}" > /dev/udp/127.0.0.1/8125
  sleep "${EACH}"
done

####

echo "[info] terminating"
