# KDex CLI Tools

## gurl

Parse URLs from the CLI like a boss.

```shell
Usage: gurl +<format> <url>

Format masks:
  %s  scheme (e.g., https)
  %S  scheme with delimiter (e.g., https://)
  %a  auth (e.g., user:pass)
  %A  auth with delimiter (e.g., user:pass@)
  %u  username (e.g., user)
  %U  password (e.g., pass)
  %H  host (e.g., example.com:port)
  %d  domain (e.g., example.com)
  %P  port (e.g., 8080)
  %p  path (e.g., /index.html)
  %q  query (e.g., a=1&b=2)
  %Q  query with delimiter (e.g., ?a=1&b=2)
  %f  fragment (e.g., section1)
  %F  fragment with delimiter (e.g., #section1)
```

## Examples

```shell
# Get the domain
]$ gurl +%d https://example.com:8080
example.com

# Add a user and password
]$ gurl +%Susername:password@%H%p%Q%F https://example.com:8080/path?a=1&b=2#section1
https://username:password@example.com:8080/path?a=1&b=2#section1

```