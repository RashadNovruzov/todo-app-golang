FROM golang:latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /go-docker-demo
EXPOSE 8080
CMD [ "/go-docker-demo" ]