FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.work go.mod ./ ./
COPY pkg/go.mod pkg/go.mod
COPY apps/identity-server/go.mod apps/identity-server/go.mod
RUN go mod download

COPY pkg/ pkg/
COPY apps/identity-server/ apps/identity-server/
WORKDIR apps/identity-server/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/identity-server main.go

FROM gcr.io/distroless/base-debian12 AS publisher

WORKDIR /

COPY --from=builder /app/identity-server /identity-server
EXPOSE 80
ENTRYPOINT ["/identity-server"]


