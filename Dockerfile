FROM golang:alpine AS build-stage

WORKDIR /app

COPY . ./

RUN go env -w GOPROXY=https://goproxy.cn \
    && go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -o /app/nautilus

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/nautilus /app/nautilus

EXPOSE 3000

ENTRYPOINT ["/app/nautilus"]
