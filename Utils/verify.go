package verify

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"runtime"
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

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

func Remove_duplicates_paths(paths []string, url string) {

	keys := make(map[string]bool)
	list := []string{}

	for _, path := range paths {
		if !inside_deny_list(path) {
			if _, value := keys[path]; !value {
				keys[path] = true
				if path != "" {
					list = append(list, path)
				}
			}
		}
	}

	show_results(list, url)

}

func inside_deny_list(path string) bool {
	slice_deny_list := []string{"momentjs", "googletagmanager", "github.com", "gstatic.com", "googleapis.com", "twitter.com", "fb.me", "youtube.com", "w3.org", "developer.mozilla.org", "npms.io", "googleoptimize.com", "modernizr.com", "dev.w3.org", "schema.org", "google.com", "bit.ly", "schema.org", "linkdedin.com", "example.com", "mozilla.org", "jquery.org", "reactjs.org", "doubleclick.net", "google", "google-analytics"}

	for _, deny_url := range slice_deny_list {
		if strings.Contains(path, deny_url) {
			return true
		} else {
			return false
		}
	}
	return false
}

func show_results(results []string, url string) {
	for _, v := range results {
		if v != "" {
			fmt.Println("##############################")
			u := fmt.Sprintf("\n"+Yellow+"[*] Results: [%s]"+Reset, url)
			u_found := fmt.Sprintf("\n"+Cyan+"[X] Found: %s\n"+Reset, v)
			fmt.Printf("\n%s\n%s\n", u, u_found)
			fmt.Println("##############################")
		}
	}
}
