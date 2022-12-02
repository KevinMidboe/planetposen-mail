# Build the project
FROM golang:1.19 as builder

WORKDIR /go/src/github.com/kevinmidboe/planetposen-mail
ADD . .

RUN make build
# RUN make test

# Create production image for application with needed files
FROM golang:1.19-alpine

EXPOSE 8000

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/kevinmidboe/planetposen-mail .

CMD ["./main"]