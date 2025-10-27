package securitytxt

import (
	"net/url"
	"time"

	iso6391 "github.com/emvi/iso-639-1"
)

type URLSet []*url.URL

type SecurityTxt struct {
	// security.txt fields
	acknowledgments    URLSet
	canonical          URLSet
	contact            []string
	encryption         []string
	expires            *time.Time
	hiring             URLSet
	policy             URLSet
	preferredLanguages []iso6391.Language

	// File comments
	comments []string

	// Contains parsing errors. For example, if the file is missing the "Contact:" field, an error will be stored here.
	errors []error
}
