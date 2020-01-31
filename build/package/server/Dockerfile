# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang v1.11 base image
FROM golang:1.13.4-alpine as build

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/pepeunlimited/rpc-starter-kit

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/server ./cmd/server/
###############################################
######## Start a new stage from scratch #######
###############################################
FROM alpine:latest as final

LABEL description="Show case of the using docker hub."
# This is intended to be run from the hooks/build script, which sets this
# appropriately based on git tags.
ARG version
LABEL com.pepeunlimited.rpcstarterkit.version=$version

# Add Maintainer Info
LABEL maintainer="Simo Ala-Kotila <simo.alakotila@gmail.com>"
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=build /go/bin/server .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ./server