FROM golang:1.21 as gobuilder
WORKDIR /build

COPY ./ /build

ENV CGO_ENABLED=0 GOOS=linux GOOS=linux GOARCH=amd64 GOPATH=/build/go

RUN go build -o ./test-app ./app/


FROM alpine:3.15.0

WORKDIR /app/

COPY --from=gobuilder /build/test-app .

CMD ["./test-app"]
