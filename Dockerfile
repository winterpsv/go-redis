# 1 choose a compiler OS
FROM golang:1.18-alpine as builder

RUN go version
ENV GOPATH=/

# 2 copy all the source files
WORKDIR /workspace
COPY . .

# 3 get dependency
RUN go mod download

# 4 build the GO program
RUN go build -o task3_4 ./cmd/app/main.go

# 5 choose a runtime OS
FROM alpine:latest as production
ARG ENV
WORKDIR /

# 6 copy from builder the GO executable file
COPY --from=builder /workspace/task3_4 .

# 7 execute the program upon start
CMD ["./task3_4"]