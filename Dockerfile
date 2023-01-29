FROM golang:alpine AS builder

WORKDIR /go/src/github.com/zalgonoise/cloaki

COPY ./ ./

# this app's sqlite requires gcc 
RUN apk add build-base

RUN go mod download
RUN mkdir /build \
    && go build -o /build/cloaki . \
    && chmod +x /build/cloaki


FROM alpine:edge

RUN mkdir -p /cloaki/server

COPY --from=builder /build/cloaki /cloaki-server

CMD ["/cloaki-server"]