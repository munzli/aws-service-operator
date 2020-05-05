# builder
FROM golang:alpine as builder

WORKDIR /go/src

RUN apk --no-cache --update add \
    make \
    git

COPY . .
RUN make build

# runner
FROM frolvlad/alpine-glibc

RUN apk --no-cache --update add openssl ca-certificates dumb-init
COPY --from=builder /go/src/aws-service-operator /usr/local/bin/

# setup environment
RUN addgroup -S aws-operator && \
    adduser -S aws-operator -G aws-operator

USER aws-operator

ENTRYPOINT ["dumb-init", "--", "aws-service-operator"]