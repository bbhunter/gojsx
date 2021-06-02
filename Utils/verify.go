package verify

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
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
