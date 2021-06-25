FROM scratch
ENTRYPOINT ["/growth-pull"]
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY growth-pull /
