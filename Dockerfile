FROM golang:latest as builder

RUN mkdir -p $GOPATH/src/microservice/simple
WORKDIR $GOPATH/src/microservice/simple

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ./

#RUN go get github.com/swaggo/swag/cmd/swag

#CMD make swag-init

RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/simple /

FROM alpine

COPY --from=builder simple .

RUN apk add --no-cache tzdata

ENV TZ Asia/Tashkent

ENTRYPOINT ["/simple"]
