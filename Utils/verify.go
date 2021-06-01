package verify

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func Verify_url(url_input string) string {
	u, err := url.Parse(url_input)

	if err != nil {
		log.Println("error while parsing the URL")
		log.Println(err)
	}

	if u.Host == "" || u.Scheme == "" {
		fmt.Println("=> Scheme or Host is invalid..")
		os.Exit(1)
	}

	if u.Path == "" {
		target_with_slash := url_input + "/"
		return target_with_slash
	} else {
		return url_input
	}

	/*
		URL could be https://foo.bar/path1/path2/pathN..
	*/

}

func Target_is_alive(url_input string) bool {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	timeout := time.Duration(100 * time.Second)

	cli := http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	req, err := http.NewRequest("GET", url_input, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
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
		fmt.Printf("=> Status code from %s is [%d] OK!\n", url_input, resp.StatusCode)
		return true
	} else {
		fmt.Println("=> Bad status code: ", resp.StatusCode)
	}

	return false
}

func Remove_duplicates_paths(paths []string) {

	keys := make(map[string]bool)
	list := []string{}

	for _, path := range paths {
		if !inside_deny_list(path) {
			if _, value := keys[path]; !value {
				keys[path] = true
				list = append(list, path)
			}
		}
	}

	show_results(list)

}

func inside_deny_list(path string) bool {
	slice_deny_list := []string{"https://momentjs.com", "twitter.com", "http://fb.me", "www.youtube.com", "http://www.w3.org", "https://developer.mozilla.org", "http://momentjs.com", "http://www.w3.org/1999/xlink", "http://www.w3.org/2000/svg", "https://npms.io/", "http://www.w3.org/2000", "https://npms.io/search?q=ponyfill", "www.w3.org", "www.googleoptimize.com", "modernizr.com", "dev.w3.org", "schema.org", "google.com", "bit.ly", "https://schema.org", "linkdedin.com", "example.com", "https://developer.mozilla.org/en-US/docs/Web/API/CustomEvent/CustomEvent#Polyfill"}

	for _, deny_url := range slice_deny_list {
		if strings.Contains(path, deny_url) {
			return true
		}
	}
	return false
}

func show_results(results []string) {
	for _, v := range results {
		fmt.Println("=> Found: ", v)
	}
}
