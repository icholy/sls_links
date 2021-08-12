package main

import (
	"fmt"
	"net/url"
	"strings"
)

type Query map[string][]string

func (q Query) Add(key, value string, quote bool) {
	value = url.QueryEscape(value)
	value = strings.ReplaceAll(value, "%", "*")
	if quote {
		value = "'" + value
	}
	q[key] = append(q[key], value)
}

func (q Query) Encode(name string) string {
	var b strings.Builder
	b.WriteString("~(")
	first := true
	for key, values := range q {
		if first {
			first = false
		} else {
			b.WriteByte('~')
		}
		b.WriteString(key)
		b.WriteByte('~')
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
		if ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z') ||
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
