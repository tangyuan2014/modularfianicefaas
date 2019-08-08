# version: 1.0
FROM golang:latest
# go stuff
RUN go get -u github.com/kardianos/govendor 
WORKDIR /go/src/github.com/faas/operations
COPY ./cmd/operationsd .
COPY ./Operations /go/src/modularfianicefaas/mathoperations/Operations
EXPOSE 8080
RUN   go build .
CMD ["./operations"]
