
# Building stage

FROM golang:1.24-alpine AS builder

RUN apk --no-cache add git

WORKDIR /app

COPY go.mod go.sum ./


COPY . .

RUN go build -o main .

# Running stage

FROM alpine:latest 

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]