FROM golang:1.26.8-bookworm AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /out/letseat-api ./cmd/api

FROM debian:bookworm-slim

COPY --from=build /out/letseat-api /usr/local/bin/letseat-api

ENTRYPOINT ["/usr/local/bin/letseat-api"]



# FROM golang:1.26.8-bookworm

# WORKDIR /src

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .
# RUN CGO_ENABLED=0 go build -o /usr/local/bin/letseat-api ./cmd/api

# ENTRYPOINT ["/usr/local/bin/letseat-api"]