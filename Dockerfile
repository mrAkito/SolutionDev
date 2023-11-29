FROM golang:latest

WORKDIR /src

COPY . .

RUN go mod download
# RUN apk add --no-cache gcc
# RUN apk add --no-cache musl-dev


CMD ["go", "run", "./main.go"]
