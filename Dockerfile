FROM golang:1.21.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /api-server-go

EXPOSE 3000

CMD ["/api-server-go"]