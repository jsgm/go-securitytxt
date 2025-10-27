package securitytxt

import (
	"net/url"
	"time"

	iso6391 "github.com/emvi/iso-639-1"
)

type URLSet []*url.URL

type SecurityTxt struct {
	// security.txt fields
	Acknowledgments    URLSet
	Canonical          URLSet
	Contact            []Contact
	Encryption         []string
	Expires            *time.Time
	Hiring             URLSet
	Policy             URLSet
	PreferredLanguages []iso6391.Language

	// Conformance flag according to RFC 9116
	IsValid bool

	// Contains parsing errors. For example, if the file is missing the "Contact:" field, an error will be stored here.
	Errors []error
}
