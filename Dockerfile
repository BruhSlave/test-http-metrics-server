FROM golang:1.26.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./ 
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o goMetricsApp ./cmd/server/


FROM alpine:3.21

WORKDIR /app

RUN apk upgrade --no-cache

COPY --from=builder /app/goMetricsApp .

EXPOSE 8090

CMD [ "./goMetricsApp" ]
