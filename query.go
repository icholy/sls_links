package main

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// Go's url.QueryEscape differes from javascript's encodeURIComponent.
// It does follow the spec, but it causes issues with the aws urls.
var unescape = strings.NewReplacer(
	"+", "%20",
	"%21", "!",
	"%27", "'",
	"%28", "(",
	"%29", ")",
	"%2A", "*",
)

func EncodeURIComponent(s string) string {
	s = url.QueryEscape(s)
	return unescape.Replace(s)
}

type Query map[string][]string

func (q Query) Add(key, value string, quote bool) {
	value = url.QueryEscape(value)
	value = unescape.Replace(value)
	value = strings.ReplaceAll(value, "%", "*")
	if quote {
		value = "'" + value
	}
	q[key] = append(q[key], value)
}

func (q Query) Encode(name string) string {
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
	query := "?" + name + "=" + Escape(b.String())
	query = url.QueryEscape(query)
	query = strings.ReplaceAll(query, "%", "$")
	return query
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
