# syntax=docker/dockerfile:1

FROM golang:1.20-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /alertmanager-prometheus-wrapper

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage
ENV GIN_MODE=release

WORKDIR /

COPY --from=build-stage /alertmanager-prometheus-wrapper /alertmanager-prometheus-wrapper

EXPOSE 80

USER nonroot:nonroot

ENTRYPOINT ["/alertmanager-prometheus-wrapper"]