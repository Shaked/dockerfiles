# syntax=docker/dockerfile:1.8-labs
FROM golang:1.21 AS base
WORKDIR /workspace

ARG CGO_ENABLED=0
ARG GOMODCACHE=/root/.cache/go-lib
ARG GOCACHE=/root/.cache/go-build

COPY go.mod go.mod
COPY go.sum go.sum
RUN --mount=type=cache,target="$GOMODCACHE" go mod download

COPY --parents cmd internal ./

FROM base AS test

ENTRYPOINT [ "go" ]
CMD [ "test", "./...", "-v", "-coverprofile", "cover.out" ]

FROM base AS builder
ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target="$GOMODCACHE" \
    --mount=type=cache,target="$GOCACHE" \
    GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -o manager cmd/main.go

FROM gcr.io/distroless/static:nonroot AS runner
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]
