FROM golang as build-env
WORKDIR /go/src/sshesame-log-to-mongodb
ADD . /go/src/sshesame-log-to-mongodb
RUN go build -o /go/bin/sshesame-log-to-mongodb
FROM gcr.io/distroless/base
VOLUME /data
CMD ["/sshesame-log-to-mongodb"]