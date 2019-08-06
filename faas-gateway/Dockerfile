FROM golang:latest
# go stuff
RUN go get -u github.com/kardianos/govendor && \
go get -u -v github.com/codeskyblue/fswatch && \
go get -u github.com/sparrc/go-ping
#RUN mkdir -p /go/src/github.com/faas/operations
WORKDIR /go/src/github.com/faas/gateway
ADD ./gatewayfd .
#COPY . /go/src/github.com/faas/operations
EXPOSE 80
RUN  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
#RUN go build /go/src/github.com/faas/ping/cmd/pingd/ping.go
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ping /go/src/github.com/#faas/ping/cmd/pingd/ping.go
#/go/src/github.com/faas/ping
CMD ["./gateway"]