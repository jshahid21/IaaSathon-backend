FROM golang:1.10.3 as builder
WORKDIR /go/src/github.com/schmidtp0740/iaas-be
RUN go get github.com/gorilla/mux github.com/rs/cors gopkg.in/goracle.v2
COPY main.go .
RUN  go build -o app .

FROM ubuntu:latest
WORKDIR /opt/oracle
RUN apt update \
    && apt install -y unzip libaio1
COPY --from=builder /go/src/github.com/schmidtp0740/iaas-be/app .
COPY instantclient-basiclite-linux.x64-12.2.0.1.0.zip .
RUN unzip instantclient-basiclite-linux.x64-12.2.0.1.0.zip 
ENV LD_LIBRARY_PATH=/opt/oracle/instantclient_12_2:$LD_LIBRARY_PATH
CMD [ "./app" ]


