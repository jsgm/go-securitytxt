package securitytxt

import (
	"errors"
	"strings"
)

type Contact struct {
	Type     ContactType // Email, URL, Phone, Other
	RawValue string
	Value    string
	Priority int
}

type ContactType string

const (
	ContactEmail   ContactType = "email"
	ContactURL     ContactType = "url"
	ContactPhone   ContactType = "phone"
	ContactUnknown ContactType = "unknown"
)

func (s *SecurityTxt) PreferredContact() (Contact, error) {
	for _, contact := range s.Contact {
		if !contact.IsUnknown() {
			return contact, nil
		}
	}

	return Contact{}, nil
}

func (c Contact) IsEmail() bool {
	return c.Type == ContactEmail
}

func (c Contact) IsPhone() bool {
	return c.Type == ContactPhone
}

func (c Contact) IsURL() bool {
	return c.Type == ContactURL
}

func (c Contact) IsUnknown() bool {
	return c.Type == ContactUnknown
}

func (c Contact) Valid() bool {
	if c.IsUnknown() {
		return false
	}

	return false
}

func parseContact(v string) (Contact, []error) {
	v = strings.TrimSpace(v)
	if len(v) == 0 {
		return Contact{
			Type:     ContactUnknown,
			RawValue: v,
			Value:    v,
			Priority: -1,
		}, []error{errors.New("empty contact field")}
	}
	switch {
	case strings.HasPrefix(v, "mailto:"):
		return Contact{
			Type:     ContactEmail,
			RawValue: v,
			Value:    v,
			Priority: 0,
		}, nil

	case strings.HasPrefix(v, "tel:"):
		return Contact{
			Type:     ContactPhone,
			RawValue: v,
			Value:    v,
			Priority: 0,
		}, nil

	case strings.HasPrefix(v, "http://"), strings.HasPrefix(v, "https://"):
		return Contact{
			Type:     ContactURL,
			RawValue: v,
			Value:    v,
			Priority: 0,
		}, nil

	default:
		return Contact{
			Type:     ContactUnknown,
			RawValue: v,
			Value:    v,
			Priority: 0,
		}, nil
	}
}
