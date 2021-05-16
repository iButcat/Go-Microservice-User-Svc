FROM golang:1.16

WORKDIR /app

COPY go.mod .
RUN go mod download usersvc

COPY . .

RUN go build -o main .
ENTRYPOINT ["/app/main"]