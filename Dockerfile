FROM alpine:3.19

RUN apk add --no-cache bind-tools curl git jq tree; \
    curl -sSL https://raw.githubusercontent.com/rotty3000/durl/main/scripts/install.sh | sh

WORKDIR /

COPY scripts/ /usr/local/bin/

RUN chmod 777 /usr/local/bin/git_checkout; \
    chmod 777 /usr/local/bin/git_push; \
    chmod 777 /usr/local/bin/patch_source_status
