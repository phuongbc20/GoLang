#Cài đặt ubuntu làm môi trường 
FROM ubuntu
RUN apt-get update
RUN apt-get install -y curl
RUN rm -rf /var/lib/apt/lists/*

#Cài đặt golang
ENV GOLANG_VERSION 1.4.2
RUN curl -sSL https://storage.googleapis.com/golang/go$GOLANG_VERSION.linux-amd64.tar.gz \
		| tar -v -C /usr/local -xz
ENV PATH /usr/local/go/bin:$PATH

#Tạo thư mục để làm GOPATH
RUN mkdir -p /go/src /go/bin && chmod -R 777 /go
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

#Lấy tập tin server.go từ thư mục hiện tại copy qua cho ubuntu
COPY Server.go /go/src/server.go
RUN mkdir /go/src/Uploads

#Chuyển thư mục làm việc đến /go/src
WORKDIR /go/src

#Chọn cổng giao tiếp của docker là 8080
EXPOSE 8080

#build và run chương trình
RUN go build -o ./go/src .
CMD ["./go/src"]