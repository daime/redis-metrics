# redis-metrics

## Development

### Requirements

redis-metrics depends on [Go](https://golang.org/doc/install), [Docker](https://docs.docker.com/installation/) and [Docker Compose](https://docs.docker.com/compose/install/).

### Running

```bash
make docker
```

See [docker-compose.yml](https://github.com/daime/redis-metrics/blob/master/docker-compose.yml) to get an idea about what is running.

### Testing

```console
$ make test
```

### Coverage

```console
$ make coverage
```
