FROM golang:alpine

WORKDIR /go/src/app

COPY ./go.mod ./go.sum ./

RUN go mod download && go mod verify

COPY ./ ./

RUN go build -v -o /go/bin/app ./cmd/

EXPOSE 8008

CMD [ "app" ]
