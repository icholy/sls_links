package main

import (
	"fmt"
	"net/url"
	"strings"
)

func Link(region, name string) string {
	return fmt.Sprintf("https://console.aws.amazon.com/lambda/home?region=%s#/functions/%s", region, name)
}

func LogLink(region, lambda string) string {
	group := "/aws/lambda/" + lambda
	group = url.QueryEscape(group)
	group = url.QueryEscape(group)
	group = strings.ReplaceAll(group, "%", "$")
	return fmt.Sprintf(
		"https://console.aws.amazon.com/cloudwatch/home?region=%s#logsV2:log-groups/log-group/%s",
		region,
		group,
	)
}
