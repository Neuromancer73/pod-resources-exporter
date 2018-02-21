FROM golang:latest
RUN mkdir /app
ADD . /app/

RUN export GOPATH="/tmp/"

RUN go get k8s.io/client-go/kubernetes

WORKDIR /app
RUN go build -o main .

RUN rm -rf /tmp/*

CMD ["/app/main"]