FROM golang:latest as builder
WORKDIR /go/src/github.com/jakobvarmose/redirect/
COPY ./ ./
RUN CGO_ENABLED=0 go build -o /root/app

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /root/app app
CMD ["./app"]
EXPOSE 80
