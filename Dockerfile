FROM golang:1.11-alpine3.7

WORKDIR /go/src/github.com/pospome/cicd
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["cicd"]
