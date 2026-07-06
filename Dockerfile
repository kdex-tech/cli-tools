FROM alpine:3.19

RUN apk add --no-cache bash bind-tools coreutils curl git jq kubectl openssl tree

# durl is pinned + sha256-verified rather than installed via install.sh (which
# fetches `latest` with no checksum and no `curl -f`). durl <= v0.1.7 shipped its
# arm64 asset UPX-packed, and a UPX-packed static-musl aarch64 binary segfaults
# on real aarch64 hardware (musl's mallocng allocator fails to init on the memory
# layout UPX leaves behind) — it crashed package_image's first `durl` call 50/50
# on GKE Axion (c4a) nodes, while QEMU masked it in every image-level check.
# v0.1.8 drops UPX and ships the arm64 binary uncompressed. Pinning + checksum
# makes rebuilds reproducible and stops a bad upstream durl from silently
# breaking the packager again. See kdex-tech/cli-tools#2.
RUN set -eux; \
    VERSION="0.1.8"; \
    ARCH=$(uname -m | sed -e 's/x86_64/amd64/g' -e 's/aarch64/arm64/g'); \
    case "${ARCH}" in \
        amd64) SHA256='3d3bd92fc3a23306b8942daed29c00d7180e223b13f170fe9a899f453426114c' ;; \
        arm64) SHA256='5c66153aae2a3575254d225371f89b32649e6e270d6ffc78aadc3981a87b1f4e' ;; \
        *) echo "Unsupported ARCH for durl: '${ARCH}'" >&2; exit 1 ;; \
    esac; \
    curl -fsSL -o /usr/local/bin/durl \
        "https://github.com/rotty3000/durl/releases/download/v${VERSION}/durl-linux-${ARCH}"; \
    echo "${SHA256}  /usr/local/bin/durl" | sha256sum -c -; \
    chmod 0755 /usr/local/bin/durl

RUN OS="$(uname -s | tr A-Z a-z)"; \
    ARCH=$(uname -m | sed -e 's/x86_64/amd64/g' -e 's/aarch64/arm64/g'); \
    VERSION="1.3.0"; \
    curl -LO "https://github.com/oras-project/oras/releases/download/v${VERSION}/oras_${VERSION}_${OS}_${ARCH}.tar.gz"; \
    mkdir -p oras-install/; \
    tar -zxf oras_${VERSION}_${OS}_${ARCH}.tar.gz -C oras-install/; \
    chown root:0 oras-install/oras; \
    mv oras-install/oras /usr/local/bin/; \
    rm -rf oras_${VERSION}_${OS}_${ARCH}.tar.gz oras-install/

WORKDIR /

COPY scripts/ /usr/local/bin/

RUN chmod 777 /usr/local/bin/git_checkout; \
    chmod 777 /usr/local/bin/git_push; \
    chmod 777 /usr/local/bin/package_image; \
    chmod 777 /usr/local/bin/patch_source_status
