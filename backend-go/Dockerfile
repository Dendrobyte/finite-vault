# syntax=docker/dockerfile:1

FROM golang:1.23

WORKDIR /usr/src/finite-vault
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

# Build
RUN go build -v -o /usr/local/bin/finite-vault

EXPOSE 5001

# Run
CMD ["finite-vault"]