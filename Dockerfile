# build go binary
FROM golang AS builder
WORKDIR /go/src/github.com/bugitt/cloudapi-daemon-trigger
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/cloudapi-daemon-trigger

# build final image
FROM alpine
COPY --from=builder /go/bin/cloudapi-daemon-trigger /usr/local/bin/cloudapi-daemon-trigger
ENTRYPOINT ["/usr/local/bin/cloudapi-daemon-trigger"]
ENV TIMEZONE=Asia/Shanghai