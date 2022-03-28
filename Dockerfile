# global arg
ARG APP_NAME='sneed'

# FIRST STAGE
FROM golang:1.17 AS builder

WORKDIR /app
# make global arg available
ARG APP_NAME

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go ./

# CGO_ENABLED is mandatory as alpine doesn't have the dylib files required by the compiled binary
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/$APP_NAME

# SECOND STAGE
FROM alpine:3.14.2

WORKDIR /app
# make global arg available
ARG APP_NAME

COPY --from=builder /app/bin/$APP_NAME /app

# Pass into the environment for the cmd command
ENV APP_NAME=$APP_NAME

CMD ./$APP_NAME
