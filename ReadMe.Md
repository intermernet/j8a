[![Circleci Builds](https://circleci.com/gh/simonmittag/j8a.svg?style=shield)](https://circleci.com/gh/simonmittag/j8a)
[![Github Workflows](https://github.com/simonmittag/j8a/workflows/Go/badge.svg)](https://github.com/simonmittag/j8a/actions)
[![Dependabot](https://badgen.net/badge/Dependabot/enabled/green?icon=dependabot)](https://github.com/simonmittag/j8a/pulls?q=is%3Aopen+is%3Apr)
[![Github Issues](https://img.shields.io/github/issues/simonmittag/j8a)](https://github.com/simonmittag/j8a/issues)
[![Github Activity](https://img.shields.io/github/commit-activity/m/simonmittag/j8a)](https://img.shields.io/github/commit-activity/m/simonmittag/j8a)  
[![Go Report](https://goreportcard.com/badge/github.com/simonmittag/j8a)](https://goreportcard.com/report/github.com/simonmittag/j8a)
[![Codeclimate Maintainability](https://api.codeclimate.com/v1/badges/326d3c304e32645811c8/maintainability)](https://codeclimate.com/github/simonmittag/jabba/maintainability)
[![Codeclimate Test Coverage](https://api.codeclimate.com/v1/badges/326d3c304e32645811c8/test_coverage)](https://codeclimate.com/github/simonmittag/jabba/code)
[![SSLLabs Rating](https://img.shields.io/badge/SSLLabs-A%2B-brightgreen)](https://www.ssllabs.com/ssltest/analyze.html?d=j8a.io)
[![Go Version](https://img.shields.io/github/go-mod/go-version/simonmittag/j8a)](https://img.shields.io/github/go-mod/go-version/simonmittag/j8a)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Docker Pulls](https://img.shields.io/docker/pulls/simonmittag/j8a)](https://hub.docker.com/repository/docker/simonmittag/j8a)
[![Version](https://img.shields.io/docker/v/simonmittag/j8a/0.9.5)](https://hub.docker.com/repository/docker/simonmittag/j8a)

## What's new?

### v0.9.5 (16/06/2022)
* now logging ulimit (max open files) so admin has visibility next to maxopenconns

### v0.9.4 (14/06/2022)
* changes to standard log levels. All server internal messages now at info, downstream at debug, upstream data at trace level.

### v0.9.3 (12/06/2022)
* Added homebrew distribution under brew tap simonmittag/cli
* Added thread logging and upstream connection logging
* bumped gojq 0.12.8, lestrrat-go/jwx 1.2.25, klauspost/compress 1.15.6

### v0.9.2 (12/04/2022)
* Shorter default http upstream timeout, for details see https://j8a.io/docs
* upgraded jwx dependency to v1.2.22

## What is j8a?

[j8a [ dʒʌbbʌ ]](https://j8a.io) is a modern **TLS 1.3 reverse proxy server** designed as a lightweight API gateway for
micro-service style JSON APIs. It is built for DevOps engineers and allows fine-grained traceability of HTTP and
Websocket traffic without sacrificing performance.

![](images/route.gif)

## Features

* Fast! 5k POST req/s traffic.
* Secure. TLS1.2, TLS/1.3 termination w/ [A+ SSLLabs rating](https://www.ssllabs.com/ssltest/analyze.html?d=j8a.io).
* Zero downtime. Auto-renew certs using ACME [RFC 8555](https://tools.ietf.org/html/rfc8555).
* Observable. API request tracing w/ correlation ID for upstream micro-services.
* APM. CPU, memory logging built-in. Daily TLS health and validity check for your full certificate chain.
* JWT token validation with full JWK(S) support for [RFC 7519](https://tools.ietf.org/html/rfc7519).
* HTTP/1.1 and HTTP/2 support w/ upstream <> downstream protocol translation.
* Websocket Support for [RFC 6455](https://datatracker.ietf.org/doc/html/rfc6455)
* [Docker](https://hub.docker.com/repository/docker/simonmittag/j8a) native

## Up and running

### Docker
```
docker pull simonmittag/j8a &&
  docker run -e J8ACFG_YML -p443:443 simonmittag/j8a
```

### Homebrew
```bash
brew tap simonmittag/cli && 
  brew install j8a && 
  j8a
```

### Golang
```bash
go install github.com/simonmittag/j8a/cmd/j8a &&
  j8a
```

## Configuration

### Environment Variables

#### J8ACFG_YML

Supply the server configuration to the running process in .yml format

```bash
J8ACFG_YML="---
            connection:
              downstream:
                readTimeoutSeconds: 3
                roundTripTimeoutSeconds: 20
                idleTimeoutSeconds: 30
                http:
                  port: 80
                maxBodyBytes: 65535
              upstream:
                socketTimeoutSeconds: 3
                readTimeoutSeconds: 30
                idleTimeoutSeconds: 10
                maxAttempts: 4
                poolSize: 8
                tlsInsecureSkipVerify: true
            routes:
              - path: "/todos"
                resource: jsonplaceholder
            resources:
              jsonplaceholder:
                - url:
                    scheme: https
                    host: jsonplaceholder.typicode.com
                    port: 443"
```

#### LOGLEVEL

Loglevel for stdout, one of TRACE, DEBUG, INFO or WARN

```bash
LOGLEVEL="DEBUG"
```

#### TZ

TimeZone

```bash
TZ="Australia/Sydney"
```

## Contributions

The j8a team welcomes all [contributors](https://github.com/simonmittag/j8a/blob/master/CONTRIBUTING.md). Everyone
interacting with the project's codebase, issue trackers, chat rooms and mailing lists is expected to follow
the [code of conduct](https://github.com/simonmittag/j8a/blob/master/CODE_OF_CONDUCT.md)
