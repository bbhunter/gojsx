package Html_parser

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Base struct {
	Url         string
	Cookies     string
	Auth        string
	Yaml_config Config
}

func (b *Base) Target_is_alive() bool {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	timeout := time.Duration(100 * time.Second)

	cli := http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	req, err := http.NewRequest("GET", b.Url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")

	if b.Cookies != "" {
		req.Header.Set("Cookie", b.Cookies)
	}

	if b.Auth != "" {
		req.Header.Set("Authorization", b.Auth)
	}

	req.Header.Set("Connection", "close")

	if err != nil {
		log.Println(err)
	}

	resp, err := cli.Do(req)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("=> Status code from %s is [%d] OK! Starting parser\n", b.Url, resp.StatusCode)
		return true
	} else {
		fmt.Println("=> Bad status code: ", resp.StatusCode)
	}

	return false
}

func (b *Base) Get_content_body() io.ReadCloser {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	timeout := time.Duration(20 * time.Second)

	cli := http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	req, err := http.NewRequest("GET", b.Url, nil)

	if b.Cookies != "" {
		req.Header.Set("Cookie", b.Cookies)
	}

	if b.Auth != "" {
		req.Header.Set("Authorization", b.Auth)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.3611")

	if err != nil {
		log.Println(err)
	}

	resp, err := cli.Do(req)

	if err != nil {
		log.Println(err)
	}

	return resp.Body

}

func (b *Base) processElement(index int, element *goquery.Selection) {

	paths := make([]string, 0)

	src, exists := element.Attr("src")
	if exists {
		if src != "" {
			paths = append(paths, src)
		}
	}

	b.parse_paths(paths)
}

func (b *Base) only_scripts_with_text() {

	html_content := b.Get_content_body()

	document, err := goquery.NewDocumentFromReader(html_content)

	if err != nil {
		log.Fatal(err)
	}

	only_text := document.Find("script").Text()
	b.Yaml_config.Regex_Matcher_Text(only_text)
}

func (b *Base) scripts_found_HTML() {

	html_content := b.Get_content_body()

	document, err := goquery.NewDocumentFromReader(html_content)

	if err != nil {
		log.Fatal(err)
	}

	document.Find("script").Each(b.processElement)
}

func (b *Base) parse_paths(paths []string) {

	parsed_paths := make([]string, 0)

	for _, path := range paths {

		if string(path[0]) == "/" && string(path[1]) != "/" {
			u := b.Url + path[1:len(path)]
			parsed_paths = append(parsed_paths, u)
		}

		if string(path[0]) == "/" && string(path[1]) == "/" {
			u := "https://" + strings.Split(path, "//")[1]
			parsed_paths = append(parsed_paths, u)
		}

		if string(path[0]) == "." && string(path[1]) == "/" {
			u := b.Url + path[2:len(path)]
			parsed_paths = append(parsed_paths, u)
		}

		if strings.Contains(path, "http://") || strings.Contains(path, "https://") {
			parsed_paths = append(parsed_paths, path)
		}
	}

	b.bulking(parsed_paths)
}

func (b *Base) Get_Page_body(url string, wg *sync.WaitGroup) {

	defer wg.Done()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	timeout := time.Duration(20 * time.Second)

	cli := http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	req.Header.Set("Connection", "Close")

	if b.Cookies != "" {
		req.Header.Set("Cookie", b.Cookies)
	}

	if b.Auth != "" {
		req.Header.Set("Authorization", b.Auth)
	}

	if err != nil {
		log.Println(err)
	}

	resp, err := cli.Do(req)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	b.Yaml_config.Regex_Matcher(resp.Body)
}

func (b *Base) bulking(paths []string) {

	var wg sync.WaitGroup

	for i := 1; i <= len(paths); i++ {
		wg.Add(1)
		go b.Get_Page_body(paths[i-1], &wg)
	}
	wg.Wait()

}

func (b *Base) Runner() {

	b.scripts_found_HTML()
	b.only_scripts_with_text()
}
