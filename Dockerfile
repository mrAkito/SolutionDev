FROM golang:1.22.0-alpine3.19

WORKDIR /try/backend
COPY . .

RUN go install github.com/cosmtrek/air@latest
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
RUN go mod tidy

CMD [ "air" ]