FROM golang:latest
# go stuff
RUN go get -u github.com/kardianos/govendor && \
    go get -u -v github.com/codeskyblue/fswatch

RUN mkdir -p /go/src/github.com/tangyuan2014/modularfianicefaas/faas-gateway
WORKDIR /go/src/github.com/tangyuan2014/modularfianicefaas/faas-gateway
EXPOSE 80
COPY ./fsw.yml /fsw.yml
CMD govendor init && govendor fetch -v +m && fswatch -config /fsw.yml