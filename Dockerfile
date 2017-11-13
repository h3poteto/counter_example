FROM h3poteto/golang:1.9.1

COPY --chown=go:go . /go/src/github.com/h3poteto/counter

USER go
WORKDIR /go/src/github.com/h3poteto/counter

RUN set -x \
    && go build

CMD "./counter"


