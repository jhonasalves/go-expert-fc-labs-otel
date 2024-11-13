FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ARG SERVICE_NAME
ENV SERVICE_NAME=${SERVICE_NAME}

COPY ${SERVICE_NAME}/ ${SERVICE_NAME}/

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ${SERVICE_NAME}/cmd/server/main.go

FROM scratch

ARG SERVICE_NAME

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/${SERVICE_NAME}/cmd/server/.env .

CMD ["./server"]
