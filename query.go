package main

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

var unescape = strings.NewReplacer(
	"+", "%20",
	"%21", "!",
	"%27", "'",
	"%28", "(",
	"%29", ")",
	"%2A", "*",
)

// Go's url.QueryEscape differes from javascript's encodeURIComponent.
// It does follow the spec, but it causes issues with the aws urls.
func QueryEscape(s string) string {
	s = url.QueryEscape(s)
	return unescape.Replace(s)
}

// Escape is an implementation of javascript's deprecated escape() function
func Escape(s string) string {
	var b strings.Builder
	for _, r := range s {
		if ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z') || ('0' <= r && r <= '9') ||
			r == '@' || r == '*' || r == '_' || r == '+' || r == '-' || r == '.' || r == '/' {
			b.WriteRune(r)
			continue
		}
		if r >= 256 {
			fmt.Fprintf(&b, "%%%04X", r)
		} else {
			fmt.Fprintf(&b, "%%%02X", r)
		}
	}
	return b.String()
}

type QueryDetails map[string][]string

func (q QueryDetails) Add(key, value string, quote bool) {
	escaped := QueryEscape(value)
	escaped = strings.ReplaceAll(escaped, "%", "*")
	if quote {
		escaped = "'" + escaped
	}
	q[key] = append(q[key], escaped)
}

func (q QueryDetails) Encode() string {
	var b strings.Builder
	b.WriteString("~(")
	// sort the keys to get a deterministic output order
	keys := make([]string, 0, len(q))
	for key := range q {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for i, key := range keys {
		if i > 0 {
			b.WriteByte('~')
		}
		b.WriteString(key)
		b.WriteByte('~')
		values := q[key]
		switch len(values) {
		case 0:
		case 1:
			b.WriteString(values[0])
		default:
			b.WriteByte('(')
			for _, v := range values {
				b.WriteByte('~')
				b.WriteString(v)
			}
			b.WriteByte(')')
		}
	}
	b.WriteByte(')')
	escaped := QueryEscape("?queryDetail=" + Escape(b.String()))
	escaped = strings.ReplaceAll(escaped, "%", "$")
	return escaped
}
