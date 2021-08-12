package main

import (
	"testing"
)

func TestLogLink(t *testing.T) {
	tests := []struct {
		name string
		link string
	}{
		{
			name: "cdl-email-staging-email_catchall",
			link: "https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcdl-email-staging-email_catchall",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := LogLink("us-east-1", tt.name)
			if link != tt.link {
				t.Fatalf("wanted %s, got %s", tt.link, link)
			}
		})
	}
}
