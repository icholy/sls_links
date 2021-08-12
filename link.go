package main

import (
	"fmt"
	"strings"
)

func Link(region, name string) string {
	return fmt.Sprintf("https://console.aws.amazon.com/lambda/home?region=%s#/functions/%s", region, name)
}

func LogLink(region, lambda string) string {
	group := "/aws/lambda/" + lambda
	group = QueryEscape(group)
	group = QueryEscape(group)
	group = strings.ReplaceAll(group, "%", "$")
	return fmt.Sprintf(
		"https://console.aws.amazon.com/cloudwatch/home?region=%s#logsV2:log-groups/log-group/%s",
		region,
		group,
	)
}
