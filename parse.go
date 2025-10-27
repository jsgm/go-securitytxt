package securitytxt

import (
	"bufio"
	"errors"
	"fmt"
	"net/url"
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

		fmt.Println(line)

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			data.comments = append(data.comments, line)
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			data.errors = append(data.errors, errors.New("unrecognized line format"))
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Encryption":
			fmt.Println(value)

		case "Contact":
			switch {
			case strings.HasPrefix(value, "mailto:"):
				addr := strings.TrimPrefix(value, "mailto:")

				if decoded, err := url.PathUnescape(addr); err == nil {
					addr = decoded
				}

				fmt.Println("IsEmail: " + addr)

			case strings.HasPrefix(value, "tel:"):
				phone := strings.TrimPrefix(value, "tel:")
				phone = strings.ReplaceAll(phone, "-", "")
				phone = strings.ReplaceAll(phone, " ", "")
				fmt.Println("IsPhone: " + phone)

			case strings.HasPrefix(value, "http://") || strings.HasPrefix(value, "https://"):
				fmt.Println("IsURL: " + value)

			default:
				fmt.Println("Error")
			}

		case "Preferred-Languages":
			languages := strings.Split(value, ",")
			for _, language := range languages {
				language = strings.TrimSpace(strings.ToLower(language))
				if len(language) == 2 && iso6391.ValidCode(language) {
					data.preferredLanguages = append(data.preferredLanguages, iso6391.FromCode(language))
				}
			}

		case "Expires":
			// Must be uppercase for 'z' -> 'Z' (RFC3339).
			t, err := time.Parse(time.RFC3339, strings.ToUpper(value))
			if err == nil {
				data.expires = &t
			} else {
				fmt.Println(err)
			}

		case "Hiring":
			e := appendURL(&data.hiring, value)
			if e != nil {
				data.errors = append(data.errors, e)
			}

		case "Policy":
			e := appendURL(&data.policy, value)
			if e != nil {
				data.errors = append(data.errors, e)
			}

		case "Acknowledgments":
			e := appendURL(&data.acknowledgments, value)
			if e != nil {
				data.errors = append(data.errors, e)
			}

		case "Canonical":
			e := appendURL(&data.canonical, value)
			if e != nil {
				data.errors = append(data.errors, e)
			}

		default:
			data.errors = append(data.errors, fmt.Errorf("unknown key %s", key))
		}

	}

	return data, nil
}
