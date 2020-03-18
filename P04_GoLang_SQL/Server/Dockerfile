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

#Cài đặt Git & Pigo
RUN apt update
RUN apt-get install -y git
RUN cd go && go get -u -f github.com/esimov/pigo/cmd/pigo
RUN cd go/src/github.com/esimov/pigo/cmd/pigo && go install

#Cài đặt MysqlGolang
RUN cd go && go get -u github.com/go-sql-driver/mysql
RUN cd go/src/github.com/go-sql-driver/mysql && go install
RUN apt-get install -y net-tools

#Lấy những file cần thiết
COPY Server.go /go/src/server.go
RUN mkdir /go/src/Uploads

#Chạy chương trình
WORKDIR /go/src
EXPOSE 8080
RUN go build -o ./go/src .
CMD ["./go/src"]
