# Current status: draft :seedling:

# redis-metrics

redis-metrics inspects [Redis](http://redis.io/) and send metrics to [StatsD](https://github.com/etsy/statsd).

## How it works

redis-metrics executes [INFO](http://redis.io/commands/INFO) command at a given interval, filters statistics and sends metrics to StatsD.

## Usage

### Download

Download redis-metrics from [GitHub Releases](https://github.com/daime/redis-metrics/releases).

### Configuration

Create a `redis-metrics.json` anywhere.

Requirements and options are listed bellow:

```javascript
{
    // interval is the time in seconds of a inspection interval (default: 10)
    "interval": 10,
    // redis is a list of redis addresses to inspect (default: [{"host": "localhost", "port": 6379}])
    "redis": [{
        // host is required
        "host": "redis_1",
        // port is optional (default: 6379)
        "port": 6379
    }, {
        "host": "redis_2"
    }],
    // statistics is a list of INFO command fields that will be send to StatsD
    // only values that are valid floats are eligible to reach StatsD
    // check http://redis.io/commands/INFO#notes for the meaning of each field
    // special sections like keyspace are handled in a special way
    // example:
    // for each database, the following line is added:
    //  dbXXX: keys=XXX,expires=XXX
    // let's say only db0 is available and you want to track the number of keys from it
    // make sure statistics sections contains an entry "d0.keys"
    "statistics": [
        "used_memory",
        "mem_fragmentation_ratio",
        "db0.keys"
    ],
    // statsd is the host and port of StatsD that will receive the metrics
    "statsd": {
        "host": "statsd",
        // port is optional (default: 8125)
        "port": 8125
    },
    // metric is a pattern to compose each metric name that will reach StatsD
    // (default "redis-metrics.{{ Redis.Host }}:{{ Redis.Port }}.{{ Statistic }}")
    "metric": "redis-metrics.{{ .Redis.Host }}:{{ .Redis.Port }}.{{ .Statistic }}"
}
```

### Run

 ```console
 $ redis-metrics --config=/path/to/redis-metrics.json
 ```
