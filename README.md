
# go-account-service-client

This is an http web client implemented in Go and using Go Modules.  

# Usage

Two requirements:

-  [go-account-service](https://github.com/obiwandsilva/go-account-service) running.

- **ACCOUNT_SERVICE_HOST**environment variable configured with its address.

You can either start the application using a binary generated by `go build cmd/account-service-client/main.go` or by running a container with its image `wandsilva/go-account-service-client:1.0.0`. For the later, you can use the `Dockerfile` or get the image from Docker Hub.

If your are going to run using the binary, use a port number as the first and only argument for it. Example: `./main 4000`.

If you are going to run on a container, don't forget to expose and bind a port to the container. The default port used by the container is `3000`, so you could use `docker container run -p 80:3000 -e ACCOUNT_SERVICE_HOST=localhost:3001 wandsilva/go-account-service-client:1.0.0`.

# API

It's a simple app for studies purposes only, so it has only one valid endpoint.

## Get account info

`GET` on `/123456`
