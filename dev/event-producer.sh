#!/bin/bash

set -eo pipefail

trap exit INT TERM

####

[ -z $EACH ] && EACH=2 || :
[ -z $MESSAGE ] && MESSAGE="deploys.test.myservice:1|c" || :

printf "[info] sending \"%s\" every %ssec\n" $MESSAGE $EACH

####

while true; do
    echo $MESSAGE > /dev/udp/127.0.0.1/8125
    sleep $EACH
done

####

printf "[info] terminating\n"