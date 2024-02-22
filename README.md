# reqdump - dump the request out

A dumb little go web app to dump the request it receives both to a
logger and echo it back to the client.

## Docker

Build the image thusly:

```shell
$ docker build -t rweir/reqdump .
```

To run it, exposing it via random port you have to look up via `docker ps`:

```shell
$ docker run -P rweir/reqdump
```

To run it, using a fixed port:

```shell
$ docker run -p9999:8080 rweir/reqdump
```
