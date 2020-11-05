FROM amazon/aws-cli
COPY entrypoint.sh /usr/bin/
ENTRYPOINT ["entrypoint.sh"]
