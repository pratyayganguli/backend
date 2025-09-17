## Golang Handbook: A Beginner to Advanced Guide

###### Last Draft Date: 14th December, 2025

**Go:** Golang is a general purpose programming language which is statically typed and compiled. It was designed at Google.

### Execution of a Go Program - 

As, Go is a compiled language when we run commands like `go build` or `go run` the source code is converted into a binary (machine level code) which contains the Go runtime and all the other dependency.

### Important concepts:

- [Memory Management](./Memory-Management.md)
- [Concurrency](./concurrency/Concurrency.md)
- [Error Handling]()
- [Mongo DB](./concurrency/Mongo.md)
- [Reflection](./Reflection.md)
- [Design Patterns](./Design-Patterns.md)

### Building the container image and running the container to run all the eligible tests and the code.

You will be requiring the [Dockerfile](#dockerfile) and the Docker run command to start the container. A container registry can also be set up for pushing the image from a remote repository and then starting the container by pulling from a remote repository.

*Alternate option:* You may also use [Docker Compose]()

**Command:**

```bash
docker build --build-arg GO_IMAGE=golang:1.25.1 --build-arg RUNTIME_IMAGE=alpine:latest -t github.com/pratyayganguli/backend/golang:latest .
```

#### Dockerfile: ####

```Dockerfile
#Global vars
ARG GO_IMAGE=golang:1.25.1
ARG RUNTIME_IMAGE=alpine:latest

#Building stage
FROM ${GO_IMAGE} as builder

RUN apt install git

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

#Running stage

FROM ${RUNTIME_IMAGE}

RUN apk add --no-cache ca-certificates

WORkDIR /root/

COPY --from = builder/app/main .

CMD["./main"]
```
