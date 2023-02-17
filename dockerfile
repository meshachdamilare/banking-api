FROM golang:1.19-alpine3.17

WORKDIR go/src/banking-api

COPY . .

COPY go.mod ./

RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]