FROM golang:1.11.5

WORKDIR /go/src/ap
COPY server.go server.go
COPY Gopkg.toml Gopkg.toml

RUN go get -u github.com/golang/dep/...
RUN mkdir log
RUN dep ensure

CMD ["go", "run", "server.go"]
