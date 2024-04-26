# Builder stage
FROM golang:1.22-alpine3.19 as builder

# Download ca-certificates to be copied into the final image later
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

# Build the binary
ADD . /src
WORKDIR /src
RUN go get
RUN go build -o hastecat


# Start a new build stage
FROM scratch

# Copy in HasteCat binary
COPY --from=builder src/hastecat /hastecat

# Copy in CA certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Set the entrypoint
ENTRYPOINT ["/hastecat"]

