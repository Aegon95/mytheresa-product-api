FROM golang:1.17.3-alpine3.14 AS build_base

ENV CGO_ENABLED=1
ENV GO111MODULE=on
RUN apk add --no-cache git  git gcc g++

# Set the Current Working Directory inside the container
WORKDIR /src

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/app ./cmd/api/main.go

# Start fresh from a smaller image
FROM alpine:3.14.3
RUN apk add ca-certificates

WORKDIR /app

COPY --from=build_base /src/out/app /app/api
COPY --from=build_base /src/config /app/config
COPY --from=build_base /src/migrations /app/migrations

RUN chmod +x api

# This container exposes port 3000 to the outside world
EXPOSE 3000

# Run the binary program produced by `go install`
ENTRYPOINT ./api