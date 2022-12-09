FROM golang:1.19-alpine
RUN apk add build-base
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /app/fetch cmd/app/main.go
ENTRYPOINT ["/app/fetch"]