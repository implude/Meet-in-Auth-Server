FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN mkdir -p /go/src/pentag.kr/Meet-n-Auth-Server

ENV GOPATH=/go/src
WORKDIR /go/src/pentag.kr/Meet-n-Auth-Server

COPY go.mod go.sum main.go ./
RUN go mod download

COPY ./ ./
RUN go build -o main .
WORKDIR /dist
RUN cp -r /go/src/pentag.kr/Meet-n-Auth-Server/main ./main
RUN cp -r /go/src/pentag.kr/Meet-n-Auth-Server/.env ./.env

FROM scratch
COPY --from=builder /dist/main ./main
COPY --from=builder /dist/.env ./.env
ENTRYPOINT ["./main"]