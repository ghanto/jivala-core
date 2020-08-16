FROM golang:1.14.5-alpine3.12 as builder

WORKDIR /app
COPY go.mod go.sum ./
COPY cmd/jivala-core/main.go ./
RUN go mod download
COPY . .

# Build binary
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/main .

CMD ["/main"]