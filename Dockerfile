FROM golang:latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-api-gin

EXPOSE 8080

CMD [ "/go-api-gin" ]

