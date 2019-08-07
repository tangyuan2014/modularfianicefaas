FROM golang:latest
# go stuff
RUN go get -u github.com/kardianos/govendor 
WORKDIR /go/src/github.com/faas/gateway
ADD ./gatewayd .

EXPOSE 80
RUN govendor init && \
    govendor fetch -v +m

RUN    go build .
#RUN go build /go/src/github.com/faas/ping/cmd/pingd/ping.go
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ping /go/src/github.com/#faas/ping/cmd/pingd/ping.go
#/go/src/github.com/faas/ping
CMD ["./gateway"]