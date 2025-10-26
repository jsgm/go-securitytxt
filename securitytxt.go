package securitytxt

import (
	"errors"
	"net/url"
	"time"

	iso6391 "github.com/emvi/iso-639-1"
)

var (
	ErrContactNotSet = errors.New("contact method is required, but not set")
)

type SecurityTxt struct {
	// security.txt fields
	Acknowledgments    []*url.URL
	Canonical          []*url.URL
	Contact            []string
	Encryption         []string
	Expires            *time.Time
	Hiring             []*url.URL
	Policy             []*url.URL
	preferredLanguages []iso6391.Language

	// File comments
	comments []string

	// Contains parsing errors. For example, if the file is missing the "Contact:" field, an error will be stored here.
	errors []error
}
