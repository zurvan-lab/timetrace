FROM golang:1.21.5-alpine3.18 as builder

RUN apk add --no-cache git gmp-dev build-base g++ openssl-dev
ADD . /ttrace

# Building ttrace-cli
RUN cd /ttrace && \
    go build -ldflags "-s -w" -trimpath -o ./build/ttrace ./cmd


## Copy binary files from builder into second container
FROM golang:1.21.5-alpine3.18

COPY --from=builder /timetrace/build/ttrace /usr/bin

ENV WORKING_DIR "/ttrace"

VOLUME $WORKING_DIR
WORKDIR $WORKING_DIR
