package securitytxt

import iso6391 "github.com/emvi/iso-639-1"

func (s *SecurityTxt) Language() iso6391.Language {
	// Returns the first appearing language.

	// Note that the order in which they appear is not an indication of priority; the
	// listed languages are intended to have equal priority.
	return s.PreferredLanguages[0]
}

func (s *SecurityTxt) ContainsLanguage(code string) bool {
	for _, language := range s.PreferredLanguages {
		if language.Code == code {
			return true
		}
	}
	return false
}
