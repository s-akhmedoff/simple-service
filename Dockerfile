FROM golang:latest as builder

#
RUN mkdir -p $GOPATH/src/microservice/simple
WORKDIR $GOPATH/src/microservice/simple

#Copy go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# Copy the local package files to the container's workspace.
COPY . ./

## Get swaggo
#RUN go get github.com/swaggo/swag/cmd/swag

##Generate swagger documentation
#CMD make swag-init

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/simple /

FROM alpine
COPY --from=builder simple .
RUN apk add --no-cache tzdata
ENV TZ Asia/Tashkent


ENTRYPOINT ["/simple"]
