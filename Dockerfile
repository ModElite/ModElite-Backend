FROM golang:1.23.2 AS builder
WORKDIR /app

ENV TZ="Asia/Bangkok"

COPY ./ ./
RUN go mod download

ARG CGO_ENABLED=0
ARG GOOS=linux

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init || exit 1

RUN go build -o /bin/app

FROM scratch AS runner
WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY .env ./
COPY --from=builder /app/docs/ ./docs/
COPY --from=builder /bin/app ./

EXPOSE 8080
CMD ["./app"]