package main

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	format := os.Args[1]
	if !strings.HasPrefix(format, "+") {
		fmt.Fprintf(os.Stderr, "error: format must start with +\n")
		os.Exit(1)
	}
	format = format[1:]

	rawURL := os.Args[2]
	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing URL: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(parseFormat(format, u))
}

func parseFormat(format string, u *url.URL) string {
	var sb strings.Builder
	for i := 0; i < len(format); i++ {
		if format[i] == '%' && i+1 < len(format) {
			char := format[i+1]
			switch char {
			case 'a':
				if u.User != nil {
					sb.WriteString(u.User.String())
				}
			case 'A':
				if u.User != nil {
					sb.WriteString(u.User.String())
					sb.WriteByte('@')
				}
			case 'D':
				sb.WriteString(u.Hostname())
			case 'd':
				hostname := u.Hostname()
				parts := strings.Split(hostname, ".")
				if len(parts) > 2 {
					sb.WriteString(strings.Join(parts[:len(parts)-2], "."))
				} else {
					sb.WriteString("")
				}
			case 'f':
				sb.WriteString(u.Fragment)
			case 'F':
				sb.WriteByte('#')
				sb.WriteString(u.Fragment)
			case 'H':
				sb.WriteString(u.Host)
			case 'p':
				sb.WriteString(u.Path)
			case 'b':
				sb.WriteString(path.Base(u.Path))
			case 'P':
				sb.WriteString(u.Port())
			case 'q':
				sb.WriteString(u.RawQuery)
			case 'Q':
				sb.WriteByte('?')
				sb.WriteString(u.RawQuery)
			case 's':
				sb.WriteString(u.Scheme)
			case 'S':
				if u.Scheme != "" {
					sb.WriteString(u.Scheme)
					sb.WriteString("://")
				}
			case 'u':
				if u.User != nil {
					sb.WriteString(u.User.Username())
				}
			case 'U':
				if u.User != nil {
					if p, ok := u.User.Password(); ok {
						sb.WriteString(p)
					}
				}
			default:
				sb.WriteByte('%')
				sb.WriteByte(char)
			}
			i++
		} else {
			sb.WriteByte(format[i])
		}
	}
	return sb.String()
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: gurl +<format> <url>\n\n")
	fmt.Fprintf(os.Stderr, "Format masks:\n\n")
	fmt.Fprintf(os.Stderr, "  e.g. given: https://foo:bar@www.example.com:8443/path/to/file.txt?query=value#section1\n\n")
	fmt.Fprintf(os.Stderr, "  %%s  scheme                  (gives: https)\n")
	fmt.Fprintf(os.Stderr, "  %%S  scheme with delimiter   (gives: https://)\n")
	fmt.Fprintf(os.Stderr, "  %%a  auth                    (gives: foo:bar)\n")
	fmt.Fprintf(os.Stderr, "  %%A  auth with delimiter     (gives: foo:bar@)\n")
	fmt.Fprintf(os.Stderr, "  %%u  username                (gives: foo)\n")
	fmt.Fprintf(os.Stderr, "  %%U  password                (gives: bar)\n")
	fmt.Fprintf(os.Stderr, "  %%H  host                    (gives: www.example.com:8443)\n")
	fmt.Fprintf(os.Stderr, "  %%D  domain                  (gives: www.example.com)\n")
	fmt.Fprintf(os.Stderr, "  %%d  subdomain               (gives: www)\n")
	fmt.Fprintf(os.Stderr, "  %%P  port                    (gives: 8443)\n")
	fmt.Fprintf(os.Stderr, "  %%p  path                    (gives: /path/to/file.txt)\n")
	fmt.Fprintf(os.Stderr, "  %%b  base                    (gives: file.txt)\n")
	fmt.Fprintf(os.Stderr, "  %%q  query                   (gives: query=value)\n")
	fmt.Fprintf(os.Stderr, "  %%Q  query with delimiter    (gives: ?query=value)\n")
	fmt.Fprintf(os.Stderr, "  %%f  fragment                (gives: section1)\n")
	fmt.Fprintf(os.Stderr, "  %%F  fragment with delimiter (gives: #section1)\n")
	fmt.Fprintf(os.Stderr, "\n")
}
