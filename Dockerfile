FROM frolvlad/alpine-glibc

RUN apk --no-cache add openssl ca-certificates dumb-init
COPY aws-service-operator /usr/local/bin/

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/usr/local/bin/aws-service-operator"]