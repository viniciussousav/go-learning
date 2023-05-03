FROM golang:latest

ENV CGO_ENABLED=1
ENV CGO_CFLAGS=-I/usr/local/include
ENV CGO_LDFLAGS=-L/usr/local/lib

RUN apt-get update && \
    apt-get install -y wget && \
    wget https://github.com/edenhill/librdkafka/archive/v1.7.0.tar.gz && \
    tar -xzf v1.7.0.tar.gz && \
    cd librdkafka-1.7.0 && \
    ./configure --prefix /usr/local && \
    make && \
    make install && \
    cd .. && \
    rm -rf librdkafka-1.7.0 v1.7.0.tar.gz

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "main.go"]
