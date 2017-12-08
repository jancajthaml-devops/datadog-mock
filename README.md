# datadog-mock #

datadog-mock is a golang statsd mock server listening on port 8125 and relaying events to stdout.

[![Docker Version](https://images.microbadger.com/badges/version/jancajthaml/datadog_mock.svg)](https://microbadger.com/images/jancajthaml/datadog_mock)
[![Static Analysis](https://api.codacy.com/project/badge/Grade/c5c255a292f84cf88972f92f74f9174d)](https://www.codacy.com/app/jan-cajthaml/datadog-mock?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=jancajthaml/datadog-mock&amp;utm_campaign=Badge_Grade) [![Go Report Card](https://goreportcard.com/badge/jancajthaml-bash/datadog-mock)](https://goreportcard.com/report/jancajthaml-bash/datadog-mock)

## Getting started ##

Bootstrap environment with `make install test` then grab `./target/datadog_mock`
or docker image `datadog/mock`.

## Testing ##

Run in docker with `make run` or `./target/datadog_mock` locally.

When datadog/mock is running you can either test simple relay
`./dev/event-producer.sh` or siege with `make perf`.

## License ##

This service is distributed under the Apache License, Version 2.0 license found
in the [LICENSE](./LICENSE) file.


