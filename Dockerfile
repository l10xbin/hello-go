FROM golang:1.8 as builder

WORKDIR /go/src/github.com/l10xbin/hello-go/

COPY ./src/app.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/l10xbin/hello-go/app .

EXPOSE 8080

CMD ["./app"]