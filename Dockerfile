FROM golang:latest

ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig
ENV GOPATH=/root/go
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN
ENV LD_LIBRARY_PATH=/usr/lib
ENV GIT_SSL_NO_VERIFY=1

WORKDIR $GOPATH/src/github.com/poc/logging-service

ADD . .

RUN go get -u github.com/beego/bee

ENTRYPOINT ["/bin/sh", "-c" , "echo 192.168.254.10   database-server >> /etc/hosts && echo 192.168.239.62   redis-ms-server >> /etc/hosts && exec java -jar ./botblocker.jar " ]

#CMD ["./main"]