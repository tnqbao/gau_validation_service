FROM golang:1.23-alpine AS builder
WORKDIR /gau_validation 
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env .
RUN go build -o main .

FROM alpine:latest
WORKDIR /gau_validation
COPY --from=builder /gau_validation/main .
COPY .env .
EXPOSE 8081
CMD ["./main"]
