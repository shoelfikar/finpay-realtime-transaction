FROM golang:1.23.7-alpine


ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

WORKDIR '/finpay-transaction'

COPY go.mod .
COPY go.sum .

COPY . .

RUN go mod download
RUN go mod vendor
RUN go mod verify


# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
RUN go build -o finpay
EXPOSE 8000

CMD ["./finpay"]