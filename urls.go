package securitytxt

import (
	"fmt"
	"net/url"
)

func (s URLSet) First() *url.URL {
	if len(s) == 0 {
		return nil
	}
	return s[0]
}

func appendURL(target *URLSet, value string) error {
	u, err := url.Parse(value)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return fmt.Errorf(`the line %q is not a valid url`, value)
	}

	if u.Scheme != "https" {
		return fmt.Errorf(`url %q must begin with "https"`, value)
	}
	*target = append(*target, u)

	return nil
}
