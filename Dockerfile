FROM golang:1.17.5-alpine3.14 as builder

WORKDIR /code 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/signalbox cmd/signalbox/main.go
RUN ls .

FROM alpine:3.15.0

WORKDIR /app

COPY --from=builder /code/bin/signalbox /usr/bin/

ENTRYPOINT ["signalbox"]