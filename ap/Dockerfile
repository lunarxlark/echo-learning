# golang for build {{{
FROM golang:alpine AS build

WORKDIR /go/src/github.com/lunarxlark/echo-study/ap
COPY . /go/src/github.com/lunarxlark/echo-study/ap/
#RUN go get -u github.com/golang/dep/...
#RUN dep ensure
RUN go build -tags netgo -o server server.go
# }}}

# alpine for execute {{{
#FROM alpine:3.9
FROM busybox

WORKDIR /work
COPY --from=build /go/src/github.com/lunarxlark/echo-study/ap/server /work/server
RUN mkdir /work/log

CMD ["/work/server"]
# }}}