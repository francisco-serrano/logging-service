FROM golang:latest

ENV GOPATH=/root/go
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN

WORKDIR $GOPATH/src/github.com/poc/logging-service

ADD . .

#RUN dep ensure
RUN go build main.go

CMD ["./main"]