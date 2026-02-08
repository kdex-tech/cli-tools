package main

import (
	"fmt"
	"net/url"
	"os"
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
			case 'd':
				sb.WriteString(u.Hostname())
			case 'f':
				sb.WriteString(u.Fragment)
			case 'F':
				sb.WriteByte('#')
				sb.WriteString(u.Fragment)
			case 'H':
				sb.WriteString(u.Host)
			case 'p':
				sb.WriteString(u.Path)
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
	fmt.Fprintf(os.Stderr, "Format masks:\n")
	fmt.Fprintf(os.Stderr, "  %%s  scheme (e.g., https)\n")
	fmt.Fprintf(os.Stderr, "  %%S  scheme delimiter (e.g., ://)\n")
	fmt.Fprintf(os.Stderr, "  %%a  auth (e.g., user:pass)\n")
	fmt.Fprintf(os.Stderr, "  %%A  auth with delimiter (e.g., user:pass@)\n")
	fmt.Fprintf(os.Stderr, "  %%u  username (e.g., user)\n")
	fmt.Fprintf(os.Stderr, "  %%U  password (e.g., pass)\n")
	fmt.Fprintf(os.Stderr, "  %%H  host (e.g., example.com:port)\n")
	fmt.Fprintf(os.Stderr, "  %%d  domain (e.g., example.com)\n")
	fmt.Fprintf(os.Stderr, "  %%P  port (e.g., 8080)\n")
	fmt.Fprintf(os.Stderr, "  %%p  path (e.g., /index.html)\n")
	fmt.Fprintf(os.Stderr, "  %%q  query (e.g., a=1&b=2)\n")
	fmt.Fprintf(os.Stderr, "  %%Q  query with delimiter (e.g., ?a=1&b=2)\n")
	fmt.Fprintf(os.Stderr, "  %%f  fragment (e.g., section1)\n")
	fmt.Fprintf(os.Stderr, "  %%F  fragment with delimiter (e.g., #section1)\n")
}
