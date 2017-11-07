#!/bin/bash

set -eo pipefail

trap exit INT TERM

####
sample="deploys.test.myservice:1|c"

[ -z $EACH ] && EACH=2 || :

printf "[info] sending \"%s\" every %ssec\n" $sample $EACH

####

while true; do
    echo "deploys.test.myservice:1|c" > /dev/udp/127.0.0.1/8125
    sleep $EACH
done

####

printf "[info] terminating\n"