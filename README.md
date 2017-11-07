# DataDog Server mock

2MB golang server listening on port 8125 writing metrics to stdout.

### Bootstrap

```
make install
make test
```

### Running

```
make run
```

or

```
docker run -it --log-driver none --rm -p 0.0.0.0:8125:8125/UDP datadog/mock
```

or

```
./target/datadog_mock
```


### Verify

When datadog/mock is running 

```
echo "deploys.test.myservice:1|c" > /dev/udp/127.0.0.1/8125
```

or

```
EACH=.001 ./dev/event-producer.sh
```

