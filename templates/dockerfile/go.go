package dockerfile

import "fmt"

type GoTemplate struct{}

func NewGoTemplate() *GoTemplate {
	return &GoTemplate{}
}

func (t *GoTemplate) Execute(port int, appname, url string) string {
	return fmt.Sprintf(`FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y wget git gcc

RUN wget -P /tmp https://dl.google.com/go/go1.17.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf /tmp/go1.17.linux-amd64.tar.gz
RUN rm /tmp/go1.17.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR /data

RUN git clone %s /data/app

WORKDIR /data/app

RUN go mod download

COPY .env .env

RUN go build -o main .

EXPOSE %d

CMD ["./main"]`, url, port)
}
