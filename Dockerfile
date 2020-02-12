FROM frolvlad/alpine-glibc

RUN apk --no-cache add openssl ca-certificates dumb-init
COPY aws-service-operator /usr/local/bin/

ENTRYPOINT ["dumb-init", "--", "aws-service-operator"]