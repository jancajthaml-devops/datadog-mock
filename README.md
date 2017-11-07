# DataDog Server mock

golang server listening on port 8125 writing metrics to stdout. (datadog/mock 3.02MB)

### Bootstrap

Require docker to assembly

`make install test`

### Running

automagically

`make run`

run in docker

`docker run -it --log-driver none --rm -p 0.0.0.0:8125:8125/UDP datadog/mock`

run on host

`./target/datadog_mock`


### Verify

When datadog/mock is running run manually

`echo "deploys.test.myservice:1|c" > /dev/udp/127.0.0.1/8125`

or send MESSAGE repeat EACH second

`MESSAGE="deploys.test.myservice:1|c" EACH=.001 ./dev/event-producer.sh`

