# The build stage
FROM golang:latest-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o monomatch /app/api/main.go

# The run stage
FROM golang:latest-alpine
WORKDIR /app
COPY --from=builder /app/monomatch .
EXPOSE 1982
CMD ["./monomatch"]