package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

func (s *ServerlessYML) Lambdas(region, env string) []Lambda {
	lambdas := []Lambda{}
	for key := range s.Functions {
		name := fmt.Sprintf("%s-%s-%s", s.Service, env, key)
		lambdas = append(lambdas, Lambda{
			Name:    name,
			Link:    Link(region, name),
			LogLink: LogLink(region, name),
		})
	}
	return lambdas
}

func (s *ServerlessYML) LogInsightsURL(region, env string) string {
	query := Query{}
	for _, lambda := range s.Lambdas(region, env) {
		query.Add("source", "/aws/lambda/"+lambda.Name, true)
	}
	return fmt.Sprintf(
		"https://console.aws.amazon.com/cloudwatch/home?region=%s#logsV2:logs-insights%s",
		region,
		query.Encode("queryDetail"),
	)
}

func main() {
	var region, env, filename string
	var openlambda, openlogs, openall bool
	flag.StringVar(&region, "region", "us-east-1", "aws region")
	flag.StringVar(&env, "env", "staging", "deployment environment")
	flag.StringVar(&filename, "file", "serverless.yml", "serverless configuration file")
	flag.BoolVar(&openlambda, "open.lambdas", false, "open all lambda links in default browser")
	flag.BoolVar(&openlogs, "open.logs", false, "open all log links in default browser")
	flag.BoolVar(&openall, "open", false, "open all links in default browser")
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
		if openall || openlambda {
			if err := OpenBrowser(lambda.Link); err != nil {
				log.Println(err)
			}
		}
		if openall || openlogs {
			if err := OpenBrowser(lambda.LogLink); err != nil {
				log.Println(err)
			}
		}
	}
	if err := tw.Flush(); err != nil {
		log.Fatal(err)
	}
}
