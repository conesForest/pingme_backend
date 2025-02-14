FROM golang:1.23-alpine

WORKDIR /pingme-backend

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main ./cmd/main.go

EXPOSE 8080

ENTRYPOINT ["./main"]
