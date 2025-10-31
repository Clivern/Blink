FROM golang:1.27.3 AS builder

ARG BLINK_VERSION=0.1.0
ARG BLINK_COMMIT=none
ARG BLINK_BUILD_DATE=unknown
ARG BLINK_BUILT_BY=docker
ARG BLINK_EDITION=oss

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w -X main.version=${BLINK_VERSION} -X main.commit=${BLINK_COMMIT} -X main.date=${BLINK_BUILD_DATE} -X main.builtBy=${BLINK_BUILT_BY} -X main.edition=${BLINK_EDITION}" \
    -o /out/blink .

RUN mkdir -p /out/configs /out/var/logs && \
    cp config.dist.yml /out/configs/config.dist.yml

FROM gcr.io/distroless/static-debian13:nonroot

WORKDIR /app

COPY --from=builder --chown=65532:65532 /out/blink /app/blink
COPY --from=builder --chown=65532:65532 /out/configs /app/configs
COPY --from=builder --chown=65532:65532 /out/var /app/var

EXPOSE 8080

VOLUME ["/app/configs", "/app/var"]

USER 65532:65532

ENTRYPOINT ["/app/blink"]
CMD ["server", "-c", "/app/configs/config.dist.yml"]
