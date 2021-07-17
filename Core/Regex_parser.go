package Html_parser

import (
	"fmt"
	verify "gojsx/Utils"

	"io"
	"io/ioutil"
	"log"
	"regexp"

	"gopkg.in/yaml.v2"
)

type Yaml_File struct {
	Regexs []string
}

type Config struct {
	Yaml_path string
}

func (c *Config) Regex_Matcher_Text(js_text string) {

	var yaml_file Yaml_File

	source, err := ioutil.ReadFile(c.Yaml_path)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &yaml_file)
	if err != nil {
		panic(err)
	}

	for _, v := range yaml_file.Regexs {
		matched, err := regexp.MatchString(v, string(js_text))
		if err != nil {
			log.Fatalln(err)
			fmt.Println("=> Unable to use the regex: ", v)
		}

		if matched {
			re := regexp.MustCompile(v).FindAllString(string(js_text), 10)
			for _, m := range re {
				fmt.Println("=> Found: ", m)
			}
		}
	}

}

func (c *Config) Regex_Matcher(js_content io.ReadCloser, url string) {

	var yaml_file Yaml_File

	source, err := ioutil.ReadFile(c.Yaml_path)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &yaml_file)
	if err != nil {
		panic(err)
	}

	slice_of_matchs := make([]string, 0)

	data, err := ioutil.ReadAll(js_content)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("=> Unable to read the js content")
	}

	for _, v := range yaml_file.Regexs {
		matched, err := regexp.MatchString(v, string(data))
		if err != nil {
			log.Fatalln(err)
			fmt.Println("=> Unable to use the regex: ", v)
			fmt.Println("=> with: ", string(data))
		}

		if matched {
			re := regexp.MustCompile(v).FindAllString(string(data), 10)
			for _, m := range re {
				if m != "" {
					slice_of_matchs = append(slice_of_matchs, m)
				}
			}
		}
	}

	verify.Remove_duplicates_paths(slice_of_matchs, url)

}
