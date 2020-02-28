FROM ubuntu


RUN apt-get update
RUN apt-get install -y curl
RUN rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.4.2

RUN curl -sSL https://storage.googleapis.com/golang/go$GOLANG_VERSION.linux-amd64.tar.gz \
		| tar -v -C /usr/local -xz

ENV PATH /usr/local/go/bin:$PATH

RUN mkdir -p /go/src /go/bin && chmod -R 777 /go
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
COPY . /go/src
WORKDIR /go/src
EXPOSE 8080
RUN go build -o ./go/src .
CMD ["./go/src"]