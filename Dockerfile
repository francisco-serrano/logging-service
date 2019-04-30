FROM golang:latest

ENV GOPATH=/root/go
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN

WORKDIR $GOPATH/src/github.com/poc/logging-service

ADD . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/beego/bee
#RUN go get -u github.com/astaxie/beego

RUN dep ensure -v