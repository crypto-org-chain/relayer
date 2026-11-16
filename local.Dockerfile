FROM golang:1-alpine3.17 AS build-env

RUN apk add --update --no-cache curl make git libc-dev bash gcc linux-headers eudev-dev

ADD . .

RUN CGO_ENABLED=1 LDFLAGS='-linkmode external -extldflags "-static"' make install

# CA certificates and the relayer user for the final scratch image
FROM alpine:3.22 AS base-files
RUN addgroup -g 1000 -S relayer && adduser -u 100 -S relayer -G relayer

# Statically linked busybox for the shell and utils of the final scratch image
FROM busybox:1.34.1-musl AS busybox-full

# Build final image from scratch
FROM scratch

LABEL org.opencontainers.image.source="https://github.com/cosmos/relayer"

WORKDIR /bin

# Install ln (for making hard links) and rm (for cleanup) from full busybox image (will be deleted, only needed for image assembly)
COPY --from=busybox-full /bin/ln /bin/rm ./

# Install busybox as shell binary (will create hardlinks for the rest of the binaries to this data)
COPY --from=busybox-full /bin/busybox /bin/sh

# Add hard links for read-only utils, then remove ln and rm
# Will then only have one copy of the busybox minimal binary file with all utils pointing to the same underlying inode
RUN ln sh pwd && \
    ln sh ls && \
    ln sh cat && \
    ln sh less && \
    ln sh grep && \
    ln sh sleep && \
    ln sh env && \
    ln sh tar && \
    ln sh tee && \
    ln sh du && \
    rm ln rm

# Install chain binaries
COPY --from=build-env /bin/rly /bin

# Install trusted CA certificates
COPY --from=base-files /etc/ssl/cert.pem /etc/ssl/cert.pem

# Install relayer user
COPY --from=base-files /etc/passwd /etc/passwd
COPY --from=base-files --chown=100:1000 /home/relayer /home/relayer

WORKDIR /home/relayer
USER relayer
