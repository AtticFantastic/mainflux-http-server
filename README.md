# Mainflux HTTP Server

[![License](https://img.shields.io/badge/license-Apache%20v2.0-blue.svg)](LICENSE)
[![Build Status](https://travis-ci.org/Mainflux/mainflux-http-server.svg?branch=master)](https://travis-ci.org/Mainflux/mainflux-http-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/Mainflux/mainflux-http-server)](https://goreportcard.com/report/github.com/Mainflux/mainflux-http-server)
[![Join the chat at https://gitter.im/Mainflux/mainflux](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/Mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

HTTP API Microservice for Mainflux IoT Platform.

### Installation
Mainflux HTTP Sevrer can be installed either via official Docker image (maintained by Mainflux team and available on [Mainflux DockerHub](https://hub.docker.com/u/mainflux)) or by downloading and compiling the source code from GitHub.

#### Docker
```bash
# NATS prerequisite
docker pull apcera/gnatsd
docker run --name=nats -it apcera/gnatsd
# Mainflux HTTP Server
docker pull mainflux/mainflux-http-server
docker run -p 7070:7070 --link=nats:nats -it mainflux/mainflux-http-server
```
#### Code
##### Prerequisite
If not set already, please set your `GOPATH` and `GOBIN` environment variables. For example:
```bash
mkdir -p ~/go
export GOPATH=~/go
export GOBIN=$GOPATH/bin
```

##### Get the code
Use [`go`](https://golang.org/cmd/go/) tool to "get" (i.e. fetch and build) `mainflux-http-server` package:
```bash
go get github.com/mainflux/mainflux-http-server
```

This will download the code to `$GOPATH/src/github.com/mainflux/mainflux-http-server` directory,
and then compile it and install the binary in `$GOBIN` directory.

Now you can run the server:
```bash
$GOBIN/mainflux-http-server
```

Please note that the binary `mainflux-http-server` expects to find configuration file `config.yml` in
direcotry provided by `MAINFLUX_HTTP_SERVER_CONFIG_DIR` if this variable is set. Otherwise it looks for `config.yml`
in `$GOPATH/src/github.com/mainflux/mainflux-http-server`.

Note also that using `go get` is prefered than out-of-gopath code fetch by cloning the git repo like this:
```
git clone https://github.com/Mainflux/mainflux-http-server && cd mainflux-http-server
go get
go build
MAINFLUX_HTTP_SERVER_CONFIG_DIR=. ./mainflux-http-server
```

#### Dependencies
Mainflux HTTP Server is connected to `NATS` on southbound interface.

Following diagram illustrates the architecture:
![Mainflux Arch](https://github.com/Mainflux/mainflux-doc/blob/master/mermaid/arch.png)

Upon booting it will expect that [NATS](https://github.com/nats-io/gnatsd) is up and running. This can be obtained either by fetching and compiling `NATS` source:
```bash
go get github.com/nats-io/gnatsd
$GOBIN/gnatsd
```
or pulling the [official `NATS` Docker container](https://hub.docker.com/r/apcera/gnatsd/)
```
docker pull apcera/gnatsd
docker run -p 4222:4222 -it apcera/gnatsd
```
> N.B. If you used official Mainflux
> [`docker-compose.yml`](https://github.com/Mainflux/mainflux/blob/master/docker-compose.yml) to
> fetch the images and run the composition, then `NATS` image is already downloaded on your host and `mainflux-nats`
> container is already created. In that case all you have to run is `docker start mainflux-nats`

After starting `NATS` (from the binary or from the continer), we can now start `mainflux-http-server` in separate terminal:
```
MAINFLUX_HTTP_SERVER_CONFIG_DIR=. ./mainflux-http-server
```
Note that since we are starting `NATS` now with ports mapped to `localhost` we must customize [config.yml](config.yml) and change `NATS` hostname to `localhost` as this is where `mainflux-http-server` will expect the `NATS` service to be available.

> N.B. `NATS` host name in the `config.yml` is defined as `nats`,
> which corresponds to the name of the service in [`docker-compose.yml`](https://github.com/Mainflux/mainflux/blob/master/docker-compose.yml).
> If you are running `mainflux-http-server` locally (and not via `docker-compose`), then you must change
> NATS hostname to `localhost` - this will work weather you are running `NATS` as a `gnatsd` compiled from source
> or as a `docker run mainflux-nats`

### Documentation
Development documentation can be found on our [Mainflux GitHub Wiki](https://github.com/Mainflux/mainflux/wiki).

Swagger-generated API reference can be foud at [http://mainflux.com/apidoc](http://mainflux.com/apidoc).

### Community
#### Mailing lists
- [mainflux-dev](https://groups.google.com/forum/#!forum/mainflux-dev) - developers related. This is discussion about development of Mainflux IoT cloud itself.
- [mainflux-user](https://groups.google.com/forum/#!forum/mainflux-user) - general discussion and support. If you do not participate in development of Mainflux cloud infrastructure, this is probably what you're looking for.

#### IRC
[Mainflux Gitter](https://gitter.im/Mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

#### Twitter
[@mainflux](https://twitter.com/mainflux)

### License
[Apache License, version 2.0](LICENSE)
