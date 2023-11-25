FROM golang as build-env
WORKDIR /go/src/sshesame-log-to-mongodb
ADD . /go/src/sshesame-log-to-mongodb
RUN go build -o /go/bin/sshesame-log-to-mongodb
FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/sshesame-log-to-mongodb /
CMD ["/sshesame-log-to-mongodb", "-config", "/config.yaml", "-data_dir", "/data"]
