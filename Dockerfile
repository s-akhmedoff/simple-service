FROM golang:latest as builder

RUN mkdir -p $GOPATH/src/microservice/simple
WORKDIR $GOPATH/src/microservice/simple

COPY . ./

RUN go mod tidy

RUN go get github.com/swaggo/swag/cmd/swag

CMD make swag

RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv /go/src/microservice/simple/bin/simple_service /

FROM alpine

COPY --from=builder simple .

RUN apk add --no-cache tzdata

ENV TZ Asia/Tashkent

ENTRYPOINT ["/simple"]
