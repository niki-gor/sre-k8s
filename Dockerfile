FROM golang:1.20.4-alpine AS builder

WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./api.go ./
RUN go build api.go

FROM scratch
COPY --from=builder /app/api ./api
