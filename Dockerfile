# Only the instructions RUN, COPY, ADD create layers. Other instructions create temporary intermediate images, and don’t increase the size of the build.

FROM golang:alpine AS build-stage

ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

# https://docs.docker.com/build/guide/layers/
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /app/nautilus

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/nautilus /app/nautilus

EXPOSE 3000

ENTRYPOINT ["/app/nautilus"]
