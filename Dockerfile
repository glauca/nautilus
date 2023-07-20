FROM golang:alpine

WORKDIR /app

COPY . ./

RUN go env -w GOPROXY=https://goproxy.cn
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /nautilus

EXPOSE 3000

CMD [ "/nautilus" ]
