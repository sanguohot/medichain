FROM golang:1.11.0-alpine3.7
RUN apk add --no-cache git\
 build-base \
 && rm -rf /var/cache/apk/*
WORKDIR /opt/medichain
COPY . .
RUN ls -al /opt/medichain
#RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main -v main.go
ENV GO11MODULE=on
RUN go build -o main -v main.go

FROM busybox:1.28
WORKDIR /root/
ENV MEDICHAIN_PATH=/root
COPY ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /opt/medichain/main .
EXPOSE 8443/tcp
ENTRYPOINT ["./main"]
