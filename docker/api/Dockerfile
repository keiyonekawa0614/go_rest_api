FROM golang:1.14

ENV GO111MODULE=on

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build .

RUN go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]
