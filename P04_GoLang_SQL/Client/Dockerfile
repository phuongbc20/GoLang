#Cài đặt ubuntu và các gói cần thiết
FROM ubuntu

RUN apt-get update
RUN apt-get install -y curl
RUN rm -rf /var/lib/apt/lists/*

#Cài đặt golang
ENV GOLANG_VERSION 1.4.2
RUN curl -sSL https://dl.google.com/go/go1.14.linux-amd64.tar.gz \
		| tar -v -C /usr/local -xz
ENV PATH /usr/local/go/bin:$PATH
RUN mkdir -p /go/src /go/bin && chmod -R 777 /go
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

#Lấy những file cần thiết
COPY Client.go /go/src/Client.go
COPY /Test /go/src
RUN mkdir /go/src/Image
RUN mkdir /go/src/Json

#Chạy chương trình
WORKDIR /go/src
EXPOSE 8080
RUN go build -o ./go/src .
CMD ["./go/src"]