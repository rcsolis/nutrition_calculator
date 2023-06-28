FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/ ./...

WORKDIR /app/bin
RUN chmod 777 nutrition_calculator

CMD [ "./nutrition_calculator" ]

