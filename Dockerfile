FROM alpine:3.19

RUN apk add --no-cache bash bind-tools curl git jq tree; \
    curl -sSL https://raw.githubusercontent.com/rotty3000/durl/main/scripts/install.sh | sh

RUN OS="$(uname -s | tr A-Z a-z)"; \
    ARCH=$(uname -m | sed -e 's/x86_64/amd64/g'); \
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
