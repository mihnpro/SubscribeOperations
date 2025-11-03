FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .


FROM alpine:latest


RUN apk --no-cache add ca-certificates tzdata


RUN addgroup -S app && adduser -S app -G app

WORKDIR /app


COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs/
COPY --from=builder /app/db/migrations/sql ./db/migrations/sql/


RUN chown -R app:app /app


USER app

EXPOSE 8080

CMD ["./main"]