
# Building stage
ARG GO_IMAGE=golang:1.25.1

ARG RUNTIME_IMAGE=alpine:latest

FROM ${GO_IMAGE} AS builder

RUN apt install git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# Running stage
FROM ${RUNTIME_IMAGE}

RUN apk --no-cache add ca-certificates 

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]