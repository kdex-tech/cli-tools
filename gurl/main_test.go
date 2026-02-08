package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseFormat(t *testing.T) {
	tests := []struct {
		name   string
		format string
		u      *url.URL
		want   string
	}{
		{
			name:   "scheme",
			format: "%s",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "https",
		},
		{
			name:   "separator",
			format: "%S",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "://",
		},
		{
			name:   "auth",
			format: "%a",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "user:pass",
		},
		{
			name:   "user",
			format: "%u",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "user",
		},
		{
			name:   "password",
			format: "%U",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "pass",
		},
		{
			name:   "host",
			format: "%H",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "example.com:8080",
		},
		{
			name:   "domain",
			format: "%d",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "example.com",
		},
		{
			name:   "port",
			format: "%P",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "8080",
		},
		{
			name:   "path",
			format: "%p",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "/index.html",
		},
		{
			name:   "query",
			format: "%q",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "a=1&b=2",
		},
		{
			name:   "query with question mark",
			format: "%Q",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "?a=1&b=2",
		},
		{
			name:   "fragment",
			format: "%f",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "section1",
		},
		{
			name:   "fragment with hash mark",
			format: "%F",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "#section1",
		},
		{
			name:   "all",
			format: "%s%S%A%H%p%Q%F",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "https://user:pass@example.com:8080/index.html?a=1&b=2#section1",
		},
		{
			name:   "without auth",
			format: "%s%S%H%p%Q%F",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "https://example.com:8080/index.html?a=1&b=2#section1",
		},
		{
			name:   "without port",
			format: "%s%S%H%p%Q%F",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "https://example.com/index.html?a=1&b=2#section1",
		},
		{
			name:   "without query",
			format: "%s%S%H%p%F",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com",
				Path:     "/index.html",
				Fragment: "section1",
			},
			want: "https://example.com/index.html#section1",
		},
		{
			name:   "with different scheme",
			format: "%s%S%H%p%F",
			u: &url.URL{
				Scheme:   "ftp",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com",
				Path:     "/index.html",
				Fragment: "section1",
			},
			want: "ftp://example.com/index.html#section1",
		},
		{
			name:   "with litterals",
			format: "http%S%H/foo/bar%p%F",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com",
				Path:     "/index.html",
				Fragment: "section1",
			},
			want: "http://example.com/foo/bar/index.html#section1",
		},
		{
			name:   "all parts",
			format: "%s|%S|%a|%A|%u|%U|%H|%d|%P|%p|%q|%Q|%f|%F",
			u: &url.URL{
				Scheme:   "https",
				User:     url.UserPassword("user", "pass"),
				Host:     "example.com:8080",
				Path:     "/index.html",
				RawQuery: "a=1&b=2",
				Fragment: "section1",
			},
			want: "https|://|user:pass|user:pass@|user|pass|example.com:8080|example.com|8080|/index.html|a=1&b=2|?a=1&b=2|section1|#section1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseFormat(tt.format, tt.u)
			assert.Equal(t, tt.want, got)
		})
	}
}
