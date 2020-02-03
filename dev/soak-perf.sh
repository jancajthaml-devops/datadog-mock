#!/bin/bash

base_dir=$(cd $(dirname $0) && pwd)

set -eo pipefail
exec 3>&1 4>&2

post_stop() {
  exit_code=$?
  printf "\n"
  if [ -n "$error" ] ; then
    printf "Error: %s\n" $error
  fi
  exec 3>&- 4>&-
  exit $exit_code
}

trap post_stop INT TERM
trap "kill 0" EXIT

MESSAGE="deploys.test.myservice:1|c" EACH=.0001 "${base_dir}"/event-producer.sh &> /dev/null &
MESSAGE="deploys.test.myservice:2|c" EACH=.0001 "${base_dir}"/event-producer.sh &> /dev/null &
MESSAGE="deploys.test.myservice:3|c" EACH=.0001 "${base_dir}"/event-producer.sh &> /dev/null &
MESSAGE="deploys.test.myservice:4|c" EACH=.0001 "${base_dir}"/event-producer.sh &> /dev/null &
MESSAGE="deploys.test.myservice:5|c" EACH=.0001 "${base_dir}"/event-producer.sh &> /dev/null &
docker rm -f $(docker ps -aq --filter="name=datadog" --filter="status=running") &> /dev/null || :

printf "processing [0 / sec]"

error=$( { (docker-compose run --rm --no-deps --service-ports artefact 2>&4 || :) | perl -e 'open(TTY, ">:unix", "/dev/tty");while (<>) {$l++;if (time > $e) {$e=time;print TTY "\rprocessing [$l / sec]";$l=0}}' 1>&3; } 2>&1 )
