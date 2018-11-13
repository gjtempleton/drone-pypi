# drone-pypi

[![Go Report Card](https://goreportcard.com/badge/github.com/xoxys/drone-pypi)](https://goreportcard.com/report/github.com/xoxys/drone-pypi) [![](https://goreportcard.com/report/github.com/xoxys/drone-pypi)](https://microbadger.com/images/xoxys/drone-pypi "Get your own image badge on microbadger.com") [![Build Status](https://travis-ci.org/xoxys/drone-pypi.svg?branch=master)](https://travis-ci.org/xoxys/drone-pypi)

Basic Drone Plugin for PyPi publishes - Based upon the starter project for creating Drone plugins.

## Build

Build the binary with the following commands:

```shell
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-pypi
docker build --rm -t plugins/drone-pypi .
```

## Usage

```shell
docker run --rm \
  -e PLUGIN_USERNAME=jdoe \
  -e PLUGIN_PASSWORD=my_secret \
  -e PLUGIN_SKIP_BUILD=false \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/drone-pypi
```
