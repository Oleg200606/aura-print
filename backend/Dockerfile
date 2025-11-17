FROM golang:25

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/srv ./cmd/server/main.go
EXPOSE 8081 
ENTRYPOINT ["/app/srv"]
