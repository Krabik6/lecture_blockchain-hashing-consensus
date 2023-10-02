# Stage 1: Build the 'simplehash' application
FROM golang:1.16-alpine as simplehash
WORKDIR /app
COPY ./simplehash/ .
RUN go build -o main .
CMD ["./main"]

# Stage 2: Build the 'sha256' application
FROM golang:1.16-alpine as sha256
WORKDIR /app
COPY ./sha256/ .
RUN go build -o main .
CMD ["./main"]

# Stage 3: Build the 'mining' application
FROM golang:1.16-alpine as maining
WORKDIR /app
COPY mining/ .
RUN go build -o main .
CMD ["./main"]
