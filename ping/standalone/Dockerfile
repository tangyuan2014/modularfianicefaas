# version: 1.0
FROM golang:latest
# go stuff
RUN go get -u github.com/kardianos/govendor && \
    go get -u github.com/sparrc/go-ping
WORKDIR /go/src/github.com/faas/ping
ADD ./cmd/pingd/ .
EXPOSE 8080
RUN go build  . 
CMD ["./ping"]