package httputil

import "net/url"

// IsValidURL return true if input is a full and valid url,e.g. https://kubeclipper.io
func IsValidURL(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return false
	}
	u, err := url.Parse(urlStr)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}
