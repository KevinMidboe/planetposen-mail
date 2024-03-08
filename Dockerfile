# Build the project
FROM golang:1.19 as builder

LABEL org.opencontainers.image.source="https://github.com/KevinMidboe/planetposen-mail"

WORKDIR /go/src/github.com/kevinmidboe/planetposen-mail
ADD . .

RUN make build
# RUN make test

# Create production image for application with needed files
FROM golang:1.22.0-alpine

EXPOSE 8000

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/kevinmidboe/planetposen-mail .

CMD ["./main"]