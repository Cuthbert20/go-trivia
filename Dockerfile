FROM golang:1.19.0

WORKDIR /usr/src/app
#using cosmtrek/air for hot reloading
RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy
