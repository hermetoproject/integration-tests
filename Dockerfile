FROM golang:1.23-alpine AS builder

WORKDIR /workspace

# Copy both modules
COPY vendored-module/ ./vendored-module/
COPY non-vendored-module/ ./non-vendored-module/

# Build vendored-module
WORKDIR /workspace/vendored-module
RUN CGO_ENABLED=0 go build -o /bin/vendored-app .

# Build non-vendored-module
WORKDIR /workspace/non-vendored-module
RUN CGO_ENABLED=0 go build -o /bin/non-vendored-app .

# Final stage
FROM scratch

COPY --from=builder /bin/vendored-app /
COPY --from=builder /bin/non-vendored-app /

CMD ["/non-vendored-app"]
