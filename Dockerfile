FROM golang:alpine AS builder
WORKDIR /app

ENV GOTOOLCHAIN=auto
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server cmd/api/main.go

FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Jakarta

COPY --from=builder /app/server .

EXPOSE 3011

CMD ["./server"]