
# Building stage

FROM golang:1.25-alpine AS builder

RUN apk --no-cache add git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# Running stage

FROM alpine:latest 

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]