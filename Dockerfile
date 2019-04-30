FROM golang:latest

ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig
ENV GOPATH=/root/go
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN
ENV LD_LIBRARY_PATH=/usr/lib

WORKDIR $GOPATH/src/github.com/poc/logging-service

ADD . .

#RUN go get -u github.com/beego/bee
#RUN go get -u github.com/astaxie/beego