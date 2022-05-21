FROM golang:alpine AS development

#Set working dir
WORKDIR /app

#Get dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# copy all
COPY . .

EXPOSE 3333

# run application
CMD go run ./cmd .

FROM golang:alpine AS builder

#build variables
ENV GOOS linux
ENV CGO_ENABLED 0

#Set working dir
WORKDIR /app

#Get dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

#build application
RUN go build -o app ./cmd

FROM alpine:latest AS production
RUN apk add --no-cache ca-certificates
COPY --from=builder app .
EXPOSE 3333
CMD ./app