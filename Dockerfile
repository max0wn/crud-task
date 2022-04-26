FROM golang:latest
ARG APP_DIR=app
COPY . /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}
RUN go mod download
RUN go build -o main ./cmd/main.go
CMD ["./main"]
