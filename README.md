# RedisMetrics

## Build
**REMEMBER TO USE GO DIRECTORY STRUCTURE**
The project is based on go's default directory structure and uses `Makefile` to
handle the build tasks.

### Building
This will compile and create the executable in the project's root directory
```bash
make build
```

### Installing
This will compile and install the project in your `$GOPATH`
```bash
make install
```

### Testing

```bash
make test
```

### Running

Options
```bash
  [-c, --config]    Specify the configuration file directory (default: config.json)
```

Running:

```
redis-metrics
```
