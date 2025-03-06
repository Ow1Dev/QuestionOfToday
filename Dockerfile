FROM docker.io/library/golang:1.24-alpine3.21 AS build-env

RUN apk --no-cache add \
    build-base \
    nodejs \
    npm \
    && rm -rf /var/cache/apk/*

COPY . ${GOPATH}/src/code.question-of-today/qot
WORKDIR ${GOPATH}/src/code.question-of-today/qot

RUN make clean-all build

FROM docker.io/library/alpine:3.21

EXPOSE 3000

ENTRYPOINT ["/app/qot/questionofday"]

COPY --from=build-env /go/src/code.question-of-today/qot/questionofday /app/qot/questionofday
