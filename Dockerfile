FROM golang:1.22.2-bookworm

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go test ./...
RUN go build -buildvcs=false -o kanban-server ./cmd/server
EXPOSE 8081
CMD ["./kanban-server"]
