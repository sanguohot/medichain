FROM golang:1.11.0-alpine3.7
RUN apk add --no-cache git\
 build-base \
 gcc \
 musl-dev \
 && rm -rf /var/cache/apk/*
WORKDIR /opt/medichain
COPY . .
RUN ls -al /opt/medichain
ENV GO11MODULE=on
#禁用cgo
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main -v
#使用cgo并默认使用动态链接
#RUN go build -o main -v main.go
#使用cgo并使用静态链接
RUN go build --ldflags "-linkmode external -extldflags -static" -a -o main -v

FROM busybox:1.28
WORKDIR /root/
ENV MEDICHAIN_PATH=/root
COPY ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /opt/medichain/main .
EXPOSE 8443/tcp
ENTRYPOINT ["./main"]
