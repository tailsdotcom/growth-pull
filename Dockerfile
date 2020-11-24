FROM golang
COPY . /app
RUN cd /app; \
    CGO_ENABLED=0 go build -a -ldflags "-s -w" .; \
    mkdir -p etc/ssl/certs; \
    rm go.mod go.sum main.go; \
    curl \
        https://www.amazontrust.com/repository/AmazonRootCA1.pem \
        https://www.amazontrust.com/repository/AmazonRootCA2.pem \
        https://www.amazontrust.com/repository/AmazonRootCA3.pem \
        https://www.amazontrust.com/repository/AmazonRootCA4.pem \
        https://www.amazontrust.com/repository/SFSRootCAG2.pem \
    > etc/ssl/certs/ca-certificates.crt
RUN apt-get update \
 && apt-get install -y upx-ucl \
 && upx --ultra-brute /app/growth-pull
FROM scratch
COPY --from=0 /app /
ENTRYPOINT ["/growth-pull"]
