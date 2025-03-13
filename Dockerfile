FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o url-shortner cmd/main.go

FROM alpine

ENV PORT="8080"
ENV ADDRESS="0.0.0.0"
COPY --from=builder /app/url-shortner .
EXPOSE $PORT

CMD [ "./url-shortner" ]