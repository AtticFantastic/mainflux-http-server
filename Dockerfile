###
# Mainflux Dockerfile
###

FROM golang:alpine
MAINTAINER Mainflux

ENV MAINFLUX_NATS_PORT=4222
ENV MAINFLUX_HTTP_SERVER_CONFIG_DIR=/config/http

###
# Install
###

RUN apk update && apk add git && rm -rf /var/cache/apk/*

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/mainflux/mainflux-http-server

RUN mkdir -p /config/http
COPY config.yml /config/http

# Get and install the dependencies
RUN go get github.com/mainflux/mainflux-http-server

###
# Run main command from entrypoint and parameters in CMD[]
###
CMD [""]

# Run mainflux command by default when the container starts.
ENTRYPOINT /go/bin/mainflux-http-server

EXPOSE $MAINFLUX_NATS_PORT
