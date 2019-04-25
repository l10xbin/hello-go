FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/github.com/l10xbin/hello-go/src

COPY src/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/l10xbin/hello-go/src/app .

EXPOSE 8080

CMD ["./app"]