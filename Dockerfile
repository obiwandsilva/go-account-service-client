FROM golang:1.14.3-alpine3.11 AS builder
WORKDIR /tmp/go-account-service-client
COPY . .
RUN ["go", "build", "cmd/account-service-client/main.go"]

FROM alpine:3.11
WORKDIR /opt/account-service-client
COPY --from=builder /tmp/go-account-service-client/main .
EXPOSE 3000
CMD [ "./main", "3000" ]

# Set the env var ACCOUNT_SERVICE_HOST when running a container
