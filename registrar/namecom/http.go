package namecom

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strings"
)

func createClient() (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Jar: jar,
	}

	return &client, nil
}

func extractToken(client *http.Client) (string, error) {
	resp, err := client.Get("https://www.name.com/pricing")
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	pattern := "meta name=\"csrf-token\" content=\"([^\"]*)\""
	re := regexp.MustCompile(pattern)
	matches := re.FindSubmatch(body)
	if len(matches) != 2 {
		return "", errors.New("couldn't find the csrf token")
	}

	return string(matches[1]), nil
}

func getPrices(client *http.Client, token string, tld []string) ([]byte, error) {

	for i := range tld {
		tld[i] = strings.Replace(tld[i], ".", "", 1)
	}

	url := fmt.Sprintf("https://www.name.com/ajax/pricing/?duration=1&tlds=%s", strings.Join(tld, ","))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-csrf-token-auth", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func downloadPrices(tld []string) ([]byte, error) {
	client, err := createClient()
	if err != nil {
		return nil, err
	}

	token, err := extractToken(client)
	if err != nil {
		return nil, err
	}

	prices, err := getPrices(client, token, tld)
	if err != nil {
		return nil, err
	}

	return prices, nil
}
