package main

import "net/http"

func main() {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	url := "http://example.com"
	resp, err := client.Get(url)
	// ...
	req, err := http.NewRequest("GET", url, nil)
	// ...

	req.Header.Add("User-Agent", "Our Custom User-Agent")
	req.Header.Add("If-None-Match", `W/"TheFileEtag"`)
	resp, err := client.Do(req)
}
