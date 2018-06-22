FROM alpine:latest

WORKDIR /opt/oracle

COPY main.go .

COPY instantclient-basiclite-linux.x64-12.2.0.1.0.zip .

RUN apk update \
    && apk add unzip libaio git go  \
    && unzip instantclient-basiclite-linux.x64-12.2.0.1.0.zip \
    && mkdir -p /etc/ld.so.conf.d \
    && go get github.com/gorilla/mux github.com/rs/cors gopkg.in/goracle.v2 \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

ENV LD_LIBRARY_PATH=/opt/oracle/instantclient_12_2:$LD_LIBRARY_PATH

#receiving error with c compilation
CMD [ "./app" ]


