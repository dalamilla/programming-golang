package types

import (
	"errors"
	"net"
	"net/url"
)

type Shorter struct {
	ShortURL    uint64 `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type ShorterError struct {
	Error string `json:"error"`
}

type ShorterPayload struct {
	URL string `json:"url" form:"url"`
}

func (p *ShorterPayload) Validate() error {
	u, err := url.ParseRequestURI(p.URL)
	if err != nil || u.Host == "" || u.Scheme != "https" && u.Scheme != "http" {
		return errors.New("Invalid URL")
	}

	_, err = net.LookupHost(u.Host)
	if err != nil {
		return errors.New("Invalid Hostname")
	}
	return nil
}
