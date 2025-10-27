package securitytxt

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"time"

	iso6391 "github.com/emvi/iso-639-1"
)

func FromURL(content string) (*SecurityTxt, error) {
	return parse(content)
}

func FromBytes(content []byte) (*SecurityTxt, error) {
	return parse(string(content))
}

func FromString(content string) (*SecurityTxt, error) {
	return parse(content)
}

func parse(content string) (*SecurityTxt, error) {
	if len(content) == 0 {
		return nil, errors.New("content must not be empty")
	}

	data := &SecurityTxt{}

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			// data.comments = append(data.comments, line)
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			data.Errors = append(data.Errors, errors.New("unrecognized line format"))
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Encryption":
			// fmt.Println(value)

		case "Contact":
			newContact, errors := parseContact(value)
			data.Contact = append(data.Contact, newContact)
			data.Errors = append(data.Errors, errors...)

		case "Preferred-Languages":
			languages := strings.Split(value, ",")
			for _, language := range languages {
				language = strings.TrimSpace(strings.ToLower(language))
				if len(language) == 2 && iso6391.ValidCode(language) {
					data.PreferredLanguages = append(data.PreferredLanguages, iso6391.FromCode(language))
				}
			}

		case "Expires":
			t, err := time.Parse(time.RFC3339, strings.ToUpper(value))
			if err == nil {
				data.Expires = &t
			} else {
				data.Errors = append(data.Errors, err)
			}

		case "Hiring":
			e := appendURL(&data.Hiring, value)
			if e != nil {
				data.Errors = append(data.Errors, e)
			}

		case "Policy":
			e := appendURL(&data.Policy, value)
			if e != nil {
				data.Errors = append(data.Errors, e)
			}

		case "Acknowledgments":
			e := appendURL(&data.Acknowledgments, value)
			if e != nil {
				data.Errors = append(data.Errors, e)
			}

		case "Canonical":
			e := appendURL(&data.Canonical, value)
			if e != nil {
				data.Errors = append(data.Errors, e)
			}

		default:
			data.Errors = append(data.Errors, fmt.Errorf(`unknown key %q`, key))
		}

	}

	if len(data.PreferredLanguages) == 0 {
		data.PreferredLanguages = []iso6391.Language{iso6391.FromCode("en")}
	}

	return data, nil
}

func (s *SecurityTxt) Valid() bool {
	return len(s.Errors) == 0
}
