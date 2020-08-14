FROM ysicing/alpine

COPY dist/crtools_linux_amd64 /usr/local/bin/crtools

COPY hack/docker/entrypoint.sh /entrypoint.sh

RUN chmod +x /usr/local/bin/crtools /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

CMD ["crtools", "-h"]