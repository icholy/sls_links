package main

import (
	"fmt"
	"strings"
)

func Link(region, name string) string {
	return fmt.Sprintf("https://console.aws.amazon.com/lambda/home?region=%s#/functions/%s", region, name)
}

var escape = strings.NewReplacer(
	"/", "$252F",
	",", "$252C",
	"[", "$255B",
	"]", "$255D",
	"=", "$253D",
	"!", "$2521",
	"\"", "$2522",
	"_", "$252F",
)

func LogLink(region, lambda string) string {
	return fmt.Sprintf(
		"https://console.aws.amazon.com/cloudwatch/home?region=%s#logsV2:log-groups/log-group/%s",
		region,
		escape.Replace("/aws/lambda/"+lambda),
	)
}
