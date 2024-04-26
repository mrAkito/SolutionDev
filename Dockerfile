FROM --platform=linux/amd64 golang:1.21.4
#FROM golang:1.21.4-alpine
#RUN apk update && apk add git
WORKDIR /go/src
COPY . .
RUN go mod tidy
EXPOSE 8080

CMD ["go", "run", "main.go", "SaveGPS.go"]