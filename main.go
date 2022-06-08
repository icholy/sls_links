package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/tabwriter"

	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v3"
)

type Lambda struct {
	Name       string
	Region     string
	LambdaLink string
	LogLink    string
	LogGroup   string
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
			Name:       name,
			Region:     region,
			LambdaLink: LambdaLink(region, name),
			LogLink:    LogLink(region, name),
			LogGroup:   LambdaLogGroup(name),
		})
	}
	return lambdas
}

func ReadServerlessYML(name string) (*ServerlessYML, error) {
	// if we've been given a directory, look for a serverless.yml file in it
	info, err := os.Stat(name)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return ReadServerlessYML(filepath.Join(name, "serverless.yml"))
	}
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	var config ServerlessYML
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	var region, env string
	var openlambda, openlogs, openall, tail bool
	flag.StringVar(&region, "region", "us-east-1", "aws region")
	flag.StringVar(&env, "env", "staging", "deployment environment")
	flag.BoolVar(&openlambda, "open.lambdas", false, "open all lambda links in default browser")
	flag.BoolVar(&openlogs, "open.logs", false, "open all log links in default browser")
	flag.BoolVar(&openall, "open", false, "open all links in default browser")
	flag.BoolVar(&tail, "tail", false, "tail logs from all lambdas")
	flag.Parse()
	// if no paths were provided, look in the current directory
	paths := flag.Args()
	if len(paths) == 0 {
		paths = []string{"."}
	}
	lambdas := []Lambda{}
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, path := range paths {
		config, err := ReadServerlessYML(path)
		if err != nil {
			log.Fatal(err)
		}
		lambdas = append(lambdas, config.Lambdas(region, env)...)
	}
	if tail {
		var g errgroup.Group
		for _, lambda := range lambdas {
			lambda := lambda
			g.Go(func() error {
				cmd := exec.Command("aws", "logs", "tail", lambda.LogGroup, "--follow")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				fmt.Println(cmd)
				return cmd.Run()
			})
		}
		if err := g.Wait(); err != nil {
			log.Fatal(err)
		}
		return
	}
	for _, lambda := range lambdas {
		fmt.Fprintf(tw, "Name:\t%s\n", lambda.Name)
		fmt.Fprintf(tw, "Lambda:\t%s\n", lambda.LambdaLink)
		fmt.Fprintf(tw, "Logs:\t%s\n\n", lambda.LogLink)
		if openall || openlambda {
			if err := OpenBrowser(lambda.LambdaLink); err != nil {
				log.Println(err)
			}
		}
		if openall || openlogs {
			if err := OpenBrowser(lambda.LogLink); err != nil {
				log.Println(err)
			}
		}
	}
	fmt.Fprintf(tw, "LogInsights:\t%s\n", LogInsightsLink(region, lambdas))
	if err := tw.Flush(); err != nil {
		log.Fatal(err)
	}
}
