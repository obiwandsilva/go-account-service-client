FROM golang:1.14.3-alpine3.11

WORKDIR /opt/go-account-service

COPY . .

RUN ["go", "build", "cmd/account-service-client/main.go"]

EXPOSE 80

CMD [ "./main", "80" ]

# It's needed to pass the env var ACCOUNT_SERVICE_HOST when running the container
