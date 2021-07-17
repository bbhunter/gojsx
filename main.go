package main

import (
	"bufio"
	"flag"
	Html_parser "gojsx/Core"
	verify "gojsx/Utils"
	"os"
	"sync"
)

func main() {
	verify.PrintBanner()
	var url = flag.String("url", "", "=> url of your target")
	var auth = flag.String("auth", "", "=> cookies from app")
	var token = flag.String("tk", "", "=> Authorization tokens, like: Athorization Bearer.. JWT..")
	var config = flag.String("config", "", "=> Config file regex ")
	flag.Parse()

	if *url == "" {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			var wg sync.WaitGroup
			sc := bufio.NewScanner(os.Stdin)
			for sc.Scan() {
				URL := sc.Text()
				wg.Add(1)

				go func() {
					defer wg.Done()
					parsed_url := verify.Verify_url(URL)
					gojsx := new(Html_parser.Base)
					gojsx.Url = parsed_url
					if gojsx.Target_is_alive() {
						if *config != "" {
							gojsx.Yaml_config.Yaml_path = *config
							gojsx.Runner()
						} else {
							gojsx.Yaml_config.Yaml_path = "./Config/regexs.yaml"
							gojsx.Runner()
						}
					}
				}()
			}
			wg.Wait()
		}
		os.Exit(0)
	}

	parsed_url := verify.Verify_url(*url)
	gojsx := new(Html_parser.Base)
	gojsx.Url = parsed_url
	gojsx.Cookies = *auth
	gojsx.Auth = *token

	if gojsx.Target_is_alive() {
		if *config != "" {
			gojsx.Yaml_config.Yaml_path = *config
			gojsx.Runner()
		} else {
			gojsx.Yaml_config.Yaml_path = "./Config/regexs.yaml"
			gojsx.Runner()
		}
	}
}
