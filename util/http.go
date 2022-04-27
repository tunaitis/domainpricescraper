package util

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func download(u string, cookies []*http.Cookie, transport *http.Transport) (io.Reader, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	if cookies != nil {
		parsedUrl, err := url.Parse(u)
		if err != nil {
			return nil, err
		}
		jar.SetCookies(parsedUrl, cookies)
	}

	client := http.Client{
		Jar: jar,
	}

	if transport != nil {
		client.Transport = transport
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Safari/605.1.15")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func DownloadString(u string, cookies []*http.Cookie, transport *http.Transport) (string, error) {
	r, err := download(u, cookies, transport)
	if err != nil {
		return "", err
	}

	d, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(d), nil
}

func DownloadDocument(u string, cookies []*http.Cookie, transport *http.Transport) (*goquery.Document, error) {
	r, err := download(u, cookies, transport)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
