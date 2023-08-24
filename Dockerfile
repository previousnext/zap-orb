FROM golang:1.21-alpine as build
RUN apk add --no-cache ca-certificates
ADD . /go/src/github.com/previousnext/zap-orb
WORKDIR /go/src/github.com/previousnext/zap-orb
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build -a -o bin/zap-slack-notify github.com/previousnext/zap-orb/cmd/zap-slack-notify

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/github.com/previousnext/zap-orb/bin/zap-slack-notify /usr/local/bin/zap-slack-notify
ENTRYPOINT ["/usr/local/bin/zap-slack-notify"]