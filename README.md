# goaleasoluciones

[![CI](https://github.com/aleasoluciones/goaleasoluciones/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/aleasoluciones/goaleasoluciones/actions/workflows/ci.yml)
[![GoDoc](https://godoc.org/github.com/aleasoluciones/http2amqp?status.png)](http://godoc.org/github.com/aleasoluciones/http2amqp)
[![License](https://img.shields.io/github/license/aleasoluciones/http2amqp)](https://github.com/aleasoluciones/http2amqp/blob/master/LICENSE)

Common libraries for writing go services/applications.

- cirtuitbreaker
- clock
- crontask
- log
- retrier
- safemap
- scheduledtask
- timetoken
- yamlreader

## Build

You need a Go runtime installed in your system which supports [modules](https://tip.golang.org/doc/go1.16#modules). A nice way to have multiple Go versions and switch easily between them is the [g](https://github.com/stefanmaric/g) application.

```sh
make build
```

## Testing

```
make test
```
