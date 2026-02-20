# KDex CLI Tools

## gurl

Parse URLs from the CLI like a boss.

```shell
Usage: gurl +<format> <url>

Format masks:

  e.g. given: https://foo:bar@www.example.com:8443/path/to/file.txt?query=value#section1

  %s  scheme                  (gives: https)
  %S  scheme with delimiter   (gives: https://)
  %a  auth                    (gives: foo:bar)
  %A  auth with delimiter     (gives: foo:bar@)
  %u  username                (gives: foo)
  %U  password                (gives: bar)
  %H  host                    (gives: www.example.com:8443)
  %D  domain                  (gives: www.example.com)
  %d  subdomain               (gives: www)
  %P  port                    (gives: 8443)
  %p  path                    (gives: /path/to/file.txt)
  %b  base                    (gives: file.txt)
  %q  query                   (gives: query=value)
  %Q  query with delimiter    (gives: ?query=value)
  %f  fragment                (gives: section1)
  %F  fragment with delimiter (gives: #section1)
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