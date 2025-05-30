# The build stage
FROM golang:alpine AS builder

WORKDIR /app
COPY . .

RUN go generate ./... && go build -o monomatch ./api/main.go

# The run stage
FROM golang:alpine

WORKDIR /app
COPY --from=builder /app/monomatch .

EXPOSE 1982

CMD ["./monomatch"]