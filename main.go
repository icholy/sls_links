package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"gopkg.in/yaml.v3"
)

type Lambda struct {
	Name    string
	Link    string
	LogLink string
}

type ServerlessYML struct {
	Service   string              `yaml:"service"`
	Functions map[string]struct{} `yaml:"functions"`
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

func (s *ServerlessYML) Lambdas(region, env string) []Lambda {
	lambdas := []Lambda{}
	for key := range s.Functions {
		name := fmt.Sprintf("%s-%s-%s", s.Service, env, key)
		lambdas = append(lambdas, Lambda{
			Name: name,
			Link: fmt.Sprintf("https://console.aws.amazon.com/lambda/home?region=%s#/functions/%s", region, name),
			LogLink: fmt.Sprintf(
				"https://console.aws.amazon.com/cloudwatch/home?region=%s#logsV2:log-groups/log-group/%s",
				region,
				escape.Replace("/aws/lambda/"+name),
			),
		})
	}
	return lambdas
}

func main() {
	var region, env, filename string
	flag.StringVar(&region, "region", "us-east-1", "aws region")
	flag.StringVar(&env, "env", "staging", "deployment environment")
	flag.StringVar(&filename, "file", "serverless.yml", "serverless configuration file")
	flag.Parse()
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var config ServerlessYML
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, lambda := range config.Lambdas(region, env) {
		fmt.Fprintf(tw, "Lambda:\t%s\n", lambda.Name)
		fmt.Fprintf(tw, "Link:\t%s\n", lambda.Link)
		fmt.Fprintf(tw, "LogLink:\t%s\n\n", lambda.LogLink)
	}
	if err := tw.Flush(); err != nil {
		log.Fatal(err)
	}
}
