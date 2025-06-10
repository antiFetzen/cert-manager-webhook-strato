FROM golang:1.22.12-bookworm AS build_deps

WORKDIR /workspace

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_deps AS build

COPY . .

RUN go generate ./...
RUN CGO_ENABLED=0 go build -o webhook -ldflags '-w -extldflags "-static"' .

# TODO: check if container works also with -> FROM alpine:3.18
FROM debian:bookworm

COPY --from=build /workspace/webhook /usr/local/bin/webhook

ENTRYPOINT ["webhook"]
