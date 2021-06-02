package main

import (
	"flag"
	"fmt"
	Html_parser "gojsx/Core"
	verify "gojsx/Utils"
	"os"
)

const (
	author = "0xrod"
)

func main() {
	fmt.Println("=> gojsx by: ", author)
	var url = flag.String("url", "", "=> url of your target")
	var auth = flag.String("auth", "", "=> cookies from app")
	var token = flag.String("tk", "", "=> Authorization tokens, like: Athorization Bearer.. JWT..")
	var config = flag.String("config", "", "=> Config file regex ")
	flag.Parse()

	if *url == "" {
		fmt.Println("=> URL not defined..")
		fmt.Printf("\n Usage: %s -url https://target.com\n", os.Args[0])
	}

	parsed__url := verify.Verify_url(*url)

	gojsx := new(Html_parser.Base)
	gojsx.Url = parsed__url
	gojsx.Cookies = *auth
	gojsx.Auth = *token

	if gojsx.Target_is_alive() {
		if *config != "" {
			fmt.Println("=> Config file: ", *config)
			gojsx.Yaml_config.Yaml_path = *config
			gojsx.Get_content_body(parsed__url)
		} else {
			fmt.Println("=> Config file: ./Config/regexs.yaml")
			gojsx.Yaml_config.Yaml_path = "./Config/regexs.yaml"
			gojsx.Get_content_body(parsed__url)
		}
	}
}
