FROM golang:1.10
MAINTAINER Anders Romin <elektromin@hotmail.com>

WORKDIR /go/src/github.com/elektromin/gohello
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["gohello"]